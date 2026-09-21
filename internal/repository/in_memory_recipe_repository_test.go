package repository

import (
	"context"
	"testing"

	"imerscafe-backend/internal/domain"
)

func TestInMemoryRecipeRepositoryCreate(t *testing.T) {
	repository := NewInMemoryRecipeRepository()

	recipe := domain.Recipe{
		ID:   "1",
		Name: "Cappuccino",
	}

	err := repository.Create(context.Background(), recipe)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	result, ok := repository.recipes["1"]

	if !ok {
		t.Fatalf("expected recipe to exist")
	}

	if result.Name != "Cappuccino" {
		t.Errorf("expected name Cappuccino, got %s", result.Name)
	}
}
