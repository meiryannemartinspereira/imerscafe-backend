package service

import (
	"imerscafe-backend/internal/domain"
	"testing"
)

func TestRoundServiceCreateRound(t *testing.T) {
	scoreService := NewScoreService()
	roundService := NewRoundService(scoreService)

	customer := domain.Customer{
		ID:   "customer-1",
		Type: domain.CustomerCalm,
	}

	recipe := domain.Recipe{
		ID:   "recipe-1",
		Name: "Cappuccino",
	}

	preparation := domain.PreparationResult{
		Correct: true,
	}

	evaluation := domain.SoftSkillEvaluation{
		Communication: 8,
		Empathy:       8,
		Politeness:    8,
		Clarity:       8,
		Feedback:      "Bom atendimento",
	}

	round := roundService.CreateRound(
		customer,
		recipe,
		preparation,
		evaluation,
	)

	if round.ID == "" {
		t.Error("expected round ID to be generated")
	}

	if round.Customer != customer {
		t.Error("expected customer to be the same")
	}

	if round.CustomerProfile.Type != domain.CustomerCalm {
		t.Errorf(
			"expected customer profile type %s, got %s",
			domain.CustomerCalm,
			round.CustomerProfile.Type,
		)
	}

	if round.Recipe.ID != recipe.ID {
		t.Error("expected recipe ID to be the same")
	}

	if round.Recipe.Name != recipe.Name {
		t.Error("expected recipe name to be the same")
	}

	if round.PreparationResult != preparation {
		t.Error("expected preparation result to be the same")
	}

	if round.SoftSkillEvaluation != evaluation {
		t.Error("expected soft skill evaluation to be the same")
	}

	if round.Score.Total != 69 {
		t.Errorf("expected score 69, got %d", round.Score.Total)
	}
}
func TestRoundServiceCreateRoundWithIncorrectPreparation(t *testing.T) {
	scoreService := NewScoreService()
	roundService := NewRoundService(scoreService)

	customer := domain.Customer{
		ID:   "customer-1",
		Type: domain.CustomerCalm,
	}

	recipe := domain.Recipe{
		ID:   "recipe-1",
		Name: "Cappuccino",
	}

	preparation := domain.PreparationResult{
		Correct: false,
	}

	evaluation := domain.SoftSkillEvaluation{
		Communication: 8,
		Empathy:       8,
		Politeness:    8,
		Clarity:       8,
		Feedback:      "Bom atendimento",
	}

	round := roundService.CreateRound(
		customer,
		recipe,
		preparation,
		evaluation,
	)

	if round.PreparationResult.Correct {
		t.Error("expected preparation to be incorrect")
	}

	if round.Score.Total != 29 {
		t.Errorf("expected score 29, got %d", round.Score.Total)
	}
}
