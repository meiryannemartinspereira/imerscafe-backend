package repository

import (
	"context"
	"errors"

	"imerscafe-backend/internal/domain"
)

var ErrIngredientNotFound = errors.New("ingredient not found")

type InMemoryIngredientRepository struct {
	ingredients map[string]domain.Ingredient
}

func NewInMemoryIngredientRepository() *InMemoryIngredientRepository {
	return &InMemoryIngredientRepository{
		ingredients: make(map[string]domain.Ingredient),
	}
}

func (r *InMemoryIngredientRepository) Create(
	ctx context.Context,
	ingredient domain.Ingredient,
) error {
	r.ingredients[ingredient.ID] = ingredient

	return nil
}

func (r *InMemoryIngredientRepository) GetAll(
	ctx context.Context,
) ([]domain.Ingredient, error) {
	ingredients := make([]domain.Ingredient, 0, len(r.ingredients))

	for _, ingredient := range r.ingredients {
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (r *InMemoryIngredientRepository) GetByID(
	ctx context.Context,
	id string,
) (domain.Ingredient, error) {
	ingredient, exists := r.ingredients[id]

	if !exists {
		return domain.Ingredient{}, ErrIngredientNotFound
	}

	return ingredient, nil
}
