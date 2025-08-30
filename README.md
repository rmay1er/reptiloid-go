# Reptiloid Go - Simple Replicate API Client for Go

A friendly Go library for generating AI images using Replicate's models. No complex setup, just simple and clean code.

## 🚀 Quick Start

### Prerequisites

- Go 1.25 or later
- A [Replicate](https://replicate.com) account and API token

### Installation

```bash
go get github.com/rmay1er/reptiloid-go
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
    "github.com/rmay1er/reptiloid-go/reptiloid"
    "github.com/rmay1er/reptiloid-go/reptiloid/models"
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
        Prompt:      "A cute robot playing guitar in the park",
    }

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

## 📦 Available Models & Future Plans

### Supported Models



### Planned Models & Features

#### Text Models (coming soon)
- ⬜ **GPT-5**: Next-generation large language model with enhanced understanding and creativity.
- ⬜ **GPT-5-mini**: Lightweight version optimized for faster responses and lower costs.
- ⬜ **GPT-5-nano**: Ultra-compact model suitable for mobile and edge applications.
- ⬜ **GPT-5-structured**: Specialized for structured data generation and complex tasks.
- ⬜ **GPT-4.1**: Powerful language understanding and generation.
- ⬜ **GPT-4.1-mini**: Faster and cost-effective for chat-based tasks.
- ⬜ **GPT-4.1-nano**: Compact model designed for efficient inference.
- ⬜ **GPT-4o**: Optimized model for multi-turn conversations and reasoning.
- ⬜ **GPT-4o-mini**: Smaller, faster variant of GPT-4o for lightweight usage.
- ⬜ **Claude 4 Sonnet**: Advanced model with poetic and creative language skills.
- ⬜ **DeepSeek R1**: AI assistant specialized in search and information retrieval.
- ⬜ **DeepSeek V3**: Improved version of DeepSeek with expanded knowledge base.
- ⬜ **Claude 3.7 Haiku**: Focused on generating concise and expressive poetic forms.
- ⬜ **Claude 3.7 Sonnet**: Tailored for rich, structured poetry and complex writing.

#### Image Generation Models
- ✅ **FluxSchnell**: Fast image generation - great for quick experiments.
- ✅ **FluxDev**: Better at images with text, slightly slower but more precise.
- ⬜ **flux-kontext-pro**: Advanced Flux model with enhanced contextual understanding.
- ⬜ **ideogram-v3-turbo**: Ultra-fast image generation optimized for quick turnaround.
- ⬜ **ideogram-v3-quality**: High-fidelity images with detailed and nuanced output.
- ⬜ **ideogram-v3-balanced**: Balanced model combining speed and quality for versatile use.
- ⬜ **imagen-4 fast**: Google’s speedy Imagen variant for rapid prototyping.
- ⬜ **imagen-4**: Standard Imagen model offering top-tier image quality.
- ⬜ **imagen-4 ultra**: Ultra-high resolution and detail from Google’s Imagen technology.
- ⬜ **seedream-3**: Bytedance’s creative AI for generating imaginative and diverse images.

#### Video Models
- ⬜ Placeholder - video models support coming soon!

#### Utilities
- ⬜ **Whisper**: State-of-the-art speech-to-text transcription.
- ⬜ More utilities will be added in the future!

## ❓ Common Questions

### Where's my image?
The library returns a URL where your image is stored. You can download it or share it directly.

### How much does it cost?
Check [Replicate's pricing](https://replicate.com/pricing). Most images cost a few cents.


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
