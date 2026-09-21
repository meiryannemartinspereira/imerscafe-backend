package repository

import (
	"context"
	"imerscafe-backend/internal/domain"
)

type InMemoryRecipeRepository struct {
	recipes map[string]domain.Recipe
}

func NewInMemoryRecipeRepository() *InMemoryRecipeRepository {
	return &InMemoryRecipeRepository{
		recipes: make(map[string]domain.Recipe),
	}
}

func (r *InMemoryRecipeRepository) Create(ctx context.Context, recipe domain.Recipe) error {
	r.recipes[recipe.ID] = recipe

	return nil
}
