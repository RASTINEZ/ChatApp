package routes

import (
	"my-golang-react-app/handlers"
	"net/http"
)

func RegisterNotifyRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/send-push", handlers.SendPushNotification)
}

func RegisterChatbotRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/chat", handlers.ChatHandler)
}

func RegisterBookingRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/bookings", handlers.CreateBooking)
	mux.HandleFunc("/api/rooms", handlers.GetRooms)
	mux.HandleFunc("/api/timeslots", handlers.GetTimeSlots)
	mux.HandleFunc("/api/bookings/all", handlers.GetAllBookings) // GET all bookings
}

// RegisterMarketingRoutes sets up the routes for marketing campaign handlers
func RegisterMarketingRoutes(mux *http.ServeMux) {
	// POST to create, GET to list all
	mux.HandleFunc("/api/marketing/campaigns", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateMarketingCampaign(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetMarketingCampaigns(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Routes for specific campaign by ID (using query parameter `id`)
	// If you switch to a router supporting path variables, these would be cleaner e.g. /api/marketing/campaigns/{id}
	mux.HandleFunc("/api/marketing/campaigns/details", handlers.GetMarketingCampaignByID)            // Expects ?id=...
	mux.HandleFunc("/api/marketing/campaigns/update", func(w http.ResponseWriter, r *http.Request) { // Expects ?id=...
		if r.Method == http.MethodPut {
			handlers.UpdateMarketingCampaign(w, r)
		} else {
			http.Error(w, "Method not allowed, use PUT for update", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/marketing/campaigns/delete", func(w http.ResponseWriter, r *http.Request) { // Expects ?id=...
		if r.Method == http.MethodDelete {
			handlers.DeleteMarketingCampaign(w, r)
		} else {
			http.Error(w, "Method not allowed, use DELETE", http.StatusMethodNotAllowed)
		}
	})
}
