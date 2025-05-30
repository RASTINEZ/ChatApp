package tools

import "github.com/google/generative-ai-go/genai"

// BookingToolDeclarations defines tools for managing bookings.
var BookingToolDeclarations = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Name:        "getRooms",
			Description: "Retrieves a list of available meeting rooms.",
			Parameters:  &genai.Schema{Type: genai.TypeObject, Properties: map[string]*genai.Schema{}}, // No parameters for now
		},
		{
			Name:        "getTimeSlots",
			Description: "Retrieves available time slots for booking, optionally filtered by room ID and date. Current date is 2025-05-16.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"room_id": {
						Type:        genai.TypeInteger,
						Description: "The ID of the room to get time slots for. Optional.",
						Nullable:    true,
					},
					"date": {
						Type:        genai.TypeString,
						Description: "The date (YYYY-MM-DD) to get time slots for. If relative (e.g., 'today', 'tomorrow'), calculate from 2025-05-16. Defaults to today if not specified.",
						Nullable:    true,
					},
				},
			},
		},
		{
			Name:        "createBooking",
			Description: "Creates a new room booking for the user 'RASTINEZ'. Current date is 2025-05-16.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"room_id": {
						Type:        genai.TypeInteger,
						Description: "The ID of the room to book.",
					},
					"timeslot_id": {
						Type:        genai.TypeInteger,
						Description: "The ID of the timeslot to book.",
					},
					"date": {
						Type:        genai.TypeString,
						Description: "The date of the booking (YYYY-MM-DD). Calculate if relative to 2025-05-16.",
					},
					// "user_id": {Type: genai.TypeString, Description: "User ID, implicitly 'RASTINEZ'"}, // Implicit
				},
				Required: []string{"room_id", "timeslot_id", "date"},
			},
		},
		{
			Name:        "getAllBookings",
			Description: "Retrieves all bookings. Can be filtered by user ('RASTINEZ' is the current user) or date. Current date is 2025-05-16.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"user_id": {
						Type:        genai.TypeString,
						Description: "Filter bookings by user ID. For 'my bookings', use 'RASTINEZ'. Optional.",
						Nullable:    true,
					},
					"date": {
						Type:        genai.TypeString,
						Description: "Filter bookings by date (YYYY-MM-DD). Optional.",
						Nullable:    true,
					},
				},
			},
		},
		// Add GetBookingsByDate if it's a separate function
	},
}

// Add this to tools/marketingTools.go or a new tools/tools.go
/*
func GetAllTools() []*genai.Tool {
	var allTools []*genai.Tool
	if MarketingToolDeclarations != nil {
		allTools = append(allTools, MarketingToolDeclarations)
	}
	if BookingToolDeclarations != nil { // Add this check
		allTools = append(allTools, BookingToolDeclarations)
	}
	return allTools
}
*/
// If you put GetAllTools in a separate tools.go, then import it in chat_handler.go
// from "my-golang-react-app/tools" and call tools.GetAllTools()