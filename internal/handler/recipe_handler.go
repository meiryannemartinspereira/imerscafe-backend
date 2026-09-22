package handler

import (
	"encoding/json"
	"net/http"

	"imerscafe-backend/internal/service"
)

type RecipeHandler struct {
	service service.RecipeService
}

func NewRecipeHandler(service service.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		service: service,
	}
}

type createRecipeRequest struct {
	Name          string   `json:"name"`
	IngredientIDs []string `json:"ingredient_ids"`
}

type updateRecipeRequest struct {
	Name          *string   `json:"name"`
	IngredientIDs *[]string `json:"ingredient_ids"`
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request createRecipeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input := service.CreateRecipeInput{
		Name:          request.Name,
		IngredientIDs: request.IngredientIDs,
	}

	if err := h.service.Create(r.Context(), input); err != nil {
		switch err {
		case service.ErrInvalidRecipeName:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrDuplicateRecipe:
			http.Error(w, err.Error(), http.StatusConflict)
		case service.ErrRecipeWithoutIngredients:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrDuplicateRecipeIngredient:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *RecipeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (h *RecipeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	recipe, err := h.service.GetByID(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(recipe)
}

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request updateRecipeRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input := service.UpdateRecipeInput{
		ID:            r.PathValue("id"),
		Name:          request.Name,
		IngredientIDs: request.IngredientIDs,
	}

	if err := h.service.Update(r.Context(), input); err != nil {
		switch err {
		case service.ErrInvalidRecipeName:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrDuplicateRecipe:
			http.Error(w, err.Error(), http.StatusConflict)
		case service.ErrRecipeWithoutIngredients:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case service.ErrDuplicateRecipeIngredient:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
