package repository

import "imerscafe-backend/internal/domain"

type InMemoryIngredientRepository struct {
	ingredients map[string]domain.Ingredient
}
