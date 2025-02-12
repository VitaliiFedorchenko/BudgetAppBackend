package handlers

import (
	"BudgetApp/internal/services"
	"BudgetApp/internal/utils"
	"encoding/json"
	"net/http"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func SetupCategoryHandler() *CategoryHandler {
	categoryService := services.NewCategoryService()
	return NewCategoryHandler(categoryService)
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	utils.NewResponse(w).ResponseJSON(categories, http.StatusOK)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.categoryService.CreateCategory(req.Name); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	utils.NewResponse(w).ResponseJSON("Category created successfully", http.StatusCreated)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Slug string `json:"slug"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.categoryService.DeleteCategory(req.Slug); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	utils.NewResponse(w).ResponseJSON("Category deleted successfully", http.StatusOK)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.NewResponse(w).ResponseJSON("Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Slug    string `json:"slug"`
		NewName string `json:"new_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.categoryService.UpdateCategory(req.Slug, req.NewName); err != nil {
		utils.NewResponse(w).ResponseJSON(err.Error(), http.StatusInternalServerError)
		return
	}

	utils.NewResponse(w).ResponseJSON("Category updated successfully", http.StatusOK)
}
