package image

import "github.com/rmay1er/reptiloid-go/reptiloid"

/*
FluxSchnell and FluxDev are two models provided for image generation:
- FluxSchnell is a fast model optimized for quick generation.
- FluxDev is better suited for images containing text, though it is slightly more expensive and slower.
*/
var (
	FluxSchnell = reptiloid.NewReplicateModel[FluxSchnellInput]("black-forest-labs/flux-schnell")
	FluxDev     = reptiloid.NewReplicateModel[FluxDevInput]("black-forest-labs/flux-dev")
	FluxUltra   = reptiloid.NewReplicateModel[FluxUltraInput]("black-forest-labs/flux-1.1-pro-ultra")
	FluxPro     = reptiloid.NewReplicateModel[FluxProInput]("black-forest-labs/flux-1.1-pro")
)

/*
FluxSchnellInput represents input parameters for the Flux Schnell model.
It contains various configuration options to control the image generation process.
*/
type FluxSchnellInput struct {
	// Prompt is the main text input for image generation.
	Prompt string `json:"prompt"`

	// AspectRatio specifies the width-to-height ratio of the output image.
	// Allowed values include "21:9", "16:9", "3:2", "4:3", "5:4", "1:1", "4:5", "3:4", "2:3", "9:16", "9:21".
	AspectRatio *string `json:"aspect_ratio,omitempty"`

	// NumOutputs defines how many images to generate.
	NumOutputs *int `json:"num_outputs,omitempty"`

	// NumInferenceSteps defines the number of steps in the diffusion process.
	NumInferenceSteps *int `json:"num_inference_steps,omitempty"`

	// Seed sets the random seed to make generation reproducible.
	Seed *int `json:"seed,omitempty"`

	// OutputFormat sets the file format for the generated images ("jpg" or "png").
	OutputFormat *string `json:"output_format,omitempty"`

	// OutputQuality sets the quality of the output images (1-100).
	OutputQuality *int `json:"output_quality,omitempty"`

	// GoFast enables faster but less precise generation.
	GoFast *bool `json:"go_fast,omitempty"`

	// Megapixels controls the resolution of the generated image.
	Megapixels *string `json:"megapixels,omitempty"`

	// DisableSafetyChecker disables built-in content filters.
	DisableSafetyChecker *bool `json:"disable_safety_checker,omitempty"`
}

// FluxDevInput represents input parameters for the Flux Dev model,
// which supports image-based prompting.
type FluxDevInput struct {
	// Prompt is the main input text for image generation.
	Prompt string `json:"prompt"`

	// Seed is the random seed for reproducibility. Use -1 for random seed.
	Seed *int `json:"seed,omitempty"`

	// Image is the URI of an image to guide generation along with the text prompt.
	Image *string `json:"image,omitempty"`

	// AspectRatio defines the aspect ratio of the output image.
	// Allowed values are the same as FluxSchnellInput plus "21:9".
	AspectRatio *string `json:"aspect_ratio,omitempty"`

	// NumOutputs specifies how many images to generate.
	NumOutputs *int `json:"num_outputs,omitempty"`

	// PromptStrength blends the influence of the prompt and image.
	// 1.0 means fully replacing image information.
	PromptStrength *float64 `json:"prompt_strength,omitempty"`

	// NumInferenceSteps sets the number of diffusion steps.
	NumInferenceSteps *int `json:"num_inference_steps,omitempty"`

	// Guidance controls the influence of the prompt on generation.
	// Lower values produce more realistic images; try values 2 to 3.5.
	Guidance *float64 `json:"guidance,omitempty"`

	// DisableSafetyChecker disables content filters.
	DisableSafetyChecker *bool `json:"disable_safety_checker,omitempty"`

	// GoFast enables faster generation with optimizations.
	GoFast *bool `json:"go_fast,omitempty"`

	// Megapixels controls output image resolution.
	Megapixels *string `json:"megapixels,omitempty"`

	// OutputFormat specifies the file format; supports "jpg", "png", and "webp".
	OutputFormat *string `json:"output_format,omitempty"`

	// OutputQuality is the quality for jpg and webp (0-100), ignored for png.
	OutputQuality *int `json:"output_quality,omitempty"`
}

// FluxUltraInput represents input parameters for the Flux 1.1 Pro Ultra model.
// This model generates less processed, more natural-looking images.
type FluxUltraInput struct {
	// Prompt is the text prompt guiding image generation.
	Prompt string `json:"prompt"`

	// Seed is the random seed for reproducibility.
	Seed *int `json:"seed,omitempty"`

	// ImagePrompt is a URI to an image used alongside the prompt to guide generation composition.
	// Must be jpeg, png, gif, or webp.
	ImagePrompt *string `json:"image_prompt,omitempty"`

	// AspectRatio defines the output image aspect ratio.
	// Allowed values: "21:9", "16:9", "3:2", "4:3", "5:4", "1:1", "4:5", "3:4", "2:3", "9:16", "9:21".
	AspectRatio *string `json:"aspect_ratio,omitempty"`

	// NumOutputs determines how many images to generate.
	NumOutputs *int `json:"num_outputs,omitempty"`

	// ImagePromptStrength blends the influence between text prompt and image prompt.
	// Ranges from 0 (only text) to 1 (only image). Default is 0.1.
	ImagePromptStrength *float64 `json:"image_prompt_strength,omitempty"`

	// SafetyTolerance sets the tolerance level for safety filters.
	// Values range from 1 (strict) to 6 (permissive). Default is 2.
	SafetyTolerance *int `json:"safety_tolerance,omitempty"`

	// Raw when true generates less processed, more natural-looking images.
	Raw *bool `json:"raw,omitempty"`

	// OutputFormat specifies output image format: "jpg" or "png".
	OutputFormat *string `json:"output_format,omitempty"`
}

// FluxProInput represents input parameters for the Flux 1.1 Pro model.
type FluxProInput struct {
	// Prompt is the text prompt guiding image generation.
	Prompt string `json:"prompt"`

	// Seed is the random seed for reproducibility.
	Seed *int `json:"seed,omitempty"`

	// ImagePrompt is a URI of an image to guide generation composition.
	// Must be jpeg, png, gif, or webp.
	ImagePrompt *string `json:"image_prompt,omitempty"`

	// AspectRatio defines the output image aspect ratio.
	// Allowed values: "custom", "1:1", "16:9", "3:2", "2:3", "4:5", "5:4", "9:16", "3:4", "4:3".
	AspectRatio *string `json:"aspect_ratio,omitempty"`

	// NumOutputs determines how many images to generate.
	NumOutputs *int `json:"num_outputs,omitempty"`

	// SafetyTolerance sets the tolerance level for safety filters.
	// Values range from 1 (strict) to 6 (permissive).
	SafetyTolerance *int `json:"safety_tolerance,omitempty"`

	// PromptUpsampling when true modifies the prompt automatically for more creative generation.
	PromptUpsampling *bool `json:"prompt_upsampling,omitempty"`

	// OutputFormat specifies output image format: "webp", "jpg", or "png".
	OutputFormat *string `json:"output_format,omitempty"`

	// OutputQuality sets the quality of output images (0 to 100).
	OutputQuality *int `json:"output_quality,omitempty"`

	// Width for custom aspect ratio, multiple of 32, between 256 and 1440.
	Width *int `json:"width,omitempty"`

	// Height for custom aspect ratio, multiple of 32, between 256 and 1440.
	Height *int `json:"height,omitempty"`
}
