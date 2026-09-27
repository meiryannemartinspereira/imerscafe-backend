package service

import (
	"imerscafe-backend/internal/domain"
	"math"
)

var customerBonuses = map[domain.CustomerType]int{
	domain.CustomerCalm:       5,
	domain.CustomerInRush:     10,
	domain.CustomerAesthetic:  15,
	domain.CustomerExecutive:  20,
	domain.CustomerSpecialist: 25,
	domain.CustomerIndecisive: 30,
}

func customerBonus(customer domain.Customer) int {
	return customerBonuses[customer.Type]
}

type ScoreService struct{}

func NewScoreService() *ScoreService {
	return &ScoreService{}
}

func preparationScore(preparation domain.PreparationResult) int {
	if preparation.Correct {
		return 40
	}

	return 0
}

func softSkillScore(evaluation domain.SoftSkillEvaluation) int {
	total := evaluation.Communication +
		evaluation.Empathy +
		evaluation.Politeness +
		evaluation.Clarity

	return int(math.Round(float64(total) * 0.75))
}

func (s *ScoreService) Calculate(
	customer domain.Customer,
	preparation domain.PreparationResult,
	evaluation domain.SoftSkillEvaluation,
) domain.Score {
	total := customerBonus(customer) +
		preparationScore(preparation) +
		softSkillScore(evaluation)

	return domain.Score{
		Total: total,
	}
}
