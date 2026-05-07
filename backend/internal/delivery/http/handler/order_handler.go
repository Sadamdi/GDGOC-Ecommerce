package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	deliveryHttp "ecommerce-backend/internal/delivery/http"
	"ecommerce-backend/internal/delivery/http/middleware"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/pkg/validator"
)

type OrderHandler struct {
	orderUC domain.OrderUseCase
}

func NewOrderHandler(orderUC domain.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUC: orderUC,
	}
}

// Checkout godoc
// @Summary      Create a new order (Checkout)
// @Description  Creates an order from the current user's shopping cart. Will fail if cart is empty or stock is insufficient.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body domain.CreateOrderRequest true "Shipping Address details"
// @Success      201 {object} deliveryHttp.Response
// @Failure      400 {object} deliveryHttp.Response
// @Failure      401 {object} deliveryHttp.Response
// @Failure      500 {object} deliveryHttp.Response
// @Router       /orders [post]
func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "UNAUTHORIZED", nil)
		return
	}

	var req domain.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	orderResp, err := h.orderUC.Checkout(r.Context(), userID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrEmptyCart) || errors.Is(err, domain.ErrInsufficientStock) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "BAD_REQUEST", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusCreated, "Order created successfully", orderResp)
}

// GetMyOrders godoc
// @Summary      List user orders
// @Description  Get a paginated list of orders for the authenticated user
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        per_page query int false "Items per page" default(20)
// @Success      200 {object} deliveryHttp.Response
// @Failure      401 {object} deliveryHttp.Response
// @Failure      500 {object} deliveryHttp.Response
// @Router       /orders [get]
func (h *OrderHandler) GetMyOrders(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "UNAUTHORIZED", nil)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("per_page"))

	paginatedResp, err := h.orderUC.GetMyOrders(r.Context(), userID, page, limit)
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "INTERNAL_ERROR", nil)
		return
	}

	meta := deliveryHttp.Meta{
		Page:       paginatedResp.Page,
		PerPage:    paginatedResp.PerPage,
		Total:      int(paginatedResp.Total),
		TotalPages: paginatedResp.TotalPages,
	}

	deliveryHttp.SuccessList(w, http.StatusOK, "Orders retrieved successfully", paginatedResp.Data, meta)
}

// GetOrderByID godoc
// @Summary      Get order details
// @Description  Get specific order details by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Order ID"
// @Success      200 {object} deliveryHttp.Response
// @Failure      401 {object} deliveryHttp.Response
// @Failure      403 {object} deliveryHttp.Response
// @Failure      404 {object} deliveryHttp.Response
// @Failure      500 {object} deliveryHttp.Response
// @Router       /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "UNAUTHORIZED", nil)
		return
	}
	isAdmin := middleware.IsAdmin(r.Context())
	orderID := r.PathValue("id")

	orderResp, err := h.orderUC.GetOrderByID(r.Context(), orderID, userID, isAdmin)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, "Order not found", "NOT_FOUND", nil)
			return
		}
		if errors.Is(err, domain.ErrNotOrderOwner) {
			deliveryHttp.ErrorResponse(w, http.StatusForbidden, "Forbidden", "FORBIDDEN", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Order retrieved successfully", orderResp)
}

// CancelOrder godoc
// @Summary      Cancel an order
// @Description  Cancel an order that is still in pending status
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Order ID"
// @Success      200 {object} deliveryHttp.Response
// @Failure      400 {object} deliveryHttp.Response
// @Failure      401 {object} deliveryHttp.Response
// @Failure      403 {object} deliveryHttp.Response
// @Failure      404 {object} deliveryHttp.Response
// @Failure      500 {object} deliveryHttp.Response
// @Router       /orders/{id}/cancel [put]
func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserID(r.Context())
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "UNAUTHORIZED", nil)
		return
	}

	orderID := r.PathValue("id")

	if err := h.orderUC.CancelOrder(r.Context(), orderID, userID); err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, "Order not found", "NOT_FOUND", nil)
			return
		}
		if errors.Is(err, domain.ErrNotOrderOwner) {
			deliveryHttp.ErrorResponse(w, http.StatusForbidden, "Forbidden", "FORBIDDEN", nil)
			return
		}
		if errors.Is(err, domain.ErrInvalidOrderStatus) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "BAD_REQUEST", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Order cancelled successfully", nil)
}

// UpdateStatus godoc
// @Summary      Update order status
// @Description  Update the status of an order (Admin only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Order ID"
// @Param        request body domain.UpdateOrderStatusRequest true "Status details"
// @Success      200 {object} deliveryHttp.Response
// @Failure      400 {object} deliveryHttp.Response
// @Failure      401 {object} deliveryHttp.Response
// @Failure      403 {object} deliveryHttp.Response
// @Failure      404 {object} deliveryHttp.Response
// @Failure      500 {object} deliveryHttp.Response
// @Router       /admin/orders/{id}/status [put]
func (h *OrderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	if !middleware.IsAdmin(r.Context()) {
		deliveryHttp.ErrorResponse(w, http.StatusForbidden, "Admin access required", "FORBIDDEN", nil)
		return
	}

	orderID := r.PathValue("id")

	var req domain.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	if err := h.orderUC.UpdateOrderStatus(r.Context(), orderID, &req); err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, "Order not found", "NOT_FOUND", nil)
			return
		}
		if errors.Is(err, domain.ErrInvalidOrderStatus) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "BAD_REQUEST", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, err.Error(), "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Order status updated successfully", nil)
}
