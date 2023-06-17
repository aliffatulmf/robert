package robert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// NewAPIRequest creates a new APIRequest instance with the provided endpoint and key.
// The endpoint is the URL endpoint for the API request.
// The key is a token used for authentication or identification purposes.
// It is recommended to generate the key using the Token function to ensure it is properly formatted.
// Example usage:
//
//	endpoint := "https://api.robert.com"
//	token := "myToken123"
//	key, err := Token(token)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	apiRequest := NewAPIRequest(endpoint, key)
func NewAPIRequest(endpoint, key string) *APIRequest {
	return &APIRequest{
		Endpoint: endpoint,
		Key:      key,
	}
}

// SendAPIRequest sends an API request with the given payload.
// Returns an APIResponse or an error if the request failed.
func (r *APIRequest) SendAPIRequest(payload []byte) (*APIResponse, error) {
	if payload == nil {
		return nil, errors.New("payload is nil")
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

	return r.parseResponse(resp.Body)
}

// createRequest creates a new HTTP request with the given payload and endpoint.
func (r *APIRequest) createRequest(payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		r.Endpoint,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Authorization", r.Key)
	return req, nil
}

// parseResponse reads the body of an API response and decodes it into an APIResponse struct.
func (r *APIRequest) parseResponse(body io.Reader) (*APIResponse, error) {
	res := &APIResponse{}
	if err := json.NewDecoder(body).Decode(res); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return res, nil
}

// containBearer checks if the given token string contains the "Bearer " prefix.
// It returns true if the prefix is found, and false otherwise.
// The function is case-insensitive, so it will match both "Bearer " and "bearer ".
func containBearer(token string) bool {
	return strings.ToLower(token[:7]) == "bearer "
}

// containSpace returns the index of the first space character in the input token string, and a boolean indicating whether a space was found.
// If a space is found, the boolean is true, otherwise it is false.
// The index is zero-based, meaning that the first character has index 0, the second has index 1, and so on.
// If no space is found, the index is -1.
func containSpace(token string) (int, bool) {
	space := strings.IndexRune(token, ' ')
	return space, space != -1
}

// Token returns a modified version of the input token string, by concatenating its prefix and suffix parts.
// The input token must have a length of at least 7 characters, otherwise an error is returned.
// The prefix and suffix parts are extracted from the input token using the getTokenParts function.
// If an error occurs during the extraction, it is propagated to the caller.
// The returned string is the concatenation of the prefix and suffix parts, in that order.
func Token(token string) (string, error) {
	if len(token) < 7 {
		return "", fmt.Errorf("token length %d, expected at least 7", len(token))
	}
	prefix, suffix, err := getTokenParts(token)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", prefix, suffix), nil
}

// getTokenParts extracts the token type and value from a given token string.
// If the token contains the string "Bearer ", it is assumed to be in the format "Bearer <token value>".
// Otherwise, the token is assumed to be in the format "<token type> <token value>".
// The function returns the token type ("Bearer "), the token value, and an error if any.
func getTokenParts(token string) (string, string, error) {
	if containBearer(token) {
		return "Bearer ", token[7:], nil
	}
	index, hasSpace := containSpace(token)
	if hasSpace {
		return "Bearer ", token, nil
	}
	return "Bearer ", token[index+1:], nil
}
