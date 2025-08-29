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
	var newmod = models.FluxDev

	//Create Client
	client := replicate.NewClient(newmod, apiKey)

	//Construct model input
	input := models.FluxDevInput{
		Prompt:      "Orange apple on a wooden table, photorealistic, high detail",
	}

	//Generate image
	image, _ := client.GenerateImage(input)


	fmt.Printf("%+v\n", image.Output[0])
}
