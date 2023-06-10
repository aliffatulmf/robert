package robert

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	Basic = iota
	Turbo
)

const (
	System = iota
	User
	Assistant
)

type payload struct {
	Messages        []map[string]string `json:"messages"`
	Stream          bool                `json:"stream"`
	Model           string              `json:"model"`
	Temperature     float32             `json:"temperature"`
	PresencePenalty float32             `json:"presence_penalty"`
}

func NewPayload() *payload {
	return &payload{
		Stream: true,
		Model:  "gpt-3.5-turbo-0301",
	}
}

func (p *payload) SetMessage(role int, message string) error {
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
	p.Messages = append(p.Messages, map[string]string{
		"role":    roleName,
		"content": message,
	})
	return nil
}

func (p *payload) SetMessages(role int, messages ...string) error {
	if role < Basic || role > Assistant {
		return errors.New("invalid role")
	}
	for _, message := range messages {
		if err := p.SetMessage(role, message); err != nil {
			return err
		}
	}
	return nil
}

func (p *payload) SetModel(model int) error {
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

func (p *payload) SetTemperature(temp float32) error {
	if temp < -1.0 || temp > 1.0 {
		return errors.New("invalid temperature. Must be between -1.0 and 1.0")
	}
	p.Temperature = temp
	return nil
}

func (p *payload) SetPresencePenalty(penalty float32) error {
	if penalty < -1.0 || penalty > 1.0 {
		return errors.New("invalid presence penalty. Must be between -1.0 and 1.0")
	}
	p.PresencePenalty = penalty
	return nil
}

func (p *payload) SetStream(stream bool) {
	p.Stream = stream
}

func (p *payload) Payload() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to create payload: %w", errors.Unwrap(err))
	}
	return b, nil
}
