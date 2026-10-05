package service

import (
	"imerscafe-backend/internal/domain"
	"testing"
)

func TestRoundServiceCreateRound(t *testing.T) {
	scoreService := NewScoreService()
	fakeAI := &FakeAIService{}

	roundService := NewRoundService(scoreService, fakeAI)

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

	round, aiResponse := roundService.CreateRound(
		customer,
		recipe,
		preparation,
		evaluation,
	)

	t.Log("========== ROUND FLOW ==========")

	t.Log("[1] CUSTOMER")
	t.Logf("Type: %s", round.Customer.Type)
	t.Logf("Description: %s", round.CustomerProfile.Description)
	t.Logf("Behavior: %s", round.CustomerProfile.Behavior)

	t.Log("[2] RECIPE")
	t.Logf("Name: %s", round.Recipe.Name)

	t.Log("[3] AI REQUEST")
	t.Logf("CustomerType: %s", fakeAI.LastRequest.CustomerType)
	t.Logf("Behavior: %s", fakeAI.LastRequest.Behavior)
	t.Logf("RecipeName: %s", fakeAI.LastRequest.RecipeName)

	t.Log("[4] AI RESPONSE")
	t.Logf("Message: %s", aiResponse.Message)

	t.Log("[5] PREPARATION")
	t.Logf("Correct: %t", round.PreparationResult.Correct)

	t.Log("[6] SOFT SKILLS")
	t.Logf("Communication: %d", round.SoftSkillEvaluation.Communication)
	t.Logf("Empathy: %d", round.SoftSkillEvaluation.Empathy)
	t.Logf("Politeness: %d", round.SoftSkillEvaluation.Politeness)
	t.Logf("Clarity: %d", round.SoftSkillEvaluation.Clarity)
	t.Logf("Feedback: %s", round.SoftSkillEvaluation.Feedback)

	t.Log("[7] SCORE")
	t.Logf("Total: %d", round.Score.Total)

	t.Log("================================")

	if round.ID == "" {
		t.Error("expected round ID")
	}

	if round.Customer.Type != domain.CustomerCalm {
		t.Errorf("expected customer type %s, got %s",
			domain.CustomerCalm,
			round.Customer.Type,
		)
	}

	if round.Recipe.Name != "Cappuccino" {
		t.Errorf("expected recipe Cappuccino, got %s",
			round.Recipe.Name,
		)
	}

	if fakeAI.LastRequest.CustomerType != string(customer.Type) {
		t.Errorf("expected customer type %s, got %s",
			customer.Type,
			fakeAI.LastRequest.CustomerType,
		)
	}

	if fakeAI.LastRequest.Behavior != round.CustomerProfile.Behavior {
		t.Errorf("expected behavior %s, got %s",
			round.CustomerProfile.Behavior,
			fakeAI.LastRequest.Behavior,
		)
	}

	if fakeAI.LastRequest.RecipeName != recipe.Name {
		t.Errorf("expected recipe %s, got %s",
			recipe.Name,
			fakeAI.LastRequest.RecipeName,
		)
	}

	if aiResponse.Message == "" {
		t.Error("expected AI response message")
	}

	if round.Score.Total != 69 {
		t.Errorf("expected score 69, got %d",
			round.Score.Total,
		)
	}
}

func TestRoundServiceCreateRoundWithIncorrectPreparation(t *testing.T) {
	scoreService := NewScoreService()
	fakeAI := &FakeAIService{}

	roundService := NewRoundService(scoreService, fakeAI)

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

	round, aiResponse := roundService.CreateRound(
		customer,
		recipe,
		preparation,
		evaluation,
	)

	t.Log("========== INCORRECT PREPARATION ==========")
	t.Logf("Customer: %s", round.Customer.Type)
	t.Logf("Recipe: %s", round.Recipe.Name)
	t.Logf("AI Response: %s", aiResponse.Message)
	t.Logf("Preparation Correct: %t", round.PreparationResult.Correct)
	t.Logf("Score: %d", round.Score.Total)
	t.Log("===========================================")

	if round.PreparationResult.Correct {
		t.Error("expected preparation to be incorrect")
	}

	if round.Score.Total != 29 {
		t.Errorf("expected score 29, got %d",
			round.Score.Total,
		)
	}
}
