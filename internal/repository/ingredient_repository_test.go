package repository

import (
	"context"
	"testing"

	"imerscafe-backend/internal/domain"
)

func TestInMemoryIngredientRepository_CreateAndGetByID(t *testing.T) {
	repo := NewInMemoryIngredientRepository()

	ingredient := domain.Ingredient{
		ID:   "1",
		Name: "CAFE",
	}

	err := repo.Create(context.Background(), ingredient)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result, err := repo.GetByID(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != ingredient {
		t.Fatalf("expected %v, got %v", ingredient, result)
	}
}
