package models

import "reptiloid-go/internal/replicate"

/*
FluxSchnell is a fast model optimized for quick generation.
*/
var FluxSchnell = replicate.NewReplicateModel[FluxSchnellInput]("black-forest-labs/flux-schnell")

/*
FluxDev is a model better suited for images containing text, though it is slightly more expensive and slower.
*/
var FluxDev = replicate.NewReplicateModel[FluxDevInput]("black-forest-labs/flux-dev")


// FluxSchnellInput represents input parameters for the Flux Schnell model.
// It contains various configuration options to control the image generation process.
type FluxSchnellInput struct {
	// Prompt is the primary text input that guides image generation. Use clear and specific language to achieve the best results.
	Prompt string `json:"prompt"`

	// AspectRatio defines the width-to-height ratio of the output image. Common values are "1:1", "16:9", etc., to control composition.
	AspectRatio *string `json:"aspect_ratio,omitempty"`

	// NumOutputs specifies how many images to generate per request. Adjust based on your need for variety versus cost.
	NumOutputs *int `json:"num_outputs,omitempty"`

	// NumInferenceSteps controls the number of steps in the diffusion process. Higher values yield finer detail but increase runtime.
	NumInferenceSteps *int `json:"num_inference_steps,omitempty"`

	// Seed sets the random number seed for reproducibility. Use a fixed seed to regenerate the same output consistently.
	Seed *int `json:"seed,omitempty"`

	// OutputFormat determines the file format of the generated image, such as "png" or "jpeg".
	OutputFormat *string `json:"output_format,omitempty"`

	// OutputQuality sets the quality level of the output image, typically on a scale from 1 (lowest) to 100 (highest).
	OutputQuality *int `json:"output_quality,omitempty"`

	// GoFast enables a faster but potentially less precise generation mode, balancing speed and output fidelity.
	GoFast *bool `json:"go_fast,omitempty"`

	// Megapixels specifies the desired resolution in megapixels, influencing image size and detail.
	Megapixels *string `json:"megapixels,omitempty"`

	// DisableSafetyChecker disables built-in content filters. Use with caution to avoid inappropriate content generation.
	DisableSafetyChecker *bool `json:"disable_safety_checker,omitempty"`
}

// FluxDevInput represents input parameters for the Flux Dev model.
// It contains various configuration options to control the image generation process.
type FluxDevInput struct {
	// Prompt is the main input text that guides image generation.
	Prompt string `json:"prompt"`

	// Seed is the random seed for reproducibility of results. Use -1 for random seed.
	Seed *int `json:"seed,omitempty"`

	// Image is the input image for image to image mode. The aspect ratio of your output will match this image.
	Image *string `json:"image,omitempty"`

	// AspectRatio defines the aspect ratio of the output image.
	// Options include "1:1", "16:9", "21:9", "3:2", "2:3", "4:5", "5:4", "3:4", "4:3", "9:16", "9:21".
	AspectRatio *string `json:"aspect_ratio,omitempty"`

	// NumOutputs specifies how many images to generate per request.
	NumOutputs *int `json:"num_outputs,omitempty"`

	// PromptStrength is the prompt strength when using img2img. 1.0 corresponds to full destruction of information in image.
	PromptStrength *float64 `json:"prompt_strength,omitempty"`

	// NumInferenceSteps sets the number of inference steps to perform during generation.
	NumInferenceSteps *int `json:"num_inference_steps,omitempty"`

	// Guidance is the guidance scale used to influence image generation.
	// Lower values can give more realistic images. Good values to try are 2, 2.5, 3 and 3.5.
	Guidance *float64 `json:"guidance,omitempty"`

	// DisableSafetyChecker disables built-in content filters. Use with caution.
	DisableSafetyChecker *bool `json:"disable_safety_checker,omitempty"`

	// GoFast enables faster predictions with additional optimizations.
	GoFast *bool `json:"go_fast,omitempty"`

	// Megapixels is the approximate number of megapixels for generated image, e.g. "1" or "0.25".
	Megapixels *string `json:"megapixels,omitempty"`

	// OutputFormat sets the file format of the generated image.
	// Supported formats are "webp", "jpg", and "png".
	OutputFormat *string `json:"output_format,omitempty"`

	// OutputQuality determines the output quality for jpg and webp formats, ranging from 0 (lowest) to 100 (highest).
	// Not relevant for png outputs.
	OutputQuality *int `json:"output_quality,omitempty"`
}
