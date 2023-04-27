package entity

type ChatConfig struct {
	Model *Model
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Message              []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}