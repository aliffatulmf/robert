package robert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewAPIRequest(endpoint, key string) *APIRequest {
	return &APIRequest{
		Endpoint: endpoint,
		Key:      key,
	}
}

func (r *APIRequest) SendAPIRequest(payload []byte) (*APIResponse, error) {
	if payload == nil {
		return nil, fmt.Errorf("payload is nil")
	}

	req, err := http.NewRequest(http.MethodPost, r.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Authorization", r.Key)

	client := &http.Client{}
	client.Timeout = time.Second * 10

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	res := &APIResponse{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res, nil
}
