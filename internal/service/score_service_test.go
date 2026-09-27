package service

import (
	"imerscafe-backend/internal/domain"
	"testing"
)

func TestScoreService_CalculatesMaximumScore(t *testing.T) {
	customer := domain.Customer{
		Type: domain.CustomerIndecisive,
	}

	preparation := domain.PreparationResult{
		Correct: true,
	}

	evaluation := domain.SoftSkillEvaluation{
		Communication: 10,
		Empathy:       10,
		Politeness:    10,
		Clarity:       10,
	}

	scoreService := NewScoreService()

	score := scoreService.Calculate(
		customer,
		preparation,
		evaluation,
	)

	if score.Total != 100 {
		t.Fatalf("expected score 100, got %d", score.Total)
	}
}
