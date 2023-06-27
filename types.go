package robert

// APIRequest represents an API request with the endpoint to send the request to and the key to authenticate the request.
type APIRequest struct {
	Endpoint string // Endpoint is the API endpoint to send the request to.
	Key      string // Key is the API key to authenticate the request with.
}

// APIResponse represents the response from the OpenAI API, including the unique identifier, object type, creation timestamp, model name, and usage statistics.
type APIResponse struct {
	ID      string   `json:"id"`      // ID is the unique identifier of the response.
	Object  string   `json:"object"`  // Object is the endpoint type returned in the response.
	Created int64    `json:"created"` // Created is the timestamp of when the response was created.
	Model   string   `json:"model"`   // Model is the name of the AI model used to generate the response.
	Usage   struct { // Usage contains information about the API usage for the request.
		PromptTokens     int `json:"prompt_tokens"`     // PromptTokens is the number of prompt tokens used in the request.
		CompletionTokens int `json:"completion_tokens"` // CompletionTokens is the number of completion tokens generated in the response.
		TotalTokens      int `json:"total_tokens"`      // TotalTokens is the total number of tokens used in the request and generated in the response.
	} `json:"usage"`
	Choices []struct { // Choices is an array of objects that contains a collection of response messages.
		Message struct { // Message is the text message associated with the choice.
			Role    string `json:"role"`    // Role is the role of the message sender.
			Content string `json:"content"` // Content is the text content of the message.
		} `json:"message"`
		FinishReason string `json:"finish_reason"` // FinishReason is the reason why the conversation ended.
		Index        int    `json:"index"`         // Index is the index of the choice in the array of choices.
	} `json:"choices"`
}
