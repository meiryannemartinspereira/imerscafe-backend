package ai

type AIService interface {
	SimulateCustomer(request AIRequest) AIResponse
}
