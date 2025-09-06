package image_test

import (
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joho/godotenv"
	"github.com/rmay1er/reptiloid-go/reptiloid"
	"github.com/rmay1er/reptiloid-go/reptiloid/models/image"
)

func TestImageModel(t *testing.T) {

	err := godotenv.Load("../../../.env") // Load .env file if exists
	if err != nil {
		t.Log("No .env file found, proceeding with environment variables")
	}

	// Get API Key from environment variable
	apiKey := os.Getenv("REPLICATE_API_TOKEN")

	// Initialize the FluxSchnell model
	model := image.FluxDev

	// Create a new client with the model and API key
	client := reptiloid.NewClient(&model, apiKey)

	// Define the input for the model
	input := image.FluxDevInput{
		Prompt:      "A fantasy landscape, trending on artstation",
		AspectRatio: aws.String("16:9"),
	}

	// Generate an image from the model
	output, err := client.Generate(&input)
	if err != nil {
		t.Fatalf("Error generating image: %v", err)
	}

	// Check if the output is not empty
	if len(output.Output) == 0 {
		t.Fatal("Expected non-empty output, got empty")
	}

	// Print the output for manual verification (optional)
	t.Logf("Image output: %+s", strings.Join(output.Output, ""))
}
