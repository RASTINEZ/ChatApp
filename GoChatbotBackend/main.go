package main

import (
	"fmt"
	"log"
	"my-golang-react-app/db"
	"my-golang-react-app/routes"
	"my-golang-react-app/services" // <-- 1. IMPORT THE SERVICES PACKAGE
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "my-golang-react-app/docs" // For Swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// @title Booking API
	// @version 1.0
	// @description This is an API for managing room bookings and interacting with AI.
	// @host localhost:5000
	// @BasePath /

	err := godotenv.Load()
	if err != nil {
		// It's often okay if .env is not found, especially in production where env vars are set directly.
		// log.Fatal("❌ Error loading .env file")
		log.Println("⚠️ .env file not found, ensure environment variables are set if needed.")
	}

	// --- DB CONNECT WITH RETRY LOGIC ---
	const maxAttempts = 10
	const delay = 2 * time.Second
	var dbErr error

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		dbErr = db.InitDB()
		if dbErr == nil {
			fmt.Println("✅ Successfully connected to the database!")
			break
		}
		log.Printf("❌ Failed to connect to database (attempt %d/%d): %v", attempts, maxAttempts, dbErr)
		if attempts < maxAttempts {
			log.Printf("⏳ Waiting %v before retrying...", delay)
			time.Sleep(delay)
		}
	}
	if dbErr != nil {
		log.Fatalf("❌ Could not connect to database after %d attempts: %v", maxAttempts, dbErr)
		os.Exit(1)
	}

	// <-- 2. INITIALIZE GEMINI SERVICE -->
	// Choose your desired model. For "basic AI like flash", use:
	geminiModelToUse := "gemini-1.5-flash-latest"
	// Or if you want to try gemini-pro:
	// geminiModelToUse := "gemini-pro"
	services.InitGemini(geminiModelToUse)
	// No error return from InitGemini as it currently uses log.Fatal on error.
	// If you change InitGemini to return an error, handle it here:
	// if err := services.InitGemini(geminiModelToUse); err != nil {
	// 	log.Fatalf("❌ Failed to initialize Gemini client: %v", err)
	// }

	mux := http.NewServeMux()

	// Register your routes
	routes.RegisterChatbotRoutes(mux) // Assuming ChatHandler here will use GetGeminiClient()
	routes.RegisterNotifyRoutes(mux)
	routes.RegisterBookingRoutes(mux)
	routes.RegisterMarketingRoutes(mux) // Make sure this is also registered if you created it

	// Swagger endpoint
	// Note: If your BasePath is /api, swagger might need to be /api/swagger/
	// But your @BasePath in comments is /api, while mux.Handle is /swagger/
	// For now, keeping it as /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://192.168.7.195:3000"}, // Your frontend
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},            // Add PUT, DELETE if used by frontend
		AllowedHeaders: []string{"Content-Type", "Authorization"},                      // Add Authorization if you plan to use tokens
		// AllowCredentials: true, // If you need to send cookies or use session-based auth
	})

	handler := c.Handler(mux)

	fmt.Println("✅ Server is running on http://localhost:5000")
	fmt.Println("📚 Swagger UI available at http://localhost:5000/swagger/index.html")
	if err := http.ListenAndServe("0.0.0.0:5000", handler); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
