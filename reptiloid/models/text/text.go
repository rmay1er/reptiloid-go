package text

import "github.com/rmay1er/reptiloid-go/reptiloid"

/*
GPT-5 Nano is a lightweight variant suitable for fast and cost-effective text tasks, with limited reasoning capabilities.
GPT-5 Mini is a balanced model offering better performance and reasoning, slightly more expensive and slower than Nano.
GPT-5 is the full model designed for high reasoning effort and complex text understanding, typically more resource-intensive.
*/
var (
	GPT5nano = reptiloid.NewReplicateModel[GPT5SeriesInput]("openai/gpt-5-nano")
	GPT5mini = reptiloid.NewReplicateModel[GPT5SeriesInput]("openai/gpt-5-mini")
	GPT5     = reptiloid.NewReplicateModel[GPT5SeriesInput]("openai/gpt-5")
)

/*
GPT-4.1 Nano is a lightweight variant of GPT-4.1 for fast and cost-effective text tasks.
GPT-4.1 Mini is a balanced model of GPT-4.1 offering better performance and reasoning.
GPT-4.1 is the full GPT-4.1 model with advanced configurable parameters.
*/
var (
	GPT41nano = reptiloid.NewReplicateModel[GPT4SeriesInput]("openai/gpt-4.1-nano")
	GPT41mini = reptiloid.NewReplicateModel[GPT4SeriesInput]("openai/gpt-4.1-mini")
	GPT41     = reptiloid.NewReplicateModel[GPT4SeriesInput]("openai/gpt-4.1")
	GPT4o     = reptiloid.NewReplicateModel[GPT4SeriesInput]("openai/gpt-4o")
	GPT4oMini = reptiloid.NewReplicateModel[GPT4SeriesInput]("openai/gpt-4o-mini")
)

/*
DeepSeekR1 is the thoughtful, reasoning-capable model of the DeepSeek family.
It excels at understanding complex prompts and generating insightful responses.
DeepSeekV3 is a newer variant optimized for diverse text tasks with improved performance and versatility.
*/
var (
	DeepSeekR1 = reptiloid.NewReplicateModel[DeepSeekInput]("deepseek-ai/deepseek-r1")
	DeepSeekV3 = reptiloid.NewReplicateModel[DeepSeekInput]("deepseek-ai/deepseek-v3")
)

/*
GPT5Input represents input parameters for the GPT-5 models.
It contains various fields to provide prompts, control verbosity, reasoning effort, image inputs, and token limits to the model.
*/
type GPT5SeriesInput struct {
	// Prompt is the prompt to send to the model. Do not use if using Messages.
	Prompt string `json:"prompt"`

	// Messages is a JSON string representing a list of messages.
	// If provided, Prompt and SystemPrompt are ignored.
	Messages []map[string]any `json:"messages,omitempty"`

	// SystemPrompt sets the assistant's behavior.
	SystemPrompt *string `json:"system_prompt,omitempty"`

	// ImageInput is a list of images (as URIs) to send to the model.
	ImageInput []string `json:"image_input,omitempty"`

	// ReasoningEffort constrains effort on reasoning for GPT-5 models.
	// Supported values: "minimal", "low", "medium", "high".
	// "minimal" results in faster responses with less reasoning.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`

	// Verbosity constrains the verbosity of the model's response.
	// Supported values: "low", "medium", "high".
	// Lower values yield concise responses; higher values yield more comprehensive answers.
	Verbosity *string `json:"verbosity,omitempty"`

	// MaxCompletionTokens sets the max number of completion tokens to generate.
	// For higher reasoning efforts, increase this to avoid empty responses.
	MaxCompletionTokens *int `json:"max_completion_tokens,omitempty"`
}

/*
GPT4SeriesInput represents input parameters for the GPT-4o GPT-4.1 models.
It includes fields for prompt, messages, sampling parameters, penalties, image inputs, system prompt, and token limits.
*/
type GPT4SeriesInput struct {
	// Prompt is the prompt to send to the model. Do not use if using Messages.
	Prompt string `json:"prompt"`

	// Messages is a JSON string representing a list of messages.
	// If provided, prompt and system_prompt are ignored.
	Messages []map[string]any `json:"messages,omitempty"`

	// SystemPrompt sets the assistant's behavior.
	SystemPrompt *string `json:"system_prompt,omitempty"`

	// ImageInput is a list of images (as URIs) to send to the model.
	ImageInput []string `json:"image_input,omitempty"`

	// Temperature controls the sampling temperature between 0 and 2.
	Temperature *float64 `json:"temperature,omitempty"`

	// TopP is the nucleus sampling parameter between 0 and 1.
	TopP *float64 `json:"top_p,omitempty"`

	// MaxCompletionTokens sets the max number of completion tokens to generate.
	MaxCompletionTokens *int `json:"max_completion_tokens,omitempty"`

	// FrequencyPenalty penalizes repeated tokens; between -2 and 2.
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// PresencePenalty penalizes new tokens based on whether they appear in the text so far; between -2 and 2.
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`
}

/*
DeepSeekR1Input represents input parameters for the DeepSeek-R1 & DeepSeek-V3 models.
It includes fields for prompt, sampling parameters, penalties, and max tokens.
*/
type DeepSeekInput struct {
	// Prompt is the prompt to send to the model.
	Prompt string `json:"prompt"`

	// MaxTokens is the maximum number of tokens to generate.
	MaxTokens *int `json:"max_tokens,omitempty"`

	// Temperature controls the sampling temperature.
	Temperature *float64 `json:"temperature,omitempty"`

	// PresencePenalty penalizes new tokens based on whether they appear in the text so far.
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`

	// FrequencyPenalty penalizes repeated tokens.
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// TopP is the nucleus sampling parameter.
	TopP *float64 `json:"top_p,omitempty"`
}
