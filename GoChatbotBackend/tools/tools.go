package tools

import "github.com/google/generative-ai-go/genai"

// MarketingToolDeclarations will be defined in marketingTools.go (ensure it's exported or in same package)
// BookingToolDeclarations will be defined in bookingTools.go (ensure it's exported or in same package)

// GetAllTools combines all available tool declarations
// This function assumes MarketingToolDeclarations and BookingToolDeclarations are accessible
// (e.g., they are exported vars in their respective files within the same 'tools' package).
func GetAllTools() []*genai.Tool {
	var allTools []*genai.Tool
	
	// These variables would be defined in marketingTools.go and bookingTools.go
	// and be visible if they are in the same 'tools' package.
	// Example:
	// In marketingTools.go: var MarketingToolDeclarations = &genai.Tool{...}
	// In bookingTools.go:   var BookingToolDeclarations = &genai.Tool{...}


	// To make them truly modular and accessible here if they are unexported or in sub-packages
	// you might have functions in those files that return them:
	// e.g. GetMarketingTools() *genai.Tool
	
	// For simplicity, assuming they are package-level vars in the 'tools' package:
	if MarketingToolDeclarations != nil { // Defined in marketingTools.go
		allTools = append(allTools, MarketingToolDeclarations)
	}
	if BookingToolDeclarations != nil { // Defined in bookingTools.go
		allTools = append(allTools, BookingToolDeclarations)
	}
	return allTools
}