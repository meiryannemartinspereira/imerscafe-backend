package service

import (
	"imerscafe-backend/internal/ai"
	"imerscafe-backend/internal/domain"
	"testing"
)

func TestAIConversation(t *testing.T) {
	fakeAI := &FakeAIService{}

	customer := domain.Customer{
		ID:   "customer-1",
		Type: domain.CustomerInRush,
	}

	profile := domain.CustomerProfiles[customer.Type]

	recipe := domain.Recipe{
		ID:   "recipe-1",
		Name: "Cappuccino",
	}

	conversation := domain.Conversation{}

	firstRequest := ai.AIRequest{
		CustomerType: string(customer.Type),
		Behavior:     profile.Behavior,
		RecipeName:   recipe.Name,
		Conversation: conversation,
	}

	firstResponse := fakeAI.SimulateCustomer(firstRequest)

	conversation.AddMessage(domain.ConversationMessage{
		Role:    "CUSTOMER",
		Content: firstResponse.Message,
	})

	t.Log("========== AI CONVERSATION ==========")

	t.Log("[1] CUSTOMER")
	t.Logf("Message: %s", firstResponse.Message)

	bartenderMessage := "Claro! Posso preparar um cappuccino para você."

	conversation.AddMessage(domain.ConversationMessage{
		Role:    "BARTENDER",
		Content: bartenderMessage,
	})

	t.Log("[2] BARTENDER")
	t.Logf("Message: %s", bartenderMessage)

	secondRequest := ai.AIRequest{
		CustomerType: string(customer.Type),
		Behavior:     profile.Behavior,
		RecipeName:   recipe.Name,
		Conversation: conversation,
	}

	secondResponse := fakeAI.SimulateCustomer(secondRequest)

	conversation.AddMessage(domain.ConversationMessage{
		Role:    "CUSTOMER",
		Content: secondResponse.Message,
	})

	t.Log("[3] CUSTOMER")
	t.Logf("Message: %s", secondResponse.Message)

	t.Log("[4] COMPLETE CONVERSATION")

	for _, message := range conversation.Messages {
		t.Logf("%s: %s", message.Role, message.Content)
	}

	t.Log("=====================================")

	if firstResponse.Message == "" {
		t.Error("expected first AI response")
	}

	if secondResponse.Message == "" {
		t.Error("expected second AI response")
	}

	if len(secondRequest.Conversation.Messages) != 2 {
		t.Errorf(
			"expected 2 messages in second request, got %d",
			len(secondRequest.Conversation.Messages),
		)
	}

	if len(conversation.Messages) != 3 {
		t.Errorf(
			"expected 3 messages in conversation, got %d",
			len(conversation.Messages),
		)
	}

	if conversation.Messages[0].Role != "CUSTOMER" {
		t.Errorf(
			"expected first message role CUSTOMER, got %s",
			conversation.Messages[0].Role,
		)
	}

	if conversation.Messages[1].Role != "BARTENDER" {
		t.Errorf(
			"expected second message role BARTENDER, got %s",
			conversation.Messages[1].Role,
		)
	}

	if conversation.Messages[2].Role != "CUSTOMER" {
		t.Errorf(
			"expected third message role CUSTOMER, got %s",
			conversation.Messages[2].Role,
		)
	}

	if secondRequest.Conversation.Messages[0].Content != firstResponse.Message {
		t.Error("expected first customer message in conversation history")
	}

	if secondRequest.Conversation.Messages[1].Content != bartenderMessage {
		t.Error("expected bartender message in conversation history")
	}
}
