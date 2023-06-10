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

	req, err := r.createRequest(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

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

	parse, err := r.parseResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return parse, nil
}

func (r *APIRequest) createRequest(payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, r.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Authorization", r.Key)
	return req, nil
}

func (r *APIRequest) parseResponse(body io.Reader) (*APIResponse, error) {
	res := &APIResponse{}
	if err := json.NewDecoder(body).Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}
