package handlers

import (
	"encoding/json"
	"fmt"
	"my-golang-react-app/db"
	"net/http"
)

type Booking struct {
	RoomID     int    `json:"room_id"`
	TimeslotID int    `json:"timeslot_id"`
	Date       string `json:"date"` // format: "YYYY-MM-DD"
}


// CreateBooking godoc
// @Summary Create a booking
// @Description Add a new booking for a specific room and timeslot
// @Tags bookings
// @Accept json
// @Produce plain
// @Param booking body Booking true "Booking info"
// @Success 201 {string} string "Booking created"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Failed to insert booking"
// @Router /api/bookings [post]
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	fmt.Println("📥 POST /api/bookings called")

	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(
		`INSERT INTO bookings (room_id, booking_date, timeslot_id)
		 VALUES ($1, $2, $3)`,
		booking.RoomID, booking.Date, booking.TimeslotID,
	)
	if err != nil {
		http.Error(w, "Failed to insert booking: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Booking created"))
}

// GetRooms godoc
// @Summary Get all rooms
// @Tags rooms
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {string} string "Failed to fetch rooms"
// @Router /api/rooms [get]
func GetRooms(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, image_url FROM rooms")
	if err != nil {
		http.Error(w, "Failed to fetch rooms", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rooms []map[string]interface{}
	for rows.Next() {
		var id int
		var name, imageURL string
		if err := rows.Scan(&id, &name, &imageURL); err != nil {
			continue
		}
		rooms = append(rooms, map[string]interface{}{
			"id":        id,
			"name":      name,
			"image_url": imageURL,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

func GetTimeSlots(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, time_label FROM timeslots")
	if err != nil {
		http.Error(w, "Failed to fetch time slots", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var slots []map[string]interface{}
	for rows.Next() {
		var id int
		var label string
		if err := rows.Scan(&id, &label); err != nil {
			continue
		}
		slots = append(slots, map[string]interface{}{
			"id":         id,
			"time_label": label,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}


func GetBookingsByDate(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	if date == "" {
		http.Error(w, "Missing date query param", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query(`
		SELECT b.id, b.room_id, r.name, b.booking_date, b.timeslot_id, t.time_label
		FROM bookings b
		JOIN rooms r ON b.room_id = r.id
		JOIN timeslots t ON b.timeslot_id = t.id
		WHERE b.booking_date = $1
	`, date)

	if err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []map[string]interface{}
	for rows.Next() {
		var id, roomID, timeslotID int
		var roomName, timeLabel, bookingDate string
		if err := rows.Scan(&id, &roomID, &roomName, &bookingDate, &timeslotID, &timeLabel); err != nil {
			continue
		}
		bookings = append(bookings, map[string]interface{}{
			"id":           id,
			"room_id":      roomID,
			"room_name":    roomName,
			"date":         bookingDate,
			"timeslot_id":  timeslotID,
			"time_label":   timeLabel,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
		SELECT b.id, b.room_id, r.name, b.booking_date, b.timeslot_id, t.time_label, b.created_at
		FROM bookings b
		JOIN rooms r ON b.room_id = r.id
		JOIN timeslots t ON b.timeslot_id = t.id
		ORDER BY b.created_at DESC
	`)

	if err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []map[string]interface{}
	for rows.Next() {
		var id, roomID, timeslotID int
		var roomName, timeLabel string
		var bookingDate, createdAt string
		if err := rows.Scan(&id, &roomID, &roomName, &bookingDate, &timeslotID, &timeLabel, &createdAt); err != nil {
			continue
		}
		bookings = append(bookings, map[string]interface{}{
			"id":          id,
			"room_name":   roomName,
			"date":        bookingDate,
			"time_label":  timeLabel,
			"created_at":  createdAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}


