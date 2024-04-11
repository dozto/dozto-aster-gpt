package chatgpt

type ChatGPTModel string

const (
	GPT35Turbo   ChatGPTModel = "gpt-3.5-turbo"
	GPT4         ChatGPTModel = "gpt-4"
	GPT432K      ChatGPTModel = "gpt-4-32k"
	GPT4TurboPre ChatGPTModel = "gpt-4-turbo-preview"
)

type ChatGPTModelRole string

const (
	ChatGPTRoleUser      ChatGPTModelRole = "user"
	ChatGPTRoleSystem    ChatGPTModelRole = "system"
	ChatGPTRoleAssistant ChatGPTModelRole = "assistant"
)

type ChatMessage struct {
	Role    ChatGPTModelRole `json:"role"`
	Content string           `json:"content"`
}

type ChatCompletionRequest struct {
	// (Required)
	// ID of the model to use.
	Model ChatGPTModel `json:"model"`

	// Required
	// The messages to generate chat completions for
	Messages []ChatMessage `json:"messages"`

	// (Optional - default: 1)
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	// We generally recommend altering this or top_p but not both.
	Temperature float64 `json:"temperature,omitempty"`

	// (Optional - default: 1)
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	// We generally recommend altering this or temperature but not both.
	Top_P float64 `json:"top_p,omitempty"`

	// (Optional - default: 1)
	// How many chat completion choices to generate for each input message.
	N int `json:"n,omitempty"`

	// (Optional - default: infinite)
	// The maximum number of tokens allowed for the generated answer. By default,
	// the number of tokens the model can return will be (4096 - prompt tokens).
	MaxTokens int `json:"max_tokens,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far,
	// increasing the model's likelihood to talk about new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// (Optional - default: 0)
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// (Optional)
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
	User string `json:"user,omitempty"`

	Stream bool `json:"stream,omitempty"`
}
