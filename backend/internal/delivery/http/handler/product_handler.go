package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	deliveryHttp "ecommerce-backend/internal/delivery/http"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/pkg/validator"
)

type ProductHandler struct {
	productUseCase domain.ProductUseCase
}

func NewProductHandler(uc domain.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: uc}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product in the catalog. Requires admin role.
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.ProductRequest true "Product Request"
// @Success 201 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Failure 403 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req domain.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	prod, err := h.productUseCase.CreateProduct(r.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "INVALID_CATEGORY", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to create product", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusCreated, "Product created successfully", prod)
}

// GetProductByID godoc
// @Summary Get product details
// @Description Get a product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 404 {object} deliveryHttp.Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "BAD_REQUEST", nil)
		return
	}

	prod, err := h.productUseCase.GetProductByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, err.Error(), "NOT_FOUND", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve product", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Product retrieved successfully", prod)
}

// GetAllProducts godoc
// @Summary List products
// @Description Browse and search products with optional filters and pagination
// @Tags products
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Param category query string false "Category ID"
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param in_stock query bool false "Filter by stock availability"
// @Param q query string false "Search keyword"
// @Success 200 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	query := domain.ProductQuery{}

	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			query.Page = parsed
		} else {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid page format", "BAD_REQUEST", nil)
			return
		}
	}
	if p := r.URL.Query().Get("per_page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			query.PerPage = parsed
		} else {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid per_page format", "BAD_REQUEST", nil)
			return
		}
	}
	query.CategoryID = r.URL.Query().Get("category")
	query.Search = r.URL.Query().Get("q")

	if min := r.URL.Query().Get("min_price"); min != "" {
		if parsed, err := strconv.ParseFloat(min, 64); err == nil {
			query.MinPrice = parsed
		} else {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid min_price format", "BAD_REQUEST", nil)
			return
		}
	}
	if max := r.URL.Query().Get("max_price"); max != "" {
		if parsed, err := strconv.ParseFloat(max, 64); err == nil {
			query.MaxPrice = parsed
		} else {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid max_price format", "BAD_REQUEST", nil)
			return
		}
	}
	if stock := r.URL.Query().Get("in_stock"); stock != "" {
		if parsed, err := strconv.ParseBool(stock); err == nil {
			query.InStock = &parsed
		} else {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid in_stock format", "BAD_REQUEST", nil)
			return
		}
	}

	resp, err := h.productUseCase.GetAllProducts(r.Context(), query)
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve products", "INTERNAL_ERROR", nil)
		return
	}

	meta := deliveryHttp.Meta{
		Page:       resp.Page,
		PerPage:    resp.PerPage,
		Total:      int(resp.Total),
		TotalPages: resp.TotalPages,
	}

	deliveryHttp.SuccessList(w, http.StatusOK, "Products retrieved successfully", resp.Data, meta)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product. Requires admin role.
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Param request body domain.ProductRequest true "Product Request"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Failure 403 {object} deliveryHttp.Response
// @Failure 404 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "BAD_REQUEST", nil)
		return
	}

	var req domain.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	prod, err := h.productUseCase.UpdateProduct(r.Context(), id, &req)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, err.Error(), "NOT_FOUND", nil)
			return
		}
		if errors.Is(err, domain.ErrCategoryNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "INVALID_CATEGORY", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to update product", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Product updated successfully", prod)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product. Requires admin role.
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Failure 403 {object} deliveryHttp.Response
// @Failure 404 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid product ID", "BAD_REQUEST", nil)
		return
	}

	err := h.productUseCase.DeleteProduct(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, err.Error(), "NOT_FOUND", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete product", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Product deleted successfully", nil)
}
