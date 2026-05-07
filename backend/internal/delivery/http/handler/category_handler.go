package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	deliveryHttp "ecommerce-backend/internal/delivery/http"
	"ecommerce-backend/internal/domain"
	"ecommerce-backend/internal/pkg/validator"
)

type CategoryHandler struct {
	categoryUseCase domain.CategoryUseCase
}

func NewCategoryHandler(uc domain.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: uc}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new product category. Requires admin role.
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body domain.CategoryRequest true "Category Request"
// @Success 201 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Failure 403 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req domain.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	cat, err := h.categoryUseCase.CreateCategory(r.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryAlreadyExists) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "CATEGORY_EXISTS", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to create category", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusCreated, "Category created successfully", cat)
}

// GetAllCategories godoc
// @Summary List all categories
// @Description Get a list of all product categories
// @Tags categories
// @Produce json
// @Success 200 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryUseCase.GetAllCategories(r.Context())
	if err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve categories", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Categories retrieved successfully", categories)
}

// GetCategoryByID godoc
// @Summary Get category details
// @Description Get a product category by ID
// @Tags categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 404 {object} deliveryHttp.Response
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid category ID", "BAD_REQUEST", nil)
		return
	}

	category, err := h.categoryUseCase.GetCategoryByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, err.Error(), "NOT_FOUND", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve category", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Category retrieved successfully", category)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing product category. Requires admin role.
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Category ID"
// @Param request body domain.CategoryRequest true "Category Request"
// @Success 200 {object} deliveryHttp.Response
// @Failure 400 {object} deliveryHttp.Response
// @Failure 401 {object} deliveryHttp.Response
// @Failure 403 {object} deliveryHttp.Response
// @Failure 404 {object} deliveryHttp.Response
// @Failure 500 {object} deliveryHttp.Response
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid category ID", "BAD_REQUEST", nil)
		return
	}

	var req domain.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", "BAD_REQUEST", nil)
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		deliveryHttp.ErrorResponse(w, http.StatusBadRequest, "Validation error", "VALIDATION_ERROR", err.Error())
		return
	}

	cat, err := h.categoryUseCase.UpdateCategory(r.Context(), id, &req)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryNotFound) {
			deliveryHttp.ErrorResponse(w, http.StatusNotFound, err.Error(), "NOT_FOUND", nil)
			return
		}
		if errors.Is(err, domain.ErrCategoryAlreadyExists) {
			deliveryHttp.ErrorResponse(w, http.StatusBadRequest, err.Error(), "CATEGORY_EXISTS", nil)
			return
		}
		deliveryHttp.ErrorResponse(w, http.StatusInternalServerError, "Failed to update category", "INTERNAL_ERROR", nil)
		return
	}

	deliveryHttp.Success(w, http.StatusOK, "Category updated successfully", cat)
}
