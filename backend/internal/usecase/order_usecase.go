package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"

	"ecommerce-backend/internal/domain"
)

type orderUseCase struct {
	orderRepo   domain.OrderRepository
	cartRepo    domain.CartRepository
	productRepo domain.ProductRepository
}

// NewOrderUseCase creates a new order use case
func NewOrderUseCase(
	orderRepo domain.OrderRepository,
	cartRepo domain.CartRepository,
	productRepo domain.ProductRepository,
) domain.OrderUseCase {
	return &orderUseCase{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (uc *orderUseCase) Checkout(ctx context.Context, userID string, req *domain.CreateOrderRequest) (*domain.OrderResponse, error) {
	// 1. Get user cart
	cart, err := uc.cartRepo.GetByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrCartNotFound) {
			return nil, domain.ErrEmptyCart
		}
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}

	if cart == nil || len(cart.Items) == 0 {
		return nil, domain.ErrEmptyCart
	}

	var orderItems []domain.OrderItem
	var totalAmount float64

	// 2. Validate product stock and build order items
	for _, item := range cart.Items {
		product, err := uc.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product %s: %w", item.ProductID, err)
		}

		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("%w: product %s (available: %d, requested: %d)",
				domain.ErrInsufficientStock, product.Name, product.Stock, item.Quantity)
		}

		subTotal := product.Price * float64(item.Quantity)
		totalAmount += subTotal

		orderItems = append(orderItems, domain.OrderItem{
			ProductID: item.ProductID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  item.Quantity,
			SubTotal:  subTotal,
		})

		// 3. Deduct stock
		product.Stock -= item.Quantity
		if err := uc.productRepo.Update(ctx, product); err != nil {
			return nil, fmt.Errorf("failed to update product stock: %w", err)
		}
	}

	// 4. Create Order
	order := &domain.Order{
		UserID:          userID,
		Items:           orderItems,
		TotalAmount:     totalAmount,
		Status:          domain.OrderStatusPending,
		ShippingAddress: req.ShippingAddress,
	}

	if err := uc.orderRepo.Create(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// 5. Clear Cart — non-fatal: log warning but do not fail the checkout
	if err := uc.cartRepo.DeleteByUserID(ctx, userID); err != nil {
		log.Printf("warning: failed to clear cart for user %s after checkout: %v", userID, err)
	}

	return uc.mapToResponse(order), nil
}

func (uc *orderUseCase) GetMyOrders(ctx context.Context, userID string, page, limit int) (*domain.PaginatedOrderResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	orders, total, err := uc.orderRepo.FindByUserID(ctx, userID, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	var responseItems []*domain.OrderResponse
	for _, o := range orders {
		responseItems = append(responseItems, uc.mapToResponse(o))
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &domain.PaginatedOrderResponse{
		Data:       responseItems,
		Page:       page,
		PerPage:    limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, orderID, userID string, isAdmin bool) (*domain.OrderResponse, error) {
	order, err := uc.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	// Verify ownership if not admin
	if !isAdmin && order.UserID != userID {
		return nil, domain.ErrNotOrderOwner
	}

	return uc.mapToResponse(order), nil
}

func (uc *orderUseCase) CancelOrder(ctx context.Context, orderID, userID string) error {
	order, err := uc.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order: %w", err)
	}

	if order.UserID != userID {
		return domain.ErrNotOrderOwner
	}

	if order.Status != domain.OrderStatusPending {
		return fmt.Errorf("%w: only pending orders can be cancelled", domain.ErrInvalidOrderStatus)
	}

	// Restore stock — non-fatal per item: log each failure but continue
	for _, item := range order.Items {
		product, err := uc.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			log.Printf("warning: cannot restore stock for product %s (order %s): %v", item.ProductID, orderID, err)
			continue
		}
		product.Stock += item.Quantity
		if err := uc.productRepo.Update(ctx, product); err != nil {
			log.Printf("warning: failed to update stock for product %s (order %s): %v", item.ProductID, orderID, err)
		}
	}

	if err := uc.orderRepo.UpdateStatus(ctx, orderID, domain.OrderStatusCancelled); err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
}

func (uc *orderUseCase) UpdateOrderStatus(ctx context.Context, orderID string, req *domain.UpdateOrderStatusRequest) error {
	order, err := uc.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order: %w", err)
	}

	// Validate state machine transitions
	if err := validateStatusTransition(order.Status, req.Status); err != nil {
		return err
	}

	// Restore stock if admin is cancelling a non-cancelled order
	if req.Status == domain.OrderStatusCancelled && order.Status != domain.OrderStatusCancelled {
		for _, item := range order.Items {
			product, err := uc.productRepo.GetByID(ctx, item.ProductID)
			if err != nil {
				log.Printf("warning: cannot restore stock for product %s (order %s): %v", item.ProductID, orderID, err)
				continue
			}
			product.Stock += item.Quantity
			if err := uc.productRepo.Update(ctx, product); err != nil {
				log.Printf("warning: failed to update stock for product %s (order %s): %v", item.ProductID, orderID, err)
			}
		}
	}

	// Completed status: stock was already deducted at checkout — no action needed

	if err := uc.orderRepo.UpdateStatus(ctx, orderID, req.Status); err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
}

// validateStatusTransition enforces the order state machine rules.
// Valid transitions: pending → completed | cancelled
// Terminal states (completed, cancelled) cannot be changed.
func validateStatusTransition(current, next domain.OrderStatus) error {
	if current == next {
		return nil
	}

	validTransitions := map[domain.OrderStatus][]domain.OrderStatus{
		domain.OrderStatusPending:   {domain.OrderStatusCompleted, domain.OrderStatusCancelled},
		domain.OrderStatusCompleted: {}, // terminal state
		domain.OrderStatusCancelled: {}, // terminal state
	}

	allowed, ok := validTransitions[current]
	if !ok {
		return fmt.Errorf("%w: unknown current status", domain.ErrInvalidOrderStatus)
	}

	for _, s := range allowed {
		if s == next {
			return nil
		}
	}

	return fmt.Errorf("%w: cannot transition from '%s' to '%s'", domain.ErrInvalidOrderStatus, current, next)
}

// mapToResponse converts Order entity to OrderResponse DTO
func (uc *orderUseCase) mapToResponse(order *domain.Order) *domain.OrderResponse {
	items := make([]domain.OrderItemResponse, len(order.Items))
	for i, item := range order.Items {
		items[i] = domain.OrderItemResponse{
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  item.Quantity,
			SubTotal:  item.SubTotal,
		}
	}

	return &domain.OrderResponse{
		ID:              order.ID,
		UserID:          order.UserID,
		Items:           items,
		TotalAmount:     order.TotalAmount,
		Status:          string(order.Status),
		ShippingAddress: order.ShippingAddress,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
}
