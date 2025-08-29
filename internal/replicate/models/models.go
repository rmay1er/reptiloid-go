package models

import "nerualchechevica-bot/internal/replicate"

/*
FluxSchnell is a fast model optimized for quick generation.
*/
var FluxSchnell = replicate.NewReplicateModel[FluxSchnellInput]("black-forest-labs/flux-schnell")

/*
FluxDev is a model better suited for images containing text, though it is slightly more expensive and slower.
*/
var FluxDev = replicate.NewReplicateModel[FluxDevInput]("black-forest-labs/flux-dev")

// FluxSchnellInput represents input parameters for the Flux Schnell model.
// Prompt is the main input text.
// AspectRatio specifies the desired aspect ratio of the output.
type FluxSchnellInput struct {
	Prompt      string `json:"prompt"`
	AspectRatio string `json:"aspect_ratio"`
}

// FluxDevInput represents input parameters for the Flux Dev model.
// Prompt is the main input text.
// Size specifies the desired size of the output.
type FluxDevInput struct {
	Prompt string `json:"prompt"`
	Size   string `json:"size"`
}
