package robert

type APIRequest struct {
	// Endpoint is the API endpoint to send the request to.
	Endpoint string

	// Key is the API key to authenticate the request with.
	Key string
}

type APIResponse struct {
	// ID is the unique identifier of the response.
	ID string `json:"id"`

	// Object is the endpoint type returned in the response.
	Object string `json:"object"`

	// Created is the timestamp of when the response was created.
	Created int64 `json:"created"`

	// Model is the name of the AI model used to generate the response.
	Model string `json:"model"`

	// Usage contains information about the API usage for the request.
	Usage struct {
		// PromptTokens is the number of prompt tokens used in the request.
		PromptTokens int `json:"prompt_tokens"`

		// CompletionTokens is the number of completion tokens generated in the response.
		CompletionTokens int `json:"completion_tokens"`

		// TotalTokens is the total number of tokens used in the request and generated in the response.
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`

	// Choices is an array of objects that contains a collection of response messages.
	Choices []struct {
		// Message is the text message associated with the choice.
		Message struct {
			// Role is the role of the message sender.
			Role string `json:"role"`

			// Content is the text content of the message.
			Content string `json:"content"`
		} `json:"message"`

		// FinishReason is the reason why the conversation ended.
		FinishReason string `json:"finish_reason"`

		// Index is the index of the choice in the array of choices.
		Index int `json:"index"`
	} `json:"choices"`
}
