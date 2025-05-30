package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"my-golang-react-app/db"
	"my-golang-react-app/services"
	"my-golang-react-app/tools" // Import your tools package
	"net/http"
	"net/url"
	"strconv"
	"strings" // Import strings
	
	"time"

	"github.com/google/generative-ai-go/genai"
)

// var (
// 	chatMemory = make(map[string][]string)
// 	memoryMu   sync.RWMutex
// )

// Message struct for incoming messages
type Message struct {
	Message string `json:"message"`
}

// ChatBotResponse struct for chatbot replies
type ChatBotResponse struct {
	Response string `json:"response"`
}

func storeMessage(sessionID, message string) error {
	dbConn := db.DB
	_, err := dbConn.Exec(
		"INSERT INTO chat_memory (session_id, message) VALUES ($1, $2)",
		sessionID, message,
	)
	return err
}

func fetchLastNMessages(sessionID string, n int) ([]string, error) {
	dbConn := db.DB
	rows, err := dbConn.Query(
		"SELECT message FROM chat_memory WHERE session_id = $1 ORDER BY id DESC LIMIT $2",
		sessionID, n,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []string
	for rows.Next() {
		var msg string
		if err := rows.Scan(&msg); err != nil {
			return nil, err
		}
		messages = append([]string{msg}, messages...)
	}
	return messages, nil
}

func clearOldMessages(sessionID string, keep int) error {
	dbConn := db.DB
	_, err := dbConn.Exec(`
        DELETE FROM chat_memory
        WHERE session_id = $1
          AND id NOT IN (
            SELECT id FROM chat_memory
            WHERE session_id = $1
            ORDER BY id DESC
            LIMIT $2
          )
    `, sessionID, keep)
	return err
}

// Helper function to make internal API calls
func callInternalAPI(method, path string, queryParams url.Values, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Construct URL
	fullURL := "http://localhost:5000" + path // Ensure this matches your server setup
	if queryParams != nil && len(queryParams) > 0 {
		fullURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequest(method, fullURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to %s: %w", fullURL, err)
	}
	req.Header.Set("Content-Type", "application/json")
	// TODO: Add any necessary auth headers if your internal APIs require them in the future

	client := &http.Client{Timeout: 15 * time.Second} // Increased timeout slightly
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request to %s: %w", fullURL, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from %s: %w", fullURL, err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API call to %s failed with status %d: %s", fullURL, resp.StatusCode, string(respBody))
	}
	return respBody, nil
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	// ... (Initial request validation and responseText declaration) ...
	var responseText string
	var msg Message // Ensure msg is declared before goto label

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	sessionID := r.Header.Get("X-Session-Id")
	if sessionID == "" {
		sessionID = "default"
	}
	if strings.TrimSpace(msg.Message) == "" {
		http.Error(w, "Message cannot be empty", http.StatusBadRequest)
		return
	}

	userMessageLower := strings.ToLower(msg.Message)

	// if strings.Contains(userMessageLower, "reset memory") {
	// 	memoryMu.Lock()
	// 	delete(chatMemory, sessionID)
	// 	memoryMu.Unlock()
	// 	responseText = "Your chat memory has been reset."
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(ChatBotResponse{Response: responseText})
	// 	return
	// }

	// ...existing code...

	switch {
	case strings.Contains(userMessageLower, "hello") || strings.Contains(userMessageLower, "สวัสดี"):
		responseText = "Hello! How can I assist you today? สวัสดีครับ/ค่ะ!"
	// ... other predefined cases ...
	case strings.Contains(userMessageLower, "book") || strings.Contains(userMessageLower, "จอง"):
		if userMessageLower == "book" || userMessageLower == "จอง" {
			responseText = "__show_booking__"
		} else {
			goto CallGeminiWithTools
		}
	default:
		goto CallGeminiWithTools
	}

	if responseText != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ChatBotResponse{Response: responseText})
		return
	}

