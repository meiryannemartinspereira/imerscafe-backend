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

func TestInMemoryRecipeRepositoryGetAll(t *testing.T) {
	repository := NewInMemoryRecipeRepository()

	repository.Create(context.Background(), domain.Recipe{
		ID:   "1",
		Name: "Cappuccino",
	})

	repository.Create(context.Background(), domain.Recipe{
		ID:   "2",
		Name: "Latte",
	})

	recipes, err := repository.GetAll(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if len(recipes) != 2 {
		t.Fatalf("expected 2 recipes, got %d", len(recipes))
	}
}

func TestInMemoryRecipeRepositoryGetByID(t *testing.T) {
	repository := NewInMemoryRecipeRepository()

	recipe := domain.Recipe{
		ID:   "1",
		Name: "Cappuccino",
	}

	repository.Create(context.Background(), recipe)

	result, err := repository.GetByID(context.Background(), "1")

	if err != nil {
		t.Fatal(err)
	}

	if result.ID != "1" {
		t.Fatalf("expected ID 1, got %s", result.ID)
	}
}

func TestInMemoryRecipeRepositoryUpdate(t *testing.T) {
	repository := NewInMemoryRecipeRepository()

	repository.Create(context.Background(), domain.Recipe{
		ID:   "1",
		Name: "Cappuccino",
	})

	err := repository.Update(context.Background(), domain.Recipe{
		ID:   "1",
		Name: "Cappuccino Especial",
	})

	if err != nil {
		t.Fatal(err)
	}

	result, _ := repository.GetByID(context.Background(), "1")

	if result.Name != "Cappuccino Especial" {
		t.Fatalf("expected name Cappuccino Especial, got %s", result.Name)
	}
}

func TestInMemoryRecipeRepositoryDelete(t *testing.T) {
	repository := NewInMemoryRecipeRepository()

	repository.Create(context.Background(), domain.Recipe{
		ID:   "1",
		Name: "Cappuccino",
	})

	err := repository.Delete(context.Background(), "1")

	if err != nil {
		t.Fatal(err)
	}

	_, err = repository.GetByID(context.Background(), "1")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
