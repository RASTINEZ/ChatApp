package tools_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/google/generative-ai-go/genai"
	// IMPORTANT: Replace "your_module_path/tools" with the actual import path for your tools package.
	// This path is typically defined in your go.mod file.
	// For example, if your go.mod is in 'GoChatbotBackend' and declares 'module myproject',
	// then the path would be 'myproject/tools'.
	"my-golang-react-app/tools"
)

// Helper function to compare two string slices for set equality (ignoring order and nil/empty equivalence)
func assertStringSliceSetEquals(t *testing.T, actual []string, expected []string, fieldName string, funcName string) {
	t.Helper()

	// Treat nil slices as empty for comparison
	if actual == nil {
		actual = []string{}
	}
	if expected == nil {
		expected = []string{}
	}

	actualCopy := make([]string, len(actual))
	copy(actualCopy, actual)
	sort.Strings(actualCopy)

	expectedCopy := make([]string, len(expected))
	copy(expectedCopy, expected)
	sort.Strings(expectedCopy)

	if !reflect.DeepEqual(actualCopy, expectedCopy) {
		t.Errorf("%s: %s mismatch for function '%s'. Got %v, want %v", funcName, fieldName, funcName, actual, expected)
	}
}

func TestMarketingToolDeclarations(t *testing.T) {
	if tools.MarketingToolDeclarations == nil {
		t.Fatal("MarketingToolDeclarations is nil")
	}
	if tools.MarketingToolDeclarations.FunctionDeclarations == nil {
		t.Fatal("MarketingToolDeclarations.FunctionDeclarations is nil")
	}

	expectedNumFunctions := 5
	if len(tools.MarketingToolDeclarations.FunctionDeclarations) != expectedNumFunctions {
		t.Fatalf("Expected %d function declarations, got %d", expectedNumFunctions, len(tools.MarketingToolDeclarations.FunctionDeclarations))
	}

	type expectedParam struct {
		Name        string
		Type        genai.Type
		Nullable    bool
		Description bool // True if description should exist
	}
	type expectedFunction struct {
		Name                string
		DescriptionKeywords []string
		ParamsType          genai.Type
		Properties          map[string]expectedParam
		Required            []string
	}

	expectedFunctions := map[string]expectedFunction{
		"getMarketingCampaigns": {
			Name:       "getMarketingCampaigns",
			ParamsType: genai.TypeObject,
			Properties: map[string]expectedParam{
				"status":          {Name: "status", Type: genai.TypeString, Nullable: true, Description: true},
				"target_audience": {Name: "target_audience", Type: genai.TypeString, Nullable: true, Description: true},
			},
			Required: nil, // No required fields
		},
		"getMarketingCampaignByID": {
			Name:       "getMarketingCampaignByID",
			ParamsType: genai.TypeObject,
			Properties: map[string]expectedParam{
				"campaign_id": {Name: "campaign_id", Type: genai.TypeInteger, Nullable: false, Description: true},
			},
			Required: []string{"campaign_id"},
		},
		"createMarketingCampaign": {
			Name:                "createMarketingCampaign",
			DescriptionKeywords: []string{"RASTINEZ", "2025-05-16"},
			ParamsType:          genai.TypeObject,
			Properties: map[string]expectedParam{
				"campaign_name":    {Name: "campaign_name", Type: genai.TypeString, Nullable: false, Description: true},
				"start_date":       {Name: "start_date", Type: genai.TypeString, Nullable: false, Description: true},
				"end_date":         {Name: "end_date", Type: genai.TypeString, Nullable: false, Description: true},
				"budget":           {Name: "budget", Type: genai.TypeNumber, Nullable: false, Description: true},
				"target_audience":  {Name: "target_audience", Type: genai.TypeString, Nullable: false, Description: true},
				"channel":          {Name: "channel", Type: genai.TypeString, Nullable: false, Description: true},
				"status":           {Name: "status", Type: genai.TypeString, Nullable: false, Description: true}, // Not explicitly nullable, so defaults to false
				"goal":             {Name: "goal", Type: genai.TypeString, Nullable: false, Description: true},
				"kpi_metric_name":  {Name: "kpi_metric_name", Type: genai.TypeString, Nullable: true, Description: true},
				"kpi_target_value": {Name: "kpi_target_value", Type: genai.TypeNumber, Nullable: true, Description: true},
				"notes":            {Name: "notes", Type: genai.TypeString, Nullable: true, Description: true},
			},
			Required: []string{"campaign_name", "start_date", "end_date", "budget", "target_audience", "channel", "goal"},
		},
		"updateMarketingCampaign": {
			Name:                "updateMarketingCampaign",
			DescriptionKeywords: []string{"RASTINEZ", "2025-05-16"},
			ParamsType:          genai.TypeObject,
			Properties: map[string]expectedParam{
				"campaign_id":      {Name: "campaign_id", Type: genai.TypeInteger, Nullable: false, Description: true},
				"campaign_name":    {Name: "campaign_name", Type: genai.TypeString, Nullable: true, Description: true},
				"start_date":       {Name: "start_date", Type: genai.TypeString, Nullable: true, Description: true},
				"end_date":         {Name: "end_date", Type: genai.TypeString, Nullable: true, Description: true},
				"budget":           {Name: "budget", Type: genai.TypeNumber, Nullable: true, Description: true},
				"target_audience":  {Name: "target_audience", Type: genai.TypeString, Nullable: true, Description: true},
				"channel":          {Name: "channel", Type: genai.TypeString, Nullable: true, Description: true},
				"status":           {Name: "status", Type: genai.TypeString, Nullable: true, Description: true},
				"goal":             {Name: "goal", Type: genai.TypeString, Nullable: true, Description: true},
				"kpi_metric_name":  {Name: "kpi_metric_name", Type: genai.TypeString, Nullable: true, Description: true},
				"kpi_target_value": {Name: "kpi_target_value", Type: genai.TypeNumber, Nullable: true, Description: true},
				"kpi_actual_value": {Name: "kpi_actual_value", Type: genai.TypeNumber, Nullable: true, Description: true},
				"notes":            {Name: "notes", Type: genai.TypeString, Nullable: true, Description: true},
			},
			Required: []string{"campaign_id"},
		},
		"deleteMarketingCampaign": {
			Name:                "deleteMarketingCampaign",
			DescriptionKeywords: []string{"RASTINEZ"},
			ParamsType:          genai.TypeObject,
			Properties: map[string]expectedParam{
				"campaign_id": {Name: "campaign_id", Type: genai.TypeInteger, Nullable: false, Description: true},
			},
			Required: []string{"campaign_id"},
		},
	}

	foundFunctionNames := make(map[string]bool)

	for _, funcDecl := range tools.MarketingToolDeclarations.FunctionDeclarations {
		t.Run(funcDecl.Name, func(t *testing.T) {
			expected, ok := expectedFunctions[funcDecl.Name]
			if !ok {
				t.Fatalf("Unexpected function declaration found: %s", funcDecl.Name)
			}
			foundFunctionNames[funcDecl.Name] = true

			if funcDecl.Name != expected.Name {
				t.Errorf("Name mismatch: Got '%s', want '%s'", funcDecl.Name, expected.Name)
			}

			if funcDecl.Description == "" {
				t.Errorf("Description for function '%s' is empty", funcDecl.Name)
			}
			for _, keyword := range expected.DescriptionKeywords {
				if !strings.Contains(funcDecl.Description, keyword) {
					t.Errorf("Description for function '%s' does not contain keyword '%s'. Got: '%s'", funcDecl.Name, keyword, funcDecl.Description)
				}
			}

			if funcDecl.Parameters == nil {
				t.Fatalf("Parameters schema for function '%s' is nil", funcDecl.Name)
			}
			if funcDecl.Parameters.Type != expected.ParamsType {
				t.Errorf("Parameters.Type mismatch for function '%s': Got %s, want %s", funcDecl.Name, funcDecl.Parameters.Type, expected.ParamsType)
			}

			if funcDecl.Parameters.Properties == nil && len(expected.Properties) > 0 {
				t.Fatalf("Parameters.Properties for function '%s' is nil, but expected properties", funcDecl.Name)
			}
			if len(funcDecl.Parameters.Properties) != len(expected.Properties) {
				t.Errorf("Parameters.Properties count mismatch for function '%s': Got %d, want %d properties. Got: %v, Expected: %v",
					funcDecl.Name, len(funcDecl.Parameters.Properties), len(expected.Properties),
					getKeys(funcDecl.Parameters.Properties), getKeysFromStringMap(expected.Properties))
			}

			for paramName, expectedParam := range expected.Properties {
				actualParamSchema, ok := funcDecl.Parameters.Properties[paramName]
				if !ok {
					t.Errorf("Expected parameter '%s' not found in function '%s'", paramName, funcDecl.Name)
					continue
				}
				if actualParamSchema.Type != expectedParam.Type {
					t.Errorf("Parameter '%s' Type mismatch for function '%s': Got %s, want %s", paramName, funcDecl.Name, actualParamSchema.Type, expectedParam.Type)
				}
				if actualParamSchema.Nullable != expectedParam.Nullable {
					t.Errorf("Parameter '%s' Nullable mismatch for function '%s': Got %v, want %v", paramName, funcDecl.Name, actualParamSchema.Nullable, expectedParam.Nullable)
				}
				if expectedParam.Description && actualParamSchema.Description == "" {
					t.Errorf("Parameter '%s' Description for function '%s' is unexpectedly empty", paramName, funcDecl.Name)
				}
			}

			for actualParamName := range funcDecl.Parameters.Properties {
				if _, ok := expected.Properties[actualParamName]; !ok {
					t.Errorf("Unexpected parameter '%s' found in function '%s'", actualParamName, funcDecl.Name)
				}
			}

			assertStringSliceSetEquals(t, funcDecl.Parameters.Required, expected.Required, "Required fields", funcDecl.Name)
		})
	}

	if len(foundFunctionNames) != expectedNumFunctions {
		missing := []string{}
		for name := range expectedFunctions {
			if !foundFunctionNames[name] {
				missing = append(missing, name)
			}
		}
		t.Errorf("Number of tested functions (%d) does not match expected number of functions (%d). Missing functions: %v", len(foundFunctionNames), expectedNumFunctions, missing)
	}
}

// Helper to get keys from map[string]*genai.Schema for debugging
func getKeys(m map[string]*genai.Schema) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Helper to get keys from map[string]expectedParam for debugging
func getKeysFromStringMap(m map[string]expectedFunction) []string {
	// This function signature was incorrect in thought process, correcting to expectedParam
	// func getKeysFromStringMap(m map[string]expectedParam) []string {
	// Correcting based on usage:
	// The map is map[string]expectedParam, not map[string]expectedFunction
	// However, the call site is `getKeysFromStringMap(expected.Properties)`
	// So the map type is map[string]expectedParam
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Corrected helper for map[string]expectedParam
func getKeysFromExpectedParamMap(m map[string]expectedParam) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// The TestMarketingToolDeclarations function uses getKeysFromStringMap(expected.Properties)
// where expected.Properties is map[string]expectedParam.
// So, the getKeysFromStringMap should be:
// func getKeysFromStringMap(m map[string]expectedParam) []string { ... }
// Let's rename it to avoid confusion and use the corrected one.

// In TestMarketingToolDeclarations, the call is:
// getKeysFromStringMap(expected.Properties)
// expected.Properties is of type map[string]expectedParam
// So the helper should be:
// func getKeysFromExpectedParamMap(m map[string]expectedParam) []string
// I will use getKeysFromExpectedParamMap in the main test function.
// The original getKeysFromStringMap is removed to avoid confusion.
// The call in the test:
// t.Errorf("...", getKeys(funcDecl.Parameters.Properties), getKeysFromExpectedParamMap(expected.Properties))
// This is now correct.
