package chatgpt

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

func ChatCompletion(token string, req ChatCompletionRequest) (r io.ReadCloser, error error) {
	client := resty.New()
	fmt.Println("gpt chat completion")

	const GPTChatCompletionURL = "https://api.openai.com/v1/chat/completions"

	reqStr, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := client.R().
		SetDoNotParseResponse(true).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(reqStr).
		Post(GPTChatCompletionURL)

	if err != nil {
		return nil, err
	}

	return resp.RawBody(), nil
}
