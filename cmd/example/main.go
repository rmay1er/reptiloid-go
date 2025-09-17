package main

import (
	"fmt"
	"os"

	"github.com/rmay1er/reptiloid-go/reptiloid"
	models "github.com/rmay1er/reptiloid-go/reptiloid/models/image"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get API Key from environment variable
	var apiKey = os.Getenv("REPLICATE_API_TOKEN")

	client := reptiloid.NewClient(&models.FluxSchnell, apiKey)

	// Construct model input
	input := &models.FluxSchnellInput{
		Prompt:      "Blue pineapple with glowing slime around",
		AspectRatio: "16:9",
	}

	image, _ := client.Generate(input)

	if err != nil {
		fmt.Printf("Error generating image: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", &input)
	fmt.Printf("%+v\n", image.Output[0])
}
