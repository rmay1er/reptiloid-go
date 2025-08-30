package main

import (
	"fmt"
	"reptiloid-go/internal/replicate"
	"reptiloid-go/internal/replicate/models"
	"os"

	"github.com/joho/godotenv"
)

func main() {

  // Load .env file
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  var apiKey = os.Getenv("REPLICATE_API_TOKEN")

	//Init Model
	var newmod = models.FluxSchnell

	//Create Client
	client := replicate.NewClient(newmod, apiKey)

	//Construct model input
	input := models.FluxSchnellInput{
		Prompt:      "Orange apple on a wooden table, photorealistic, high detail",
	}

	//Generate image
	image, err := client.GenerateImage(input)
	if err != nil {
    fmt.Printf("Error generating image: %v\n", err)
    return
  }


	fmt.Printf("%+v\n", image)
}
