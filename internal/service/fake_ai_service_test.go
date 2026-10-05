package service

import "imerscafe-backend/internal/ai"

type FakeAIService struct {
	LastRequest ai.AIRequest
}

func (f *FakeAIService) SimulateCustomer(request ai.AIRequest) ai.AIResponse {
	f.LastRequest = request

	if len(request.Conversation.Messages) == 0 {
		return ai.AIResponse{
			Message: "Olá! Gostaria de um café, por favor.",
		}
	}

	return ai.AIResponse{
		Message: "Claro! Mas preciso que seja rápido, por favor.",
	}
}
