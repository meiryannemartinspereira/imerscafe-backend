package ai

import "imerscafe-backend/internal/domain"

type AIRequest struct {
	CustomerType string
	Behavior     string
	RecipeName   string
	Conversation domain.Conversation
}
