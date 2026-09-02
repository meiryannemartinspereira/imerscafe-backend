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
	ErrInvalidIngredientName = errors.New("nome do ingrediente é obrigatório")
	ErrDuplicateIngredient   = errors.New("já existe um ingrediente com este nome")
)

type IngredientService struct {
	repo repository.IngredientRepository
}

func NewIngredientService(repo repository.IngredientRepository) *IngredientService {
	return &IngredientService{repo: repo}
}

func (s *IngredientService) Create(ctx context.Context, name string) (*domain.Ingredient, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, ErrInvalidIngredientName
	}

	ingredients, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, ing := range ingredients {
		if strings.EqualFold(ing.Name, trimmedName) {
			return nil, ErrDuplicateIngredient
		}
	}

	newIngredient := &domain.Ingredient{
		ID:   uuid.New().String(),
		Name: trimmedName,
	}

	if err := s.repo.Create(ctx, newIngredient); err != nil {
		return nil, err
	}

	return newIngredient, nil
}

func (s *IngredientService) GetAll(ctx context.Context) ([]*domain.Ingredient, error) {
	return s.repo.GetAll(ctx)
}

func (s *IngredientService) GetByID(ctx context.Context, id string) (*domain.Ingredient, error) {
	return s.repo.GetByID(ctx, id)
}