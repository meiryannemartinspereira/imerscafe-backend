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

func executeRequest(mux http.Handler, method string, path string, body string) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)

	return recorder
}

func TestCreateIngredient(t *testing.T) {
	mux := app.New()

	recorder := executeRequest(mux, http.MethodPost, "/ingredients", `{"name": "Sugar"}`)

	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}

	var ingredient domain.Ingredient
	if err := json.NewDecoder(recorder.Body).Decode(&ingredient); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if ingredient.Name != "Sugar" {
		t.Errorf("expected ingredient name 'Sugar', got '%s'", ingredient.Name)
	}
	if ingredient.ID == "" {
		t.Error("expected ingredient ID to be set, got empty string")
	}
}

func TestGetAllIngredients(t *testing.T) {
	mux := app.New()

	recorder := executeRequest(
		mux,
		http.MethodGet,
		"/ingredients",
		"",
	)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestGetIngredientByID(t *testing.T) {
	mux := app.New()

	createRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Sugar"}`,
	)

	if createRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, createRecorder.Code)
	}

	var createdIngredient domain.Ingredient

	err := json.NewDecoder(createRecorder.Body).Decode(&createdIngredient)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	recorder := executeRequest(
		mux,
		http.MethodGet,
		"/ingredients/"+createdIngredient.ID,
		"",
	)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	var ingredient domain.Ingredient

	err = json.NewDecoder(recorder.Body).Decode(&ingredient)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if ingredient.ID != createdIngredient.ID {
		t.Errorf("expected ingredient ID %q, got %q", createdIngredient.ID, ingredient.ID)
	}

	if ingredient.Name != "Sugar" {
		t.Errorf("expected ingredient name %q, got %q", "Sugar", ingredient.Name)
	}
}
