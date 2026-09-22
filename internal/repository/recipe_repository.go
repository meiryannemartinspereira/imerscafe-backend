package repository

import (
	"context"
	"imerscafe-backend/internal/domain"
)

type RecipeRepository interface {
	Create(ctx context.Context, recipe domain.Recipe) error
	GetAll(ctx context.Context) ([]domain.Recipe, error)
	GetByID(ctx context.Context, id string) (domain.Recipe, error)
	Update(ctx context.Context, recipe domain.Recipe) error
	Delete(ctx context.Context, id string) error
}
