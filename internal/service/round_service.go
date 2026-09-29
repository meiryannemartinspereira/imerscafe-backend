package service

import (
	"imerscafe-backend/internal/domain"

	"github.com/google/uuid"
)

type RoundService struct {
	scoreService *ScoreService
}

func NewRoundService(scoreService *ScoreService) *RoundService {
	return &RoundService{
		scoreService: scoreService,
	}
}

func (s *RoundService) CreateRound(
	customer domain.Customer,
	recipe domain.Recipe,
	preparation domain.PreparationResult,
	evaluation domain.SoftSkillEvaluation,
) domain.Round {
	score := s.scoreService.Calculate(
		customer,
		preparation,
		evaluation,
	)

	profile := domain.CustomerProfiles[customer.Type]

	return domain.Round{
		ID:                  uuid.NewString(),
		Customer:            customer,
		CustomerProfile:     profile,
		Recipe:              recipe,
		PreparationResult:   preparation,
		SoftSkillEvaluation: evaluation,
		Score:               score,
	}
}
