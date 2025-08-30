package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rmay1er/reptiloid-go/reptiloid"
	"github.com/rmay1er/reptiloid-go/reptiloid/models"
	"os"

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

	//Init Model
	var newmod = models.FluxSchnell

	//Create Client
	client := reptiloid.NewClient(newmod, apiKey)

	//Construct model input
	input := models.FluxSchnellInput{
		Prompt:      "Blue pineapple with glowing slime around",
		AspectRatio: aws.String("16:9"),
	}

	//Generate image
	image, err := client.GenerateImage(input)
	if err != nil {
		fmt.Printf("Error generating image: %v\n", err)
		return
	}

	fmt.Printf("%+v\n", image.Output[0])
}
