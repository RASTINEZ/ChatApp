package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"my-golang-react-app/db" // Assuming your db connection is here
	"net/http"
	"strconv"
	"time"
	// You might need a library for handling URL parameters if you get more complex
	// e.g., "github.com/gorilla/mux" if you switch from http.ServeMux
)

// MarketingCampaign defines the structure for marketing campaign data
type MarketingCampaign struct {
	CampaignID     int             `json:"campaign_id"`
	CampaignName   string          `json:"campaign_name"`
	StartDate      string          `json:"start_date"` // Consider using time.Time for internal logic
	EndDate        string          `json:"end_date"`   // Consider using time.Time for internal logic
	Budget         float64         `json:"budget"`
	TargetAudience string          `json:"target_audience"`
	Channel        string          `json:"channel"`
	Status         string          `json:"status"`
	Goal           string          `json:"goal"`
	KpiMetricName  sql.NullString  `json:"kpi_metric_name"`  // Use sql.NullString for nullable fields
	KpiTargetValue sql.NullFloat64 `json:"kpi_target_value"` // Use sql.NullFloat64 for nullable fields
	KpiActualValue sql.NullFloat64 `json:"kpi_actual_value"` // Use sql.NullFloat64 for nullable fields
	Notes          sql.NullString  `json:"notes"`            // Use sql.NullString for nullable fields
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// Helper to convert sql.NullString to string for JSON response
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return "" // Or return null, depending on your API preference
}

// Helper to convert sql.NullFloat64 to *float64 for JSON response
func nullFloat64ToPtr(nf sql.NullFloat64) *float64 {
	if nf.Valid {
		return &nf.Float64
	}
	return nil
}

