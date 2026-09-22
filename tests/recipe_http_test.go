package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"imerscafe-backend/internal/app"
	"imerscafe-backend/internal/domain"
)

func TestCreateRecipe(t *testing.T) {
	mux := app.New()

	coffeeRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Coffee"}`,
	)

	if coffeeRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, coffeeRecorder.Code)
	}

	var coffee domain.Ingredient

	if err := json.NewDecoder(coffeeRecorder.Body).Decode(&coffee); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	milkRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Milk"}`,
	)

	if milkRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, milkRecorder.Code)
	}

	var milk domain.Ingredient

	if err := json.NewDecoder(milkRecorder.Body).Decode(&milk); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	recorder := executeRequest(
		mux,
		http.MethodPost,
		"/recipes",
		`{"name": "Cappuccino", "ingredient_ids": ["`+coffee.ID+`", "`+milk.ID+`"]}`,
	)

	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}
}

func TestGetAllRecipes(t *testing.T) {
	mux := app.New()

	recorder := executeRequest(
		mux,
		http.MethodGet,
		"/recipes",
		"",
	)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestGetRecipeByID(t *testing.T) {
	mux := app.New()

	coffeeRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Coffee"}`,
	)

	var coffee domain.Ingredient

	if err := json.NewDecoder(coffeeRecorder.Body).Decode(&coffee); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	milkRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Milk"}`,
	)

	var milk domain.Ingredient

	if err := json.NewDecoder(milkRecorder.Body).Decode(&milk); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	createRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/recipes",
		`{"name": "Cappuccino", "ingredient_ids": ["`+coffee.ID+`", "`+milk.ID+`"]}`,
	)

	if createRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, createRecorder.Code)
	}

	recipesRecorder := executeRequest(
		mux,
		http.MethodGet,
		"/recipes",
		"",
	)

	var recipes []domain.Recipe

	if err := json.NewDecoder(recipesRecorder.Body).Decode(&recipes); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	recipeID := recipes[0].ID

	recorder := executeRequest(
		mux,
		http.MethodGet,
		"/recipes/"+recipeID,
		"",
	)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	var recipe domain.Recipe

	if err := json.NewDecoder(recorder.Body).Decode(&recipe); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if recipe.ID != recipeID {
		t.Errorf("expected recipe ID %q, got %q", recipeID, recipe.ID)
	}

	if recipe.Name != "Cappuccino" {
		t.Errorf("expected recipe name %q, got %q", "Cappuccino", recipe.Name)
	}

	if len(recipe.Ingredients) != 2 {
		t.Errorf("expected 2 ingredients, got %d", len(recipe.Ingredients))
	}
}

func TestUpdateRecipe(t *testing.T) {
	mux := app.New()

	coffeeRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Coffee"}`,
	)

	var coffee domain.Ingredient

	if err := json.NewDecoder(coffeeRecorder.Body).Decode(&coffee); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	createRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/recipes",
		`{"name": "Coffee", "ingredient_ids": ["`+coffee.ID+`"]}`,
	)

	if createRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, createRecorder.Code)
	}

	recipesRecorder := executeRequest(
		mux,
		http.MethodGet,
		"/recipes",
		"",
	)

	var recipes []domain.Recipe

	if err := json.NewDecoder(recipesRecorder.Body).Decode(&recipes); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	recipeID := recipes[0].ID

	recorder := executeRequest(
		mux,
		http.MethodPatch,
		"/recipes/"+recipeID,
		`{"name": "Coffee Special"}`,
	)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNoContent, recorder.Code)
	}
}

func TestDeleteRecipe(t *testing.T) {
	mux := app.New()

	coffeeRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/ingredients",
		`{"name": "Coffee"}`,
	)

	var coffee domain.Ingredient

	if err := json.NewDecoder(coffeeRecorder.Body).Decode(&coffee); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	createRecorder := executeRequest(
		mux,
		http.MethodPost,
		"/recipes",
		`{"name": "Coffee", "ingredient_ids": ["`+coffee.ID+`"]}`,
	)

	if createRecorder.Code != http.StatusCreated {
		t.Fatalf("expected status code %d, got %d", http.StatusCreated, createRecorder.Code)
	}

	recipesRecorder := executeRequest(
		mux,
		http.MethodGet,
		"/recipes",
		"",
	)

	var recipes []domain.Recipe

	if err := json.NewDecoder(recipesRecorder.Body).Decode(&recipes); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	recipeID := recipes[0].ID

	recorder := executeRequest(
		mux,
		http.MethodDelete,
		"/recipes/"+recipeID,
		"",
	)

	if recorder.Code != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNoContent, recorder.Code)
	}
}
