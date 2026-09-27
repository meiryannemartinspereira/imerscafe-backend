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

func TestScoreService_ZeroPreparationScoreWhenIncorrect(t *testing.T) {
	customer := domain.Customer{
		Type: domain.CustomerIndecisive,
	}

	preparation := domain.PreparationResult{
		Correct: false,
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

	if score.Total != 60 {
		t.Fatalf("expected score 60, got %d", score.Total)
	}
}

func TestScoreService_ZeroSoftSkillScoreWhenEvaluationIsZero(t *testing.T) {
	customer := domain.Customer{
		Type: domain.CustomerCalm,
	}

	preparation := domain.PreparationResult{
		Correct: true,
	}

	evaluation := domain.SoftSkillEvaluation{}

	scoreService := NewScoreService()

	score := scoreService.Calculate(
		customer,
		preparation,
		evaluation,
	)

	if score.Total != 45 {
		t.Fatalf("expected score 45, got %d", score.Total)
	}
}
