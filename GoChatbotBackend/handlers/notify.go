package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PushRequest struct {
	Title   string `json:"title"`
	Message string `json:"body"` // Changed to match frontend "body" key
}

func SendPushNotification(w http.ResponseWriter, r *http.Request) {
	var req PushRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Prepare payload for OneSignal
	payload := map[string]interface{}{
		"app_id": os.Getenv("ONESIGNAL_APP_ID"), // Or hardcode temporarily
		"included_segments": []string{"All"},
		"headings": map[string]string{
			"en": req.Title,
		},
		"contents": map[string]string{
			"en": req.Message,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error preparing request: %v", err), http.StatusInternalServerError)
		return
	}

	// Create request to OneSignal API
	apiURL := "https://onesignal.com/api/v1/notifications"
	httpReq, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating request: %v", err), http.StatusInternalServerError)
		return
	}

	// Add headers
	httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpReq.Header.Set("Authorization", "Basic "+os.Getenv("ONESIGNAL_API_KEY")) // Or use raw string

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Handle non-2xx responses
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		http.Error(w, fmt.Sprintf("Failed to send push notification: status %v", resp.Status), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Push notification sent via OneSignal"))
}
