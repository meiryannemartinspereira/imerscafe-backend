package service

import (
	"imerscafe-backend/internal/ai"
	"imerscafe-backend/internal/domain"

	"github.com/google/uuid"
)

type RoundService struct {
	scoreService *ScoreService
	aiService    ai.AIService
}

func NewRoundService(scoreService *ScoreService, aiService ai.AIService) *RoundService {
	return &RoundService{
		scoreService: scoreService,
		aiService:    aiService,
	}
}

func (s *RoundService) CreateRound(
	customer domain.Customer,
	recipe domain.Recipe,
	preparation domain.PreparationResult,
	evaluation domain.SoftSkillEvaluation,
) (domain.Round, ai.AIResponse) {

	profile := domain.CustomerProfiles[customer.Type]

	request := ai.AIRequest{
		CustomerType: string(profile.Type),
		Behavior:     profile.Behavior,
		RecipeName:   recipe.Name,
	}

	response := s.aiService.SimulateCustomer(request)

	score := s.scoreService.Calculate(
		customer,
		preparation,
		evaluation,
	)

	round := domain.Round{
		ID:                  uuid.NewString(),
		Customer:            customer,
		CustomerProfile:     profile,
		Recipe:              recipe,
		PreparationResult:   preparation,
		SoftSkillEvaluation: evaluation,
		Score:               score,
	}

	return round, response
}
