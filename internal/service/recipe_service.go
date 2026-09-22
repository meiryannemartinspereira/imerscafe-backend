package service

import (
	"context"
	"errors"
	"strings"

	"imerscafe-backend/internal/domain"
	"imerscafe-backend/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrInvalidRecipeName         = errors.New("recipe name is required")
	ErrDuplicateRecipe           = errors.New("recipe already exists")
	ErrRecipeWithoutIngredients  = errors.New("recipe must have at least one ingredient")
	ErrDuplicateRecipeIngredient = errors.New("ingredient is duplicated in recipe")
)

type RecipeService interface {
	Create(ctx context.Context, recipe domain.Recipe) error
	GetAll(ctx context.Context) ([]domain.Recipe, error)
	GetByID(ctx context.Context, id string) (domain.Recipe, error)
	Update(ctx context.Context, recipe domain.Recipe) error
	Delete(ctx context.Context, id string) error
}

type recipeService struct {
	recipeRepository     repository.RecipeRepository
	ingredientRepository repository.IngredientRepository
}

func NewRecipeService(
	recipeRepository repository.RecipeRepository,
	ingredientRepository repository.IngredientRepository,
) *recipeService {
	return &recipeService{
		recipeRepository:     recipeRepository,
		ingredientRepository: ingredientRepository,
	}
}

func (s *recipeService) Create(ctx context.Context, recipe domain.Recipe) error {
	recipe.Name = strings.TrimSpace(recipe.Name)

	if recipe.Name == "" {
		return ErrInvalidRecipeName
	}

	recipes, err := s.recipeRepository.GetAll(ctx)

	if err != nil {
		return err
	}

	for _, existingRecipe := range recipes {
		if existingRecipe.Name == recipe.Name {
			return ErrDuplicateRecipe
		}
	}

	if len(recipe.Ingredients) == 0 {
		return ErrRecipeWithoutIngredients
	}

	ingredients := make(map[string]bool)

	for _, ingredient := range recipe.Ingredients {
		if ingredients[ingredient.ID] {
			return ErrDuplicateRecipeIngredient
		}

		_, err := s.ingredientRepository.GetByID(ctx, ingredient.ID)

		if err != nil {
			return err
		}

		ingredients[ingredient.ID] = true
	}

	recipe.ID = uuid.New().String()

	if err := s.recipeRepository.Create(ctx, &recipe); err != nil {
		return err
	}

	return nil
}
