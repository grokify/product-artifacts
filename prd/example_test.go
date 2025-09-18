package prd

import (
	"testing"
	"time"
)

func TestPRDMarshalUnmarshal(t *testing.T) {
	// Create a sample PRD
	now := time.Now()
	prd := &PRD{
		ID:          "PRD-TEST-001",
		Title:       "Test Product Feature",
		Version:     "1.0.0",
		CreatedDate: "2024-01-15",
		LastUpdated: &now,
		Owner: Owner{
			Name:  "Test Owner",
			Email: "test@example.com",
			Team:  "Test Team",
		},
		Status:   "draft",
		Priority: "high",
		Overview: Overview{
			ProblemStatement: "Test problem statement",
			SolutionSummary:  "Test solution summary",
			TargetAudience:   "Test users",
			MarketContext:    "Test market context",
		},
		Objectives: Objectives{
			BusinessGoals: []string{"Test goal 1", "Test goal 2"},
			SuccessMetrics: []SuccessMetric{
				{
					Metric:            "Test Metric",
					Target:            "100%",
					MeasurementMethod: "Test measurement",
				},
			},
		},
		Requirements: Requirements{
			Functional: []FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "Test functional requirement",
					Priority:    "must_have",
				},
			},
			NonFunctional: []NonFunctionalRequirement{
				{
					ID:          "NFR-001",
					Category:    "performance",
					Description: "Test non-functional requirement",
				},
			},
		},
	}

	// Test marshaling to JSON
	jsonStr, err := prd.ToJSON()
	if err != nil {
		t.Fatalf("Failed to marshal PRD to JSON: %v", err)
	}

	// Test unmarshaling from JSON
	unmarshaled, err := FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("Failed to unmarshal PRD from JSON: %v", err)
	}

	// Verify key fields
	if unmarshaled.ID != prd.ID {
		t.Errorf("Expected ID %s, got %s", prd.ID, unmarshaled.ID)
	}
	if unmarshaled.Title != prd.Title {
		t.Errorf("Expected Title %s, got %s", prd.Title, unmarshaled.Title)
	}
	if unmarshaled.Owner.Name != prd.Owner.Name {
		t.Errorf("Expected Owner Name %s, got %s", prd.Owner.Name, unmarshaled.Owner.Name)
	}
}

func TestPRDFileOperations(t *testing.T) {
	// Create a sample PRD
	prd := &PRD{
		ID:          "PRD-FILE-TEST-001",
		Title:       "File Test Product",
		Version:     "1.0.0",
		CreatedDate: "2024-01-15",
		Owner: Owner{
			Name:  "File Test Owner",
			Email: "filetest@example.com",
		},
		Status: "draft",
		Overview: Overview{
			ProblemStatement: "File test problem",
			SolutionSummary:  "File test solution",
		},
		Objectives: Objectives{
			BusinessGoals: []string{"File test goal"},
		},
		Requirements: Requirements{
			Functional: []FunctionalRequirement{
				{
					ID:          "FR-FILE-001",
					Description: "File test requirement",
					Priority:    "must_have",
				},
			},
		},
	}

	filename := "/tmp/test_prd.json"

	// Test saving to file
	err := prd.SaveToFile(filename)
	if err != nil {
		t.Fatalf("Failed to save PRD to file: %v", err)
	}

	// Test loading from file
	loaded, err := LoadFromFile(filename)
	if err != nil {
		t.Fatalf("Failed to load PRD from file: %v", err)
	}

	// Verify key fields
	if loaded.ID != prd.ID {
		t.Errorf("Expected ID %s, got %s", prd.ID, loaded.ID)
	}
	if loaded.Title != prd.Title {
		t.Errorf("Expected Title %s, got %s", prd.Title, loaded.Title)
	}
}

func TestPRDValidation(t *testing.T) {
	tests := []struct {
		name        string
		prd         *PRD
		expectError bool
	}{
		{
			name: "Valid PRD",
			prd: &PRD{
				ID:          "PRD-VALID-001",
				Title:       "Valid Product",
				Version:     "1.0.0",
				CreatedDate: "2024-01-15",
				Owner: Owner{
					Name:  "Valid Owner",
					Email: "valid@example.com",
				},
				Status: "draft",
				Overview: Overview{
					ProblemStatement: "Valid problem",
					SolutionSummary:  "Valid solution",
				},
				Objectives: Objectives{
					BusinessGoals: []string{"Valid goal"},
				},
				Requirements: Requirements{
					Functional: []FunctionalRequirement{
						{
							ID:          "FR-001",
							Description: "Valid requirement",
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "Missing ID",
			prd: &PRD{
				Title:   "Product without ID",
				Version: "1.0.0",
				Owner: Owner{
					Name:  "Owner",
					Email: "owner@example.com",
				},
				Status: "draft",
				Overview: Overview{
					ProblemStatement: "Problem",
					SolutionSummary:  "Solution",
				},
				Objectives: Objectives{
					BusinessGoals: []string{"Goal"},
				},
				Requirements: Requirements{
					Functional: []FunctionalRequirement{
						{
							ID:          "FR-001",
							Description: "Requirement",
						},
					},
				},
			},
			expectError: true,
		},
		{
			name: "Invalid Status",
			prd: &PRD{
				ID:      "PRD-INVALID-001",
				Title:   "Invalid Status Product",
				Version: "1.0.0",
				Owner: Owner{
					Name:  "Owner",
					Email: "owner@example.com",
				},
				Status: "invalid_status",
				Overview: Overview{
					ProblemStatement: "Problem",
					SolutionSummary:  "Solution",
				},
				Objectives: Objectives{
					BusinessGoals: []string{"Goal"},
				},
				Requirements: Requirements{
					Functional: []FunctionalRequirement{
						{
							ID:          "FR-001",
							Description: "Requirement",
						},
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.prd.Validate()
			if tt.expectError && err == nil {
				t.Error("Expected validation error, but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no validation error, but got: %v", err)
			}
		})
	}
}

func TestUpdateLastModified(t *testing.T) {
	prd := &PRD{
		ID:          "PRD-UPDATE-001",
		Title:       "Update Test Product",
		Version:     "1.0.0",
		CreatedDate: "2024-01-15",
		Owner: Owner{
			Name:  "Test Owner",
			Email: "test@example.com",
		},
		Status: "draft",
		Overview: Overview{
			ProblemStatement: "Test problem",
			SolutionSummary:  "Test solution",
		},
		Objectives: Objectives{
			BusinessGoals: []string{"Test goal"},
		},
		Requirements: Requirements{
			Functional: []FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "Test requirement",
				},
			},
		},
	}

	// Initially, LastUpdated should be nil
	if prd.LastUpdated != nil {
		t.Error("Expected LastUpdated to be nil initially")
	}

	// Update last modified
	prd.UpdateLastModified()

	// Now LastUpdated should be set
	if prd.LastUpdated == nil {
		t.Error("Expected LastUpdated to be set after UpdateLastModified()")
	}

	// Check that the timestamp is recent (within last minute)
	now := time.Now()
	if now.Sub(*prd.LastUpdated) > time.Minute {
		t.Error("LastUpdated timestamp seems too old")
	}
}
