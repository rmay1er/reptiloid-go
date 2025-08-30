# Reptiloid Go - Simple Replicate API Client for Go

A beginner-friendly Go library for generating AI images using Replicate's models. No complex setup, just simple and clean code.

## 🚀 Quick Start

### Prerequisites

- Go 1.25 or later
- A [Replicate](https://replicate.com) account and API token

### Installation

```bash
go get github.com/your-username/reptiloid-go
```

### Get Your API Token

1. Sign up at [Replicate](https://replicate.com)
2. Go to your [API tokens page](https://replicate.com/account/api-tokens)
3. Copy your API token

### Your First AI Image

Create a `.env` file in your project:

```env
REPLICATE_API_TOKEN=your_token_here
```

Create `main.go`:

```go
package main

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "reptiloid-go/internal/replicate"
    "reptiloid-go/internal/replicate/models"
)

func main() {
    // Load environment variables
    godotenv.Load()

    // Get your API token
    apiKey := os.Getenv("REPLICATE_API_TOKEN")

    // Choose a model (FluxSchnell for fast generation)
    model := models.FluxSchnell

    // Create the client
    client := replicate.NewClient(model, apiKey)

    // Describe your image
    input := models.FluxSchnellInput{
        Prompt:      "A cute robot playing guitar in

    // Generate the image!
    result, err := client.GenerateImage(input)
    if err != nil {
        fmt.Printf("Oops! Something went wrong: %v\n", err)
        return
    }

    // Your image is ready!
    fmt.Printf("Image generated: %s\n", result.Output[0])
}
```

Run it:

```bash
go run main.go
```

## 📦 Available Models

### Flux Schnell 🚀
Fast image generation - perfect for quick experiments.

```go
model := models.FluxSchnell
input := models.FluxSchnellInput{
    Prompt: "A magical castle in the clouds",
    AspectRatio: "1:1",
    NumOutputs: 1,
}
```

### Flux Dev 📝
Better for text in images - slightly slower but more precise.

```go
model := models.FluxDev
input := models.FluxDevInput{
    Prompt: "A sign that says 'Welcome to AI World'",
    AspectRatio: "16:9",
    Seed: 42, // Same seed = same image every time
}
```

## 🎨 Customize Your Images

### Basic Options

- `Prompt`: Describe what you want to see
- `AspectRatio`: Image shape ("1:1", "16:9", "4:3")
- `NumOutputs`: How many images to create

### Advanced Options
- `Seed`: Use the same number to get the same image
- `OutputFormat`: "png", "jpg", or "webp"
- `OutputQuality`: 1-100 (higher = better quality)

## ❓ Common Questions

### Where's my image?
The library returns a URL where your image is stored. You can download it or share it directly.

### How much does it cost?
Check [Replicate's pricing](https://replicate.com/pricing). Most images cost a few cents.

### Why did it fail?
- Check your API token is correct
- Make sure you have enough credits
- Try a simpler prompt

## 🛠️ Development

### Adding New Models

Want to use a different Replicate model? Just add it to the models package!

```go
// In internal/replicate/models/models.go
var MyNewModel = replicate.NewReplicateModel[MyInput]("username/model-name")
```

### Running Tests

```bash
go test ./...
```

## 📚 Learn More

- [Replicate Documentation](https://replicate.com/docs)
- [Go Beginner Tutorial](https://go.dev/doc/tutorial/)

## 🤝 Contributing

Found a bug? Have an idea? We'd love your help!

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Submit a pull request

## 📄 License

MIT License - feel free to use this code for your projects!

---

Made with ❤️ for Go beginners exploring AI image generation
