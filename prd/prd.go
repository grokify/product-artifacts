package prd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// PRD represents a Product Requirements Document
type PRD struct {
	ID                      string                   `json:"id"`
	Title                   string                   `json:"title"`
	Version                 string                   `json:"version"`
	CreatedDate             string                   `json:"created_date"`
	LastUpdated             *time.Time               `json:"last_updated,omitempty"`
	Owner                   Owner                    `json:"owner"`
	Stakeholders            []Stakeholder            `json:"stakeholders,omitempty"`
	Status                  string                   `json:"status"`
	Priority                string                   `json:"priority,omitempty"`
	Overview                Overview                 `json:"overview"`
	Objectives              Objectives               `json:"objectives"`
	UserPersonas            []UserPersona            `json:"user_personas,omitempty"`
	UserStories             []UserStory              `json:"user_stories,omitempty"`
	Requirements            Requirements             `json:"requirements"`
	TechnicalSpecifications *TechnicalSpecifications `json:"technical_specifications,omitempty"`
	Timeline                *Timeline                `json:"timeline,omitempty"`
	RisksAndAssumptions     *RisksAndAssumptions     `json:"risks_and_assumptions,omitempty"`
	OutOfScope              []string                 `json:"out_of_scope,omitempty"`
	Appendices              *Appendices              `json:"appendices,omitempty"`
}

// Owner represents the product owner
type Owner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Team  string `json:"team,omitempty"`
}

// Stakeholder represents a project stakeholder
type Stakeholder struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role"`
	Team  string `json:"team,omitempty"`
}

// Overview contains the high-level product overview
type Overview struct {
	ProblemStatement string `json:"problem_statement"`
	SolutionSummary  string `json:"solution_summary"`
	TargetAudience   string `json:"target_audience,omitempty"`
	MarketContext    string `json:"market_context,omitempty"`
}

// Objectives contains business goals and success metrics
type Objectives struct {
	BusinessGoals  []string        `json:"business_goals"`
	SuccessMetrics []SuccessMetric `json:"success_metrics,omitempty"`
	OKRs           []OKR           `json:"okrs,omitempty"`
}

// SuccessMetric represents a measurable success indicator
type SuccessMetric struct {
	Metric            string `json:"metric"`
	Target            string `json:"target"`
	MeasurementMethod string `json:"measurement_method,omitempty"`
}

// OKR represents an Objective and Key Results
type OKR struct {
	Objective  string   `json:"objective"`
	KeyResults []string `json:"key_results"`
}

// UserPersona represents a target user persona
type UserPersona struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Goals       []string `json:"goals,omitempty"`
	PainPoints  []string `json:"pain_points,omitempty"`
}

// UserStory represents a user story
type UserStory struct {
	ID                 string   `json:"id"`
	Story              string   `json:"story"`
	AcceptanceCriteria []string `json:"acceptance_criteria"`
	Priority           string   `json:"priority,omitempty"`
	EffortEstimate     string   `json:"effort_estimate,omitempty"`
}

// Requirements contains functional and non-functional requirements
type Requirements struct {
	Functional    []FunctionalRequirement    `json:"functional"`
	NonFunctional []NonFunctionalRequirement `json:"non_functional,omitempty"`
}

// FunctionalRequirement represents a functional requirement
type FunctionalRequirement struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Priority     string   `json:"priority,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
}

// NonFunctionalRequirement represents a non-functional requirement
type NonFunctionalRequirement struct {
	ID                 string `json:"id"`
	Category           string `json:"category"`
	Description        string `json:"description"`
	AcceptanceCriteria string `json:"acceptance_criteria,omitempty"`
}

// TechnicalSpecifications contains technical details
type TechnicalSpecifications struct {
	ArchitectureOverview   string             `json:"architecture_overview,omitempty"`
	TechnologyStack        *TechnologyStack   `json:"technology_stack,omitempty"`
	APISpecifications      []APISpecification `json:"api_specifications,omitempty"`
	SecurityConsiderations []string           `json:"security_considerations,omitempty"`
}

// TechnologyStack represents the technology stack
type TechnologyStack struct {
	Frontend       []string `json:"frontend,omitempty"`
	Backend        []string `json:"backend,omitempty"`
	Database       []string `json:"database,omitempty"`
	Infrastructure []string `json:"infrastructure,omitempty"`
}

// APISpecification represents an API endpoint specification
type APISpecification struct {
	Endpoint       string `json:"endpoint,omitempty"`
	Method         string `json:"method,omitempty"`
	Description    string `json:"description,omitempty"`
	RequestFormat  string `json:"request_format,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
}

