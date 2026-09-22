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

type CreateRecipeInput struct {
	Name          string
	IngredientIDs []string
}

type UpdateRecipeInput struct {
	ID            string
	Name          *string
	IngredientIDs *[]string
}

type RecipeService interface {
	Create(ctx context.Context, input CreateRecipeInput) error
	GetAll(ctx context.Context) ([]domain.Recipe, error)
	GetByID(ctx context.Context, id string) (domain.Recipe, error)
	Update(ctx context.Context, input UpdateRecipeInput) error
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

func (s *recipeService) Create(ctx context.Context, input CreateRecipeInput) error {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return ErrInvalidRecipeName
	}

	recipes, err := s.recipeRepository.GetAll(ctx)

	if err != nil {
		return err
	}

	for _, existingRecipe := range recipes {
		if existingRecipe.Name == input.Name {
			return ErrDuplicateRecipe
		}
	}

	if len(input.IngredientIDs) == 0 {
		return ErrRecipeWithoutIngredients
	}

	ingredients := []domain.Ingredient{}
	ingredientIDs := make(map[string]bool)

	for _, ingredientID := range input.IngredientIDs {
		if ingredientIDs[ingredientID] {
			return ErrDuplicateRecipeIngredient
		}

		ingredient, err := s.ingredientRepository.GetByID(ctx, ingredientID)

		if err != nil {
			return err
		}

		ingredients = append(ingredients, ingredient)
		ingredientIDs[ingredientID] = true
	}

	recipe := domain.Recipe{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Ingredients: ingredients,
	}

	if err := s.recipeRepository.Create(ctx, recipe); err != nil {
		return err
	}

	return nil
}

func (s *recipeService) GetAll(ctx context.Context) ([]domain.Recipe, error) {
	return s.recipeRepository.GetAll(ctx)
}

func (s *recipeService) GetByID(ctx context.Context, id string) (domain.Recipe, error) {
	return s.recipeRepository.GetByID(ctx, id)
}

func (s *recipeService) Delete(ctx context.Context, id string) error {
	return s.recipeRepository.Delete(ctx, id)
}

func (s *recipeService) Update(ctx context.Context, input UpdateRecipeInput) error {
	existingRecipe, err := s.recipeRepository.GetByID(ctx, input.ID)

	if err != nil {
		return err
	}

	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)

		if name == "" {
			return ErrInvalidRecipeName
		}

		recipes, err := s.recipeRepository.GetAll(ctx)

		if err != nil {
			return err
		}

		for _, existing := range recipes {
			if existing.ID != input.ID && existing.Name == name {
				return ErrDuplicateRecipe
			}
		}

		existingRecipe.Name = name
	}

	if input.IngredientIDs != nil {
		if len(*input.IngredientIDs) == 0 {
			return ErrRecipeWithoutIngredients
		}

		ingredients := []domain.Ingredient{}
		ingredientIDs := make(map[string]bool)

		for _, ingredientID := range *input.IngredientIDs {
			if ingredientIDs[ingredientID] {
				return ErrDuplicateRecipeIngredient
			}

			ingredient, err := s.ingredientRepository.GetByID(ctx, ingredientID)

			if err != nil {
				return err
			}

			ingredients = append(ingredients, ingredient)
			ingredientIDs[ingredientID] = true
		}

		existingRecipe.Ingredients = ingredients
	}

	return s.recipeRepository.Update(ctx, existingRecipe)
}
