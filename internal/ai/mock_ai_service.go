package ai

type MockAIService struct{}

func (m *MockAIService) SimulateCustomer(request AIRequest) AIResponse {
	return AIResponse{
		Message: "Olá! Gostaria de um café, por favor.",
	}
}
