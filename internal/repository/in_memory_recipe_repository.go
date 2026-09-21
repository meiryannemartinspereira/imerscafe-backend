package repository

import (
	"context"
	"errors"
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

func (r *InMemoryRecipeRepository) GetAll(ctx context.Context) ([]domain.Recipe, error) {
	recipes := []domain.Recipe{}

	for _, recipe := range r.recipes {
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *InMemoryRecipeRepository) GetByID(ctx context.Context, id string) (domain.Recipe, error) {
	recipe, ok := r.recipes[id]

	if !ok {
		return domain.Recipe{}, errors.New("Recipe not found")
	}

	return recipe, nil
}

func (r *InMemoryRecipeRepository) Update(ctx context.Context, recipe domain.Recipe) error {
	_, ok := r.recipes[recipe.ID]

	if !ok {
		return errors.New("recipe not found")
	}

	r.recipes[recipe.ID] = recipe

	return nil
}

func (r *InMemoryRecipeRepository) Delete(ctx context.Context, id string) error {
	_, ok := r.recipes[id]

	if !ok {
		return errors.New("recipe not found")
	}

	delete(r.recipes, id)

	return nil
}