// CreateMarketingCampaign godoc
// @Summary Create a marketing campaign
// @Description Add a new marketing campaign
// @Tags marketing
// @Accept json
// @Produce json
// @Param campaign body MarketingCampaign true "Campaign info"
// @Success 201 {object} MarketingCampaign
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Failed to create campaign"
// @Router /api/marketing/campaigns [post]
func CreateMarketingCampaign(w http.ResponseWriter, r *http.Request) {
	var mc MarketingCampaign
	fmt.Println("📥 POST /api/marketing/campaigns called")

	if err := json.NewDecoder(r.Body).Decode(&mc); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Apply default values and basic validation
	if mc.CampaignName == "" {
		http.Error(w, "Invalid request: CampaignName is required", http.StatusBadRequest)
		return
	}
	if mc.StartDate == "" {
		http.Error(w, "Invalid request: StartDate is required", http.StatusBadRequest)
		return
	}
	if mc.EndDate == "" {
		http.Error(w, "Invalid request: EndDate is required", http.StatusBadRequest)
		return
	}

	// Set default for Status if not provided
	if mc.Status == "" {
		mc.Status = "planned" // Default status as per tool definition
	}

	// Set defaults for other potentially required string fields if they are empty
	// This helps prevent NOT NULL constraint errors if the DB doesn't allow empty strings
	// and the AI omits them (though they are 'required' in the tool definition).
	if mc.TargetAudience == "" {
		// Depending on your business logic and DB constraints,
		// you might set a default or return an error if it's truly mandatory.
		// For now, setting a placeholder to avoid DB errors for NOT NULL columns.
		mc.TargetAudience = "Not Specified" // Example placeholder
		// Or, if it's truly required and "" is not acceptable:
		// http.Error(w, "Invalid request: TargetAudience is required", http.StatusBadRequest)
		// return
	}
	if mc.Channel == "" {
		mc.Channel = "Not Specified" // Example placeholder
		// Or, if it's truly required:
		// http.Error(w, "Invalid request: Channel is required", http.StatusBadRequest)
		// return
	}
	if mc.Goal == "" {
		mc.Goal = "Not Specified" // Example placeholder
		// Or, if it's truly required:
		// http.Error(w, "Invalid request: Goal is required", http.StatusBadRequest)
		// return
	}
	// Note: For 'budget' (a float64), if not provided, it defaults to 0.0.
	// If 0.0 is not acceptable for a NOT NULL 'budget' column,
	// you'd need specific validation or a different default strategy for it.
	// The tool definition marks budget as required, so the AI should ideally send it.

	// Parse dates - assuming "YYYY-MM-DD" format from input
	// ...existing code...
	query := `
        INSERT INTO marketing_campaigns 
        (campaign_name, start_date, end_date, budget, target_audience, channel, status, goal, kpi_metric_name, kpi_target_value, notes)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING campaign_id, created_at, updated_at, status;
    `
	err := db.DB.QueryRow(
		query,
		mc.CampaignName, mc.StartDate, mc.EndDate, mc.Budget, mc.TargetAudience, mc.Channel,
		mc.Status, mc.Goal, mc.KpiMetricName, mc.KpiTargetValue, mc.Notes,
	).Scan(&mc.CampaignID, &mc.CreatedAt, &mc.UpdatedAt, &mc.Status) // Read back the generated ID and timestamps

	if err != nil {
		http.Error(w, "Failed to insert campaign: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mc)
}

// GetMarketingCampaigns godoc
// @Summary Get all marketing campaigns
// @Description Retrieve a list of all marketing campaigns
// @Tags marketing
// @Produce json
// @Success 200 {array} MarketingCampaign
// @Failure 500 {string} string "Failed to fetch campaigns"
// @Router /api/marketing/campaigns [get]
func GetMarketingCampaigns(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /api/marketing/campaigns called")

	// Optional: Add query param filters here (e.g., ?status=active, ?channel=Email)
	// statusFilter := r.URL.Query().Get("status")
	// query := "SELECT campaign_id, campaign_name, start_date, end_date, budget, target_audience, channel, status, goal, kpi_metric_name, kpi_target_value, kpi_actual_value, notes, created_at, updated_at FROM marketing_campaigns"
	// var args []interface{}
	// if statusFilter != "" {
	// query += " WHERE status = $1"
	// args = append(args, statusFilter)
	// }
	// query += " ORDER BY created_at DESC"

	rows, err := db.DB.Query(`
		SELECT campaign_id, campaign_name, 
		       TO_CHAR(start_date, 'YYYY-MM-DD') as start_date, 
		       TO_CHAR(end_date, 'YYYY-MM-DD') as end_date, 
		       budget, target_audience, channel, status, goal, 
		       kpi_metric_name, kpi_target_value, kpi_actual_value, 
		       notes, created_at, updated_at 
		FROM marketing_campaigns
		ORDER BY created_at DESC
	`) // TO_CHAR used to ensure date format consistency
	if err != nil {
		http.Error(w, "Failed to fetch campaigns: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var campaigns []MarketingCampaign
	for rows.Next() {
		var mc MarketingCampaign
		var startDateStr, endDateStr string // Temporary strings for scanning dates

		if err := rows.Scan(
			&mc.CampaignID, &mc.CampaignName, &startDateStr, &endDateStr,
			&mc.Budget, &mc.TargetAudience, &mc.Channel, &mc.Status, &mc.Goal,
			&mc.KpiMetricName, &mc.KpiTargetValue, &mc.KpiActualValue, &mc.Notes,
			&mc.CreatedAt, &mc.UpdatedAt,
		); err != nil {
			fmt.Printf("Error scanning campaign row: %v\n", err) // Log error
			// Decide if you want to skip or return an error for the whole request
			continue
		}
		mc.StartDate = startDateStr
		mc.EndDate = endDateStr
		campaigns = append(campaigns, mc)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating campaign rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(campaigns)
}

// GetMarketingCampaignByID godoc
// @Summary Get a specific marketing campaign by ID
// @Description Retrieve details of a marketing campaign by its ID
// @Tags marketing
// @Produce json
// @Param id path int true "Campaign ID"
// @Success 200 {object} MarketingCampaign
// @Failure 400 {string} string "Invalid campaign ID"
// @Failure 404 {string} string "Campaign not found"
// @Failure 500 {string} string "Failed to fetch campaign"
// @Router /api/marketing/campaigns/{id} [get]
func GetMarketingCampaignByID(w http.ResponseWriter, r *http.Request) {
	// This requires a router that can handle path parameters like {id}
	// http.ServeMux doesn't support this directly in a clean way.
	// You'd typically use a library like github.com/gorilla/mux or chi.
	// For a simple approach with ServeMux, you might parse the ID from the path string.
	// Example (crude, assumes path is /api/marketing/campaigns/123):
	// pathParts := strings.Split(r.URL.Path, "/")
	// idStr := pathParts[len(pathParts)-1]
	// For now, let's assume you'll adapt this or use query params.
	// We'll use a query param for simplicity with current ServeMux setup.

	idStr := r.URL.Query().Get("id")
	fmt.Printf("GET /api/marketing/campaigns?id=%s called\n", idStr)

	if idStr == "" {
		http.Error(w, "Campaign ID is required as a query parameter (e.g., /api/marketing/campaigns/details?id=1)", http.StatusBadRequest)
		return
	}

	campaignID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid campaign ID format", http.StatusBadRequest)
		return
	}

	var mc MarketingCampaign
	var startDateStr, endDateStr string

	query := `
		SELECT campaign_id, campaign_name, 
		       TO_CHAR(start_date, 'YYYY-MM-DD') as start_date, 
		       TO_CHAR(end_date, 'YYYY-MM-DD') as end_date, 
		       budget, target_audience, channel, status, goal, 
		       kpi_metric_name, kpi_target_value, kpi_actual_value, 
		       notes, created_at, updated_at 
		FROM marketing_campaigns 
		WHERE campaign_id = $1
	`
	err = db.DB.QueryRow(query, campaignID).Scan(
		&mc.CampaignID, &mc.CampaignName, &startDateStr, &endDateStr,
		&mc.Budget, &mc.TargetAudience, &mc.Channel, &mc.Status, &mc.Goal,
		&mc.KpiMetricName, &mc.KpiTargetValue, &mc.KpiActualValue, &mc.Notes,
		&mc.CreatedAt, &mc.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Campaign not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch campaign: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	mc.StartDate = startDateStr
	mc.EndDate = endDateStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mc)
}

// UpdateMarketingCampaign godoc
// @Summary Update a marketing campaign
// @Description Update details of an existing marketing campaign by its ID
// @Tags marketing
// @Accept json
// @Produce json
// @Param id path int true "Campaign ID"  // This would ideally be a path param
// @Param campaign body MarketingCampaign true "Campaign info to update"
// @Success 200 {object} MarketingCampaign
// @Failure 400 {string} string "Invalid request or campaign ID"
// @Failure 404 {string} string "Campaign not found"
// @Failure 500 {string} string "Failed to update campaign"
// @Router /api/marketing/campaigns/{id} [put] // Assuming {id} will be handled
func UpdateMarketingCampaign(w http.ResponseWriter, r *http.Request) {
	// Similar to GetMarketingCampaignByID, handling ID from path is tricky with ServeMux.
	// Let's assume ID comes as a query param for this example.
	idStr := r.URL.Query().Get("id") // e.g. /api/marketing/campaigns/update?id=1
	fmt.Printf("PUT /api/marketing/campaigns/update?id=%s called\n", idStr)

	if idStr == "" {
		http.Error(w, "Campaign ID is required as a query parameter for update", http.StatusBadRequest)
		return
	}
	campaignID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid campaign ID format", http.StatusBadRequest)
		return
	}

	var mc MarketingCampaign
	if err := json.NewDecoder(r.Body).Decode(&mc); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure the ID from query param matches any ID in body if present, or just use query param ID
	mc.CampaignID = campaignID

	query := `
		UPDATE marketing_campaigns SET 
		campaign_name = $1, start_date = $2, end_date = $3, budget = $4, 
		target_audience = $5, channel = $6, status = $7, goal = $8, 
		kpi_metric_name = $9, kpi_target_value = $10, kpi_actual_value = $11, notes = $12,
		updated_at = CURRENT_TIMESTAMP
		WHERE campaign_id = $13
		RETURNING campaign_name, start_date, end_date, budget, target_audience, channel, status, goal, kpi_metric_name, kpi_target_value, kpi_actual_value, notes, created_at, updated_at;
	`
	// Note: kpi_actual_value is included in the update.
	// You might want separate logic if it's only updated by a different process.
	var startDateStr, endDateStr string
	err = db.DB.QueryRow(
		query,
		mc.CampaignName, mc.StartDate, mc.EndDate, mc.Budget, mc.TargetAudience,
		mc.Channel, mc.Status, mc.Goal, mc.KpiMetricName, mc.KpiTargetValue,
		mc.KpiActualValue, mc.Notes, mc.CampaignID,
	).Scan(
		&mc.CampaignName, &startDateStr, &endDateStr, &mc.Budget, &mc.TargetAudience,
		&mc.Channel, &mc.Status, &mc.Goal, &mc.KpiMetricName, &mc.KpiTargetValue,
		&mc.KpiActualValue, &mc.Notes, &mc.CreatedAt, &mc.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows { // Might not happen with RETURNING if update affects 0 rows
			http.Error(w, "Campaign not found or no update occurred", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update campaign: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	mc.StartDate = startDateStr
	mc.EndDate = endDateStr

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mc)
}

// DeleteMarketingCampaign godoc
// @Summary Delete a marketing campaign by ID
// @Description Remove a marketing campaign from the database
// @Tags marketing
// @Produce plain
// @Param id path int true "Campaign ID" // This would ideally be a path param
// @Success 200 {string} string "Campaign deleted successfully"
// @Failure 400 {string} string "Invalid campaign ID"
// @Failure 404 {string} string "Campaign not found"
// @Failure 500 {string} string "Failed to delete campaign"
// @Router /api/marketing/campaigns/{id} [delete]
func DeleteMarketingCampaign(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // e.g. /api/marketing/campaigns/delete?id=1
	fmt.Printf("DELETE /api/marketing/campaigns/delete?id=%s called\n", idStr)

	if idStr == "" {
		http.Error(w, "Campaign ID is required as a query parameter for delete", http.StatusBadRequest)
		return
	}
	campaignID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid campaign ID format", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("DELETE FROM marketing_campaigns WHERE campaign_id = $1", campaignID)
	if err != nil {
		http.Error(w, "Failed to delete campaign: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Campaign not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Campaign deleted successfully"))
}
