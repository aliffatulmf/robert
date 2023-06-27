package robert

import (
	"encoding/json"
	"fmt"
)

// Role represents the role of the chat message sender.
type Role int

const (
	System Role = iota
	User
	Assistant
)

// ModelType represents the type of model to use for generating responses.
type ModelType int

const (
	Basic ModelType = iota
	Turbo
	Basic16K
	Turbo16K
)

// ChatMessage represents a single chat message, including the role of the sender and the message content.
type ChatMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Payload represents a payload for the OpenAI API.
type Payload struct {
	ChatMessages    []ChatMessage `json:"messages"`         // A list of chat messages.
	Model           string        `json:"model"`            // The type of model to use.
	Temperature     float32       `json:"temperature"`      // The sampling temperature to use.
	PresencePenalty float32       `json:"presence_penalty"` // The presence penalty to use.
}

// NewPayload returns a new Payload with default values.
func NewPayload(modelType ModelType, temp float32, presencePenalty float32) Payload {
	p := Payload{}
	p.Model = getModelType(modelType)
	p.Temperature = temp
	p.PresencePenalty = presencePenalty
	return p
}

// getModelType returns the appropriate model type as per given ModelType.
func getModelType(modelType ModelType) string {
	switch modelType {
	case Turbo:
		return "gpt-3.5-turbo"
	case Turbo16K:
		return "gpt-3.5-turbo-16k"
	case Basic16K:
		return "gpt-3.5-turbo-16k-0613"
	case Basic:
		fallthrough
	default:
		return "gpt-3.5-turbo-0613"
	}
}

// AddMessage adds a new chat message to the Payload.
func (p *Payload) AddMessage(role Role, message string) {
	p.ChatMessages = append(p.ChatMessages, ChatMessage{
		Role:    role,
		Content: message,
	})
}

// AddMessages adds multiple chat messages to the Payload.
func (p *Payload) AddMessages(role Role, messages ...string) {
	for _, message := range messages {
		p.AddMessage(role, message)
	}
}

// SetModel sets the model type for the Payload.
func (p *Payload) SetModel(modelType ModelType) error {
	p.Model = getModelType(modelType)
	return nil
}

// SetTemperature sets the temperature for the Payload.
func (p *Payload) SetTemperature(temp float32) error {
	if temp < -1.0 || temp > 1.0 {
		return fmt.Errorf("invalid temperature. Must be between -1.0 and 1.0")
	}
	p.Temperature = temp
	return nil
}

// SetPresencePenalty sets the presence penalty for the Payload.
func (p *Payload) SetPresencePenalty(penalty float32) error {
	if penalty < -1.0 || penalty > 1.0 {
		return fmt.Errorf("invalid presence penalty. Must be between -1.0 and 1.0")
	}
	p.PresencePenalty = penalty
	return nil
}

// ToJSON serializes the Payload to JSON.
func (p *Payload) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
