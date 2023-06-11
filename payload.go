package robert

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Constants for the role of the chat message sender.
const (
	// The system role.
	System = iota
	// The user role.
	User
	// The assistant role.
	Assistant
)

// Constants for the type of model to use for generating responses.
const (
	// The basic model.
	Basic = iota
	// The turbo model.
	Turbo
)

// The ChatMessage struct represents a single chat message, including the role of the sender and the message content.
type ChatMessage struct {
	// Role field represents the role of the sender
	Role string `json:"role"`

	// Content field represents the message content
	Content string `json:"content"`
}

// Payload represents a payload for the OpenAI API.
type Payload struct {
	// ChatMessages is a list of chat messages.
	ChatMessages []ChatMessage `json:"messages"`

	// Model is the type of model to use.
	Model string `json:"model"`

	// Temperature is the sampling temperature to use.
	Temperature float32 `json:"temperature"`

	// PresencePenalty is the presence penalty to use.
	PresencePenalty float32 `json:"presence_penalty"`
}

// NewPayload returns a new Payload with default values.
func NewPayload() *Payload {
	return &Payload{
		Model:           "gpt-3.5-turbo-0301",
		Temperature:     0.5,
		PresencePenalty: 0.0,
	}
}

// AddMessage adds a new chat message to the Payload.
func (p *Payload) AddMessage(role int, message string) error {
	var roleName string
	switch role {
	case System:
		roleName = "system"
	case User:
		roleName = "user"
	case Assistant:
		roleName = "assistant"
	default:
		return errors.New("invalid role")
	}

	p.ChatMessages = append(p.ChatMessages, ChatMessage{
		Role:    roleName,
		Content: message,
	})

	return nil
}

// AddMessages adds multiple chat messages to the Payload.
func (p *Payload) AddMessages(role int, messages ...string) error {
	if role < Basic || role > Assistant {
		return errors.New("invalid role")
	}

	for _, message := range messages {
		if err := p.AddMessage(role, message); err != nil {
			return err
		}
	}
	return nil
}

// SetModel sets the model type for the Payload.
func (p *Payload) SetModel(model int) error {
	switch model {
	case Basic:
		p.Model = "gpt-3.5-turbo-0301"
	case Turbo:
		p.Model = "gpt-3.5-turbo"
	default:
		return errors.New("invalid model")
	}
	return nil
}

// SetTemperature sets the temperature for the Payload.
func (p *Payload) SetTemperature(temp float32) error {
	if temp < -1.0 || temp > 1.0 {
		return errors.New("invalid temperature. Must be between -1.0 and 1.0")
	}
	p.Temperature = temp
	return nil
}

// SetPresencePenalty sets the presence penalty for the Payload.
func (p *Payload) SetPresencePenalty(penalty float32) error {
	if penalty < -1.0 || penalty > 1.0 {
		return errors.New("invalid presence penalty. Must be between -1.0 and 1.0")
	}
	p.PresencePenalty = penalty
	return nil
}

// ToJSON serializes the Payload to JSON.
func (p *Payload) ToJSON() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create payload: %w", errors.Unwrap(err))
	}
	return b, nil
}
