package domain

type Round struct {
	ID                  string
	Customer            Customer
	CustomerProfile     CustomerProfile
	Recipe              Recipe
	PreparationResult   PreparationResult
	SoftSkillEvaluation SoftSkillEvaluation
	Score               Score
}