// Timeline contains project timeline information
type Timeline struct {
	Milestones []Milestone `json:"milestones,omitempty"`
	LaunchDate string      `json:"launch_date,omitempty"`
}

// Milestone represents a project milestone
type Milestone struct {
	Name         string   `json:"name"`
	Description  string   `json:"description,omitempty"`
	TargetDate   string   `json:"target_date"`
	Dependencies []string `json:"dependencies,omitempty"`
}

// RisksAndAssumptions contains risks and assumptions
type RisksAndAssumptions struct {
	Risks       []Risk   `json:"risks,omitempty"`
	Assumptions []string `json:"assumptions,omitempty"`
}

// Risk represents a project risk
type Risk struct {
	Description        string `json:"description"`
	Impact             string `json:"impact"`
	Probability        string `json:"probability"`
	MitigationStrategy string `json:"mitigation_strategy,omitempty"`
}

// Appendices contains supporting documents and references
type Appendices struct {
	ResearchData      string            `json:"research_data,omitempty"`
	MockupsWireframes []MockupWireframe `json:"mockups_wireframes,omitempty"`
	RelatedDocuments  []RelatedDocument `json:"related_documents,omitempty"`
}

// MockupWireframe represents a mockup or wireframe reference
type MockupWireframe struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

// RelatedDocument represents a related document reference
type RelatedDocument struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
	Type  string `json:"type,omitempty"`
}

// LoadFromFile loads a PRD from a JSON file
func LoadFromFile(filename string) (*PRD, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	var prd PRD
	if err := json.Unmarshal(data, &prd); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &prd, nil
}

// SaveToFile saves a PRD to a JSON file
func (p *PRD) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal PRD to JSON: %w", err)
	}

	if err := os.WriteFile(filename, data, 0600); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// ToJSON converts the PRD to a JSON string
func (p *PRD) ToJSON() (string, error) {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal PRD to JSON: %w", err)
	}
	return string(data), nil
}

// FromJSON creates a PRD from a JSON string
func FromJSON(jsonStr string) (*PRD, error) {
	var prd PRD
	if err := json.Unmarshal([]byte(jsonStr), &prd); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return &prd, nil
}

// UpdateLastModified updates the last_updated timestamp to now
func (p *PRD) UpdateLastModified() {
	now := time.Now()
	p.LastUpdated = &now
}

// Validate performs basic validation on the PRD
func (p *PRD) Validate() error {
	if p.ID == "" {
		return fmt.Errorf("PRD ID is required")
	}
	if p.Title == "" {
		return fmt.Errorf("PRD title is required")
	}
	if p.Version == "" {
		return fmt.Errorf("PRD version is required")
	}
	if p.Owner.Name == "" {
		return fmt.Errorf("owner name is required")
	}
	if p.Owner.Email == "" {
		return fmt.Errorf("owner email is required")
	}
	if p.Overview.ProblemStatement == "" {
		return fmt.Errorf("problem statement is required")
	}
	if p.Overview.SolutionSummary == "" {
		return fmt.Errorf("solution summary is required")
	}
	if len(p.Objectives.BusinessGoals) == 0 {
		return fmt.Errorf("at least one business goal is required")
	}
	if len(p.Requirements.Functional) == 0 {
		return fmt.Errorf("at least one functional requirement is required")
	}

	// Validate status enum
	validStatuses := map[string]bool{
		"draft":          true,
		"review":         true,
		"approved":       true,
		"in_development": true,
		"completed":      true,
		"archived":       true,
	}
	if !validStatuses[p.Status] {
		return fmt.Errorf("invalid status: %s", p.Status)
	}

	return nil
}