CallGeminiWithTools:
	ctx := context.Background()
	baseGeminiModel := services.GetGeminiClient()
	if baseGeminiModel == nil {
		http.Error(w, "AI service is not available.", http.StatusInternalServerError)
		return
	}

	allAppTools := tools.GetAllTools()

	// --- Point 1: Configure Tools on the Model ---
	// CAVEAT: If GetGeminiClient() returns a shared singleton, this is not concurrency-safe.
	// A better approach would be for GetGeminiClient() to allow for configuration
	// or to return new instances. For now, we set it on the obtained model.
	baseGeminiModel.Tools = allAppTools // Tools are part of the GenerativeModel

	cs := baseGeminiModel.StartChat()
	// cs.Tools = allAppTools // THIS WAS THE ERROR - ChatSession doesn't have this field.

	// memoryMu.Lock()
	// chatMemory[sessionID] = append(chatMemory[sessionID], msg.Message)
	// history := strings.Join(chatMemory[sessionID], "\n")
	// memoryMu.Unlock()

	// contextualUserPrompt := fmt.Sprintf(
	// 	"SYSTEM CONTEXT: You are a helpful assistant. Current user is RASTINEZ. Current date is %s.\nCHAT HISTORY:\n%s\nUSER QUERY: %s",
	// 	time.Now().UTC().Format("2006-01-02"), history, msg.Message)

	// ...existing code...

	if strings.TrimSpace(userMessageLower) == "!clearoldmessages" {
		err := clearOldMessages(sessionID, 20) // Keep last 20 messages
		if err != nil {
			responseText = "Error clearing old messages: " + err.Error()
		} else {
			responseText = "Old messages cleared! Only the last 20 are kept."
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ChatBotResponse{Response: responseText})
		return
	}

	// ...existing code...
	if strings.Contains(userMessageLower, "!resetmemory") {
		// Delete all messages for this session from the DB
		dbConn := db.DB
		_, err := dbConn.Exec("DELETE FROM chat_memory WHERE session_id = $1", sessionID)
		if err != nil {
			responseText = "Error resetting memory."
		} else {
			responseText = "Your chat memory has been reset."
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ChatBotResponse{Response: responseText})
		return
	}

	// Store the new user message in the DB
	if err := storeMessage(sessionID, msg.Message); err != nil {
		http.Error(w, "Failed to store message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch the last N messages for context (e.g., N=20)
	historySlice, err := fetchLastNMessages(sessionID, 20)
	if err != nil {
		http.Error(w, "Failed to fetch chat history: "+err.Error(), http.StatusInternalServerError)
		return
	}
	history := strings.Join(historySlice, "\n")

	contextualUserPrompt := fmt.Sprintf(
		"SYSTEM CONTEXT: You are a helpful assistant. Current user is RASTINEZ. Current date is %s.\nCHAT HISTORY:\n%s\nUSER QUERY: %s",
		time.Now().UTC().Format("2006-01-02"), history, msg.Message)

	// History for the chat session
	// The ChatSession `cs` manages history internally after the first message.
	// We only need to send the current user message.

	var resp *genai.GenerateContentResponse
	// var err error

	resp, err = cs.SendMessage(ctx, genai.Text(contextualUserPrompt)) // Send just the text part

	if err != nil {
		fmt.Printf("Error generating content from Gemini (initial call): %v\n", err)
		responseText = "Sorry, I'm having trouble processing that with the AI."
	} else {
		if resp != nil && len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
			part := resp.Candidates[0].Content.Parts[0]
			if fc, ok := part.(genai.FunctionCall); ok {
				fmt.Printf("Gemini wants to call function: %s with args: %v\n", fc.Name, fc.Args)

				// --- Point 3: functionCallOutput should be map[string]any or specific struct for clarity ---
				var functionCallOutput map[string]any // More specific type for the output to Gemini
				var apiErr error

				queryParams := make(url.Values)
				// --- Point 2: fc.Args is already map[string]any ---
				requestBodyArgs := fc.Args // Use fc.Args directly

				switch fc.Name {
				// --- Marketing Tools ---
				case "getMarketingCampaigns":
					if status, ok := requestBodyArgs["status"].(string); ok && status != "" {
						queryParams.Set("status", status)
					}
					// ... (rest of the case)
					apiRespBody, localApiErr := callInternalAPI(http.MethodGet, "/api/marketing/campaigns", queryParams, nil)
					apiErr = localApiErr
					if apiErr == nil {
						// Try to unmarshal into a structured map for Gemini
						var structuredResp map[string]any
						if json.Unmarshal(apiRespBody, &structuredResp) == nil {
							functionCallOutput = structuredResp
						} else { // Fallback for non-JSON or malformed JSON
							functionCallOutput = map[string]any{"raw_response": string(apiRespBody)}
						}
					}

				case "createMarketingCampaign": // <<< ADD THIS CASE
					// requestBodyArgs already contains all parameters from the AI (budget, campaign_name, etc.)
					// We pass these directly as the body for the POST request.
					apiRespBody, localApiErr := callInternalAPI(http.MethodPost, "/api/marketing/campaigns", nil, requestBodyArgs)
					apiErr = localApiErr
					if apiErr == nil {
						// The CreateMarketingCampaign handler should return the created campaign as JSON
						var structuredResp map[string]any
						if json.Unmarshal(apiRespBody, &structuredResp) == nil {
							functionCallOutput = structuredResp // This will be the created campaign object
						} else {
							// Fallback if the response isn't the expected JSON object
							functionCallOutput = map[string]any{"raw_response": string(apiRespBody), "message": "Campaign might have been created, but response format was unexpected."}
						}
					}

				// ... (similar changes for other cases, ensuring functionCallOutput is map[string]any)
				case "updateMarketingCampaign":
					if idFloat, ok := requestBodyArgs["campaign_id"].(float64); ok {
						queryParams.Set("id", strconv.Itoa(int(idFloat)))

						// Create a new map for the body, excluding campaign_id
						bodyForUpdate := make(map[string]any)
						for k, v := range requestBodyArgs {
							if k != "campaign_id" {
								bodyForUpdate[k] = v
							}
						}
						apiRespBody, localApiErr := callInternalAPI(http.MethodPut, "/api/marketing/campaigns/update", queryParams, bodyForUpdate)
						apiErr = localApiErr
						if apiErr == nil {
							var structuredResp map[string]any
							if json.Unmarshal(apiRespBody, &structuredResp) == nil {
								functionCallOutput = structuredResp
							} else {
								functionCallOutput = map[string]any{"raw_response": string(apiRespBody)}
							}
						}
					} else {
						apiErr = fmt.Errorf("missing campaign_id for updateMarketingCampaign")
					}
				case "deleteMarketingCampaign":
					if idFloat, ok := requestBodyArgs["campaign_id"].(float64); ok {
						queryParams.Set("id", strconv.Itoa(int(idFloat)))
						apiRespBody, localApiErr := callInternalAPI(http.MethodDelete, "/api/marketing/campaigns/delete", queryParams, nil)
						apiErr = localApiErr
						if apiErr == nil {
							// For DELETE, construct a success message as Gemini expects a JSON object
							functionCallOutput = map[string]any{"success": true, "message": "Campaign deleted successfully. Response: " + string(apiRespBody)}
						}
					} else {
						apiErr = fmt.Errorf("missing campaign_id for deleteMarketingCampaign")
					}

				// ... (rest of your switch cases for marketing and booking tools) ...
				// Ensure all paths within the switch assign to functionCallOutput or apiErr
				default:
					apiErr = fmt.Errorf("unknown function call requested by AI: %s", fc.Name)
				}

				if apiErr != nil {
					fmt.Printf("Error executing tool %s: %v\n", fc.Name, apiErr)
					functionCallOutput = map[string]any{"error": apiErr.Error()} // Ensure it's a map
				}
				// If API call was successful but functionCallOutput wasn't set (e.g. API returns 204 No Content)
				if functionCallOutput == nil && apiErr == nil {
					functionCallOutput = map[string]any{"success": true, "message": "Operation completed successfully with no specific data returned."}
				}

				// --- Point 3 (continued): Sending FunctionResponse ---
				// The `Response` field of `genai.FunctionResponse` is `any`.
				// As long as `functionCallOutput` (which is now map[string]any)
				// can be serialized to a JSON object by the SDK, this should be fine.
				finalResp, err := cs.SendMessage(ctx, genai.FunctionResponse{Name: fc.Name, Response: functionCallOutput})
				// ... (rest of the logic for handling finalResp) ...
				if err != nil {
					fmt.Printf("Error generating content from Gemini (after function call): %v\n", err)
					responseText = "I tried to use one of my tools, but something went wrong processing the result."
				} else if finalResp != nil && len(finalResp.Candidates) > 0 && len(finalResp.Candidates[0].Content.Parts) > 0 {
					if textPart, ok := finalResp.Candidates[0].Content.Parts[0].(genai.Text); ok {
						responseText = string(textPart)
					} else {
						responseText = "I processed your request with my tools, but couldn't form a final text response."
					}
				} else {
					responseText = "I processed your request with my tools. Is there anything else I can help with?"
				}

			} else if textPart, ok := part.(genai.Text); ok {
				responseText = string(textPart)
			} else {
				responseText = "Sorry, I received an unexpected response type from the AI."
			}
		} else { // This else corresponds to "if resp != nil && len(resp.Candidates) > 0 ..."
			if resp != nil && resp.PromptFeedback != nil && resp.PromptFeedback.BlockReason != genai.BlockReasonUnspecified {
				responseText = fmt.Sprintf("My response was blocked. Reason: %s. Please rephrase your request.", resp.PromptFeedback.BlockReason.String())
			} else {
				responseText = "Sorry, I couldn't process your request at this time. No suitable response generated."
			}
		}
	}
	// ... (final response sending logic) ...
	if responseText == "" {
		fmt.Println("Warning: responseText was empty before sending. Setting default error message.")
		responseText = "Sorry, an unexpected error occurred while processing your request."
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChatBotResponse{Response: responseText})
}
