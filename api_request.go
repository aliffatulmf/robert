package robert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	res := &APIResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return res, nil
}
