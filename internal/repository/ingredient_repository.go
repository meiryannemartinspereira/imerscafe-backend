package repository

import (
	"context"
	"imerscafe-backend/internal/domain"
)

type IngredientRepository interface {
	Create(ctx context.Context, ingredient *domain.Ingredient) error
	GetAll(ctx context.Context) ([]*domain.Ingredient, error)
	GetByID(ctx context.Context, id string) (*domain.Ingredient, error)
}