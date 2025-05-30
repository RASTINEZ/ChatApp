package tools

import "github.com/google/generative-ai-go/genai"

// MarketingToolDeclarations defines the set of tools related to marketing campaigns.
var MarketingToolDeclarations = &genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		{
			Name:        "getMarketingCampaigns",
			Description: "Retrieves a list of marketing campaigns. Can be filtered by status (e.g., 'active', 'completed', 'planned'), or target audience. If no filters are provided, it returns all campaigns.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"status": {
						Type:        genai.TypeString,
						Description: "Filter campaigns by status (e.g., 'active', 'completed', 'planned'). Optional.",
						Nullable:    true,
					},
					"target_audience": {
						Type:        genai.TypeString,
						Description: "Filter campaigns by target audience. Optional.",
						Nullable:    true,
					},
					// Add other potential filters like limit, offset, date ranges if your API supports them
				},
			},
		},
		{
			Name:        "getMarketingCampaignByID",
			Description: "Retrieves a specific marketing campaign by its ID.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"campaign_id": {
						Type:        genai.TypeInteger, // Assuming campaign_id is an integer
						Description: "The unique ID of the marketing campaign.",
					},
				},
				Required: []string{"campaign_id"},
			},
		},
		{
			Name:        "createMarketingCampaign",
			Description: "Creates a new marketing campaign. The user 'RASTINEZ' is creating this. Current date is 2025-05-16.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"campaign_name":    {Type: genai.TypeString, Description: "Name of the campaign."},
					"start_date":       {Type: genai.TypeString, Description: "Start date (YYYY-MM-DD). If relative (e.g., 'next Monday'), calculate based on current date 2025-05-16."},
					"end_date":         {Type: genai.TypeString, Description: "End date (YYYY-MM-DD). If relative (e.g., 'in 2 weeks'), calculate based on current date 2025-05-16."},
					"budget":           {Type: genai.TypeNumber, Description: "Budget for the campaign."},
					"target_audience":  {Type: genai.TypeString, Description: "Target audience."},
					"channel":          {Type: genai.TypeString, Description: "Marketing channel."},
					"status":           {Type: genai.TypeString, Description: "Initial status (e.g., 'planned', 'active'). Default to 'planned' if not specified."},
					"goal":             {Type: genai.TypeString, Description: "The primary goal of the campaign."},
					"kpi_metric_name":  {Type: genai.TypeString, Description: "Name of the Key Performance Indicator metric.", Nullable: true},
					"kpi_target_value": {Type: genai.TypeNumber, Description: "Target value for the KPI.", Nullable: true},
					"notes":            {Type: genai.TypeString, Description: "Additional notes about the campaign.", Nullable: true},
				},
				Required: []string{"campaign_name", "start_date", "end_date", "budget", "target_audience", "channel", "goal"},
			},
		},
		{
			Name:        "updateMarketingCampaign",
			Description: "Updates an existing marketing campaign by its ID. Only include fields that need to be changed. The user 'RASTINEZ' is updating this. Current date is 2025-05-16.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"campaign_id":      {Type: genai.TypeInteger, Description: "The ID of the campaign to update."},
					"campaign_name":    {Type: genai.TypeString, Description: "New name for the campaign.", Nullable: true},
					"start_date":       {Type: genai.TypeString, Description: "New start date (YYYY-MM-DD).", Nullable: true},
					"end_date":         {Type: genai.TypeString, Description: "New end date (YYYY-MM-DD).", Nullable: true},
					"budget":           {Type: genai.TypeNumber, Description: "New budget for the campaign.", Nullable: true},
					"target_audience":  {Type: genai.TypeString, Description: "New target audience.", Nullable: true},
					"channel":          {Type: genai.TypeString, Description: "New marketing channel.", Nullable: true},
					"status":           {Type: genai.TypeString, Description: "New status (e.g., 'active', 'completed', 'paused').", Nullable: true},
					"goal":             {Type: genai.TypeString, Description: "New primary goal of the campaign.", Nullable: true},
					"kpi_metric_name":  {Type: genai.TypeString, Description: "New KPI metric name.", Nullable: true},
					"kpi_target_value": {Type: genai.TypeNumber, Description: "New target value for the KPI.", Nullable: true},
					"kpi_actual_value": {Type: genai.TypeNumber, Description: "Actual value achieved for the KPI (usually updated when campaign is completed or ongoing).", Nullable: true},
					"notes":            {Type: genai.TypeString, Description: "Updated notes about the campaign.", Nullable: true},
				},
				Required: []string{"campaign_id"},
			},
		},
		{
			Name:        "deleteMarketingCampaign",
			Description: "Deletes a marketing campaign by its ID. This action is irreversible. The user 'RASTINEZ' is deleting this.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"campaign_id": {
						Type:        genai.TypeInteger,
						Description: "The unique ID of the marketing campaign to delete.",
					},
				},
				Required: []string{"campaign_id"},
			},
		},
	},
}

// You can add BookingToolDeclarations here as well if you create a bookingTools.go
// var BookingToolDeclarations = &genai.Tool{ ... }


// GetAllTools combines all available tool declarations
// func GetAllTools() []*genai.Tool {
// 	var allTools []*genai.Tool
// 	if MarketingToolDeclarations != nil {
// 		allTools = append(allTools, MarketingToolDeclarations)
// 	}
// 	// Assuming BookingToolDeclarations is defined in the same package or imported
// 	// For example, if bookingTools.go also has `package tools`
// 	if BookingToolDeclarations != nil { // You would have defined BookingToolDeclarations
// 		allTools = append(allTools, BookingToolDeclarations)
// 	}
// 	return allTools
// }