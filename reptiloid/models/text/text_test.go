package text_test

import (
	"os"
	// "strings"
	"testing"

	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joho/godotenv"
	"github.com/rmay1er/reptiloid-go/reptiloid"
	"github.com/rmay1er/reptiloid-go/reptiloid/models/text"
)

func TestTextModel(t *testing.T) {

	err := godotenv.Load("../../../.env") // Load .env file if exists
	if err != nil {
		t.Log("No .env file found, proceeding with environment variables")
	}

	// Replace with your actual API key for testing purposes
	apiKey := os.Getenv("REPLICATE_API_TOKEN")

	// Initialize the GPT-5 nano model
	model := text.GPT5nano

	// Create a new client with the model and API key
	client := reptiloid.NewClient(&model, apiKey)

	// Define the input for the model
	input := text.GPT5SeriesInput{
		Prompt:       "If you get this, respond with somthing funny. about it on russian language",
		SystemPrompt: "Answer like king of Orcs",
	}

	// Generate a response from the model
	output, err := client.Generate(&input)
	if err != nil {
		t.Fatalf("Error generating response: %v", err)
	}

	// Check if the output is not empty
	if len(output.Output) == 0 {
		t.Fatal("Expected non-empty output, got empty")
	}

	// Print the output for manual verification (optional)
	t.Logf("Model output: %+v", output)
}
