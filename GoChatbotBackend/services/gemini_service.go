// Example (conceptual - refer to official SDK docs for exact implementation)
// In a new file, e.g., services/gemini_service.go

package services

import (
	"context"
	
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var geminiClient *genai.GenerativeModel

func InitGemini(modelName string) { // Allow model name to be passed in
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set.")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating Gemini client: %v", err)
	}
	// Not closing client here, assuming it's a long-lived client for the application

	// Use the modelName parameter here
	if modelName == "" {
		modelName = "gemini-1.5-flash-latest" // Default to Flash if not specified
		log.Printf("No model name provided, defaulting to %s", modelName)
	}
	
	geminiClient = client.GenerativeModel(modelName)
	log.Printf("Gemini client initialized successfully with model: %s", modelName)
}

// GetGeminiClient provides access to the initialized client.
// You might need to handle the case where client.Close() from InitGemini would make this unusable
// if InitGemini is not a true singleton initializer.
// A better pattern for a shared client would be a sync.Once or a global var with explicit Close on app shutdown.
func GetGeminiClient() *genai.GenerativeModel {
    if geminiClient == nil {
        log.Println("Gemini client not initialized. Attempting to initialize with default Flash model.")
        InitGemini("gemini-1.5-flash-latest") // Initialize with Flash if called before explicit Init
        if geminiClient == nil {
            log.Fatal("Failed to initialize Gemini client on demand.")
        }
    }
	return geminiClient
}
