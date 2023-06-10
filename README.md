## Robert API Library Documentation

The Robert API library provides a convenient way to interact with the Robert API for natural language processing tasks. This documentation provides a brief overview of the library and its usage.

### Installation

To use the Robert API library, you can use the `go get` command to download and install the library directly into your Go project. Open your terminal or command prompt and run the following command:

```shell
go get -u github.com/aliffatulmf/robert
```

### Usage

After the installation, you can import the `robert` package into your Go code and start using the library.

```go
import "github.com/aliffatulmf/robert"
```

The library consists of several components:

- **APIRequest**: Represents the parameters for sending an API request.
- **APIResponse**: Represents the response received from the API.
- **Payload**: Represents the payload for the API request, including messages, model settings, temperature, and presence penalty.

#### Creating an API Request

To create an API request, you need to create an instance of `APIRequest` by providing the endpoint and key.

```go
request := robert.NewAPIRequest("https://api.robert.com/endpoint", "your-api-key")
```

#### Setting Payload

Before sending an API request, you need to create a payload using the `Payload` struct. It provides methods to set various attributes of the payload.

```go
payload := robert.NewPayload()
payload.SetMessages(robert.System, "You are GPT-4, OpenAI’s advanced language model.")
payload.SetMessages(robert.User, "Hello!", "How are you?")
payload.SetModel(robert.Turbo)
```

#### Sending the Request

Once you have created the payload, you can send the API request using the `SendAPIRequest` method of the `APIRequest` struct.

```go
response, err := request.SendAPIRequest(payload.Payload())
if err != nil {
    // handle error
}
```

The `SendAPIRequest` method takes the payload as a byte array and returns the API response or an error. If the request is successful, you can access the response data using the fields of the `APIResponse` struct.

```go
fmt.Println("Response ID:", response.ID)
fmt.Println("Model used:", response.Model)
// ...
```

### Examples

Here's an example that demonstrates the basic usage of the Robert API library:

```go
package main

import (
	"fmt"
	"github.com/aliffatulmf/robert"
)

func main() {
	request := robert.NewAPIRequest("https://api.robert.com/endpoint", "your-api-key")

	payload := robert.NewPayload()
	payload.SetMessages(robert.System, "You are GPT-4, OpenAI’s advanced language model.")
	payload.SetMessages(robert.User, "Hello!", "How are you?")
	payload.SetModel(robert.Turbo)

	response, err := request.SendAPIRequest(payload.Payload())
	if err != nil {
		// handle error
	}

	fmt.Println("Response ID:", response.ID)
	fmt.Println("Model used:", response.Model)
	// ...
}
```

Note: Replace "https://api.robert.com/endpoint" with the actual API endpoint URL and "your-api-key" with your valid API key.