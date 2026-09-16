package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"imerscafe-backend/internal/app"
	"imerscafe-backend/internal/domain"
)

func TestCreateIngredient(t *testing.T) {
	mux := app.New()

	request := httptest.NewRequest(
		http.MethodPost,
		"/ingredients",
		strings.NewReader(`{"name": "Sugar"}`),
	)

	recorder := httptest.NewRecorder()

	mux.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}

	var ingredient domain.Ingredient

	err := json.NewDecoder(recorder.Body).Decode(&ingredient)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if ingredient.Name != "Sugar" {
		t.Errorf("expected ingredient name %q, got %q", "Sugar", ingredient.Name)
	}

	if ingredient.ID == "" {
		t.Errorf("expected ingredient ID to be non-empty")
	}
}
