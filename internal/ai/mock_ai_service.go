package ai

type MockAIService struct{}

func (m *MockAIService) SimulateCustomer(request AIRequest) AIResponse {
	if len(request.Conversation.Messages) == 0 {
		return AIResponse{
			Message: "Olá! Gostaria de um café, por favor.",
		}
	}

	return AIResponse{
		Message: "Claro! Mas preciso que seja rápido, por favor.",
	}
}
