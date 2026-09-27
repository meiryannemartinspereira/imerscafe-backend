package domain

type Round struct {
	ID                  string
	Customer            Customer
	Recipe              Recipe
	PreparationResult   PreparationResult
	SoftSkillEvaluation SoftSkillEvaluation
	Score               Score
}
