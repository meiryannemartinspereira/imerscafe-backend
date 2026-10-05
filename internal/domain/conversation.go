package domain

type Conversation struct {
	Messages []ConversationMessage
}

func (c *Conversation) AddMessage(message ConversationMessage) {
	c.Messages = append(c.Messages, message)
}
