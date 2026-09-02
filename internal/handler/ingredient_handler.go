package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"imerscafe-backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type IngredientHandler struct {
	service *service.IngredientService
}

func NewIngredientHandler(service *service.IngredientService) *IngredientHandler {
	return &IngredientHandler{service: service}
}

type createIngredientRequest struct {
	Name string `json:"name"`
}

func (h *IngredientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createIngredientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "payload inválido", http.StatusBadRequest)
		return
	}

	ingredient, err := h.service.Create(r.Context(), req.Name)
	if err != nil {
		if errors.Is(err, service.ErrInvalidIngredientName) || errors.Is(err, service.ErrDuplicateIngredient) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "erro interno no servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ingredient)
}

func (h *IngredientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "erro interno no servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ingredients)
}

func (h *IngredientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	ingredient, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "erro interno no servidor", http.StatusInternalServerError)
		return
	}

	if ingredient == nil {
		http.Error(w, "ingrediente não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ingredient)
}