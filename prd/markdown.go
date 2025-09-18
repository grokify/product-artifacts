package prd

import (
	"fmt"
	"strings"
)

// ToMarkdown converts the PRD struct into a formatted Markdown document
func (p *PRD) ToMarkdown() string {
	var md strings.Builder

	// Header
	md.WriteString(fmt.Sprintf("# %s\n\n", p.Title))

	// Metadata table
	md.WriteString("| Field | Value |\n")
	md.WriteString("|-------|-------|\n")
	md.WriteString(fmt.Sprintf("| **ID** | %s |\n", p.ID))
	md.WriteString(fmt.Sprintf("| **Version** | %s |\n", p.Version))
	md.WriteString(fmt.Sprintf("| **Created Date** | %s |\n", p.CreatedDate))
	if p.LastUpdated != nil {
		md.WriteString(fmt.Sprintf("| **Last Updated** | %s |\n", p.LastUpdated.Format("2006-01-02 15:04:05")))
	}
	md.WriteString(fmt.Sprintf("| **Status** | %s |\n", p.Status))
	if p.Priority != "" {
		md.WriteString(fmt.Sprintf("| **Priority** | %s |\n", p.Priority))
	}
	md.WriteString("\n")

	// Owner and Stakeholders
	md.WriteString("## Owner\n\n")
	md.WriteString(fmt.Sprintf("**Name:** %s  \n", p.Owner.Name))
	md.WriteString(fmt.Sprintf("**Email:** %s  \n", p.Owner.Email))
	if p.Owner.Team != "" {
		md.WriteString(fmt.Sprintf("**Team:** %s  \n", p.Owner.Team))
	}
	md.WriteString("\n")

	if len(p.Stakeholders) > 0 {
		md.WriteString("## Stakeholders\n\n")
		for _, stakeholder := range p.Stakeholders {
			md.WriteString(fmt.Sprintf("- **%s** (%s)", stakeholder.Name, stakeholder.Role))
			if stakeholder.Email != "" {
				md.WriteString(fmt.Sprintf(" - %s", stakeholder.Email))
			}
			if stakeholder.Team != "" {
				md.WriteString(fmt.Sprintf(" - Team: %s", stakeholder.Team))
			}
			md.WriteString("\n")
		}
		md.WriteString("\n")
	}

	// Overview
	md.WriteString("## Overview\n\n")
	md.WriteString("### Problem Statement\n\n")
	md.WriteString(fmt.Sprintf("%s\n\n", p.Overview.ProblemStatement))

	md.WriteString("### Solution Summary\n\n")
	md.WriteString(fmt.Sprintf("%s\n\n", p.Overview.SolutionSummary))

	if p.Overview.TargetAudience != "" {
		md.WriteString("### Target Audience\n\n")
		md.WriteString(fmt.Sprintf("%s\n\n", p.Overview.TargetAudience))
	}

	if p.Overview.MarketContext != "" {
		md.WriteString("### Market Context\n\n")
		md.WriteString(fmt.Sprintf("%s\n\n", p.Overview.MarketContext))
	}

	// Objectives
	md.WriteString("## Objectives\n\n")
	md.WriteString("### Business Goals\n\n")
	for i, goal := range p.Objectives.BusinessGoals {
		md.WriteString(fmt.Sprintf("%d. %s\n", i+1, goal))
	}
	md.WriteString("\n")

	if len(p.Objectives.SuccessMetrics) > 0 {
		md.WriteString("### Success Metrics\n\n")
		md.WriteString("| Metric | Target | Measurement Method |\n")
		md.WriteString("|--------|--------|--------------------|")
		md.WriteString("\n")
		for _, metric := range p.Objectives.SuccessMetrics {
			md.WriteString(fmt.Sprintf("| %s | %s | %s |\n", metric.Metric, metric.Target, metric.MeasurementMethod))
		}
		md.WriteString("\n")
	}

	if len(p.Objectives.OKRs) > 0 {
		md.WriteString("### OKRs\n\n")
		for _, okr := range p.Objectives.OKRs {
			md.WriteString(fmt.Sprintf("**Objective:** %s\n\n", okr.Objective))
			md.WriteString("**Key Results:**\n")
			for _, kr := range okr.KeyResults {
				md.WriteString(fmt.Sprintf("- %s\n", kr))
			}
			md.WriteString("\n")
		}
	}

	// User Personas
	if len(p.UserPersonas) > 0 {
		md.WriteString("## User Personas\n\n")
		for _, persona := range p.UserPersonas {
			md.WriteString(fmt.Sprintf("### %s\n\n", persona.Name))
			md.WriteString(fmt.Sprintf("%s\n\n", persona.Description))

			if len(persona.Goals) > 0 {
				md.WriteString("**Goals:**\n")
				for _, goal := range persona.Goals {
					md.WriteString(fmt.Sprintf("- %s\n", goal))
				}
				md.WriteString("\n")
			}

			if len(persona.PainPoints) > 0 {
				md.WriteString("**Pain Points:**\n")
				for _, pain := range persona.PainPoints {
					md.WriteString(fmt.Sprintf("- %s\n", pain))
				}
				md.WriteString("\n")
			}
		}
	}

	// User Stories
	if len(p.UserStories) > 0 {
		md.WriteString("## User Stories\n\n")
		for _, story := range p.UserStories {
			md.WriteString(fmt.Sprintf("### %s\n\n", story.ID))
			md.WriteString(fmt.Sprintf("**Story:** %s\n\n", story.Story))

			if len(story.AcceptanceCriteria) > 0 {
				md.WriteString("**Acceptance Criteria:**\n")
				for _, criteria := range story.AcceptanceCriteria {
					md.WriteString(fmt.Sprintf("- %s\n", criteria))
				}
				md.WriteString("\n")
			}

			if story.Priority != "" {
				md.WriteString(fmt.Sprintf("**Priority:** %s\n", story.Priority))
			}
			if story.EffortEstimate != "" {
				md.WriteString(fmt.Sprintf("**Effort Estimate:** %s\n", story.EffortEstimate))
			}
			md.WriteString("\n")
		}
	}

	// Requirements
	md.WriteString("## Requirements\n\n")
	md.WriteString("### Functional Requirements\n\n")
	for _, req := range p.Requirements.Functional {
		md.WriteString(fmt.Sprintf("#### %s\n\n", req.ID))
		md.WriteString(fmt.Sprintf("%s\n\n", req.Description))

		if req.Priority != "" {
			md.WriteString(fmt.Sprintf("**Priority:** %s\n", req.Priority))
		}

		if len(req.Dependencies) > 0 {
			md.WriteString("**Dependencies:**\n")
			for _, dep := range req.Dependencies {
				md.WriteString(fmt.Sprintf("- %s\n", dep))
			}
		}
		md.WriteString("\n")
	}

	if len(p.Requirements.NonFunctional) > 0 {
		md.WriteString("### Non-Functional Requirements\n\n")
		for _, req := range p.Requirements.NonFunctional {
			md.WriteString(fmt.Sprintf("#### %s (%s)\n\n", req.ID, req.Category))
			md.WriteString(fmt.Sprintf("%s\n\n", req.Description))

			if req.AcceptanceCriteria != "" {
				md.WriteString(fmt.Sprintf("**Acceptance Criteria:** %s\n\n", req.AcceptanceCriteria))
			}
		}
	}

	// Technical Specifications
	if p.TechnicalSpecifications != nil {
		md.WriteString("## Technical Specifications\n\n")

		if p.TechnicalSpecifications.ArchitectureOverview != "" {
			md.WriteString("### Architecture Overview\n\n")
			md.WriteString(fmt.Sprintf("%s\n\n", p.TechnicalSpecifications.ArchitectureOverview))
		}

		if p.TechnicalSpecifications.TechnologyStack != nil {
			md.WriteString("### Technology Stack\n\n")
			stack := p.TechnicalSpecifications.TechnologyStack

			if len(stack.Frontend) > 0 {
				md.WriteString("**Frontend:**\n")
				for _, tech := range stack.Frontend {
					md.WriteString(fmt.Sprintf("- %s\n", tech))
				}
				md.WriteString("\n")
			}

			if len(stack.Backend) > 0 {
				md.WriteString("**Backend:**\n")
				for _, tech := range stack.Backend {
					md.WriteString(fmt.Sprintf("- %s\n", tech))
				}
				md.WriteString("\n")
			}

			if len(stack.Database) > 0 {
				md.WriteString("**Database:**\n")
				for _, tech := range stack.Database {
					md.WriteString(fmt.Sprintf("- %s\n", tech))
				}
				md.WriteString("\n")
			}

			if len(stack.Infrastructure) > 0 {
				md.WriteString("**Infrastructure:**\n")
				for _, tech := range stack.Infrastructure {
					md.WriteString(fmt.Sprintf("- %s\n", tech))
				}
				md.WriteString("\n")
			}
		}

		if len(p.TechnicalSpecifications.APISpecifications) > 0 {
			md.WriteString("### API Specifications\n\n")
			for _, api := range p.TechnicalSpecifications.APISpecifications {
				if api.Endpoint != "" && api.Method != "" {
					md.WriteString(fmt.Sprintf("#### %s %s\n\n", api.Method, api.Endpoint))
				}
				if api.Description != "" {
					md.WriteString(fmt.Sprintf("%s\n\n", api.Description))
				}
				if api.RequestFormat != "" {
					md.WriteString(fmt.Sprintf("**Request Format:** %s\n", api.RequestFormat))
				}
				if api.ResponseFormat != "" {
					md.WriteString(fmt.Sprintf("**Response Format:** %s\n", api.ResponseFormat))
				}
				md.WriteString("\n")
			}
		}

		if len(p.TechnicalSpecifications.SecurityConsiderations) > 0 {
			md.WriteString("### Security Considerations\n\n")
			for _, consideration := range p.TechnicalSpecifications.SecurityConsiderations {
				md.WriteString(fmt.Sprintf("- %s\n", consideration))
			}
			md.WriteString("\n")
		}
	}

	// Timeline
	if p.Timeline != nil {
		md.WriteString("## Timeline\n\n")

		if len(p.Timeline.Milestones) > 0 {
			md.WriteString("### Milestones\n\n")
			for _, milestone := range p.Timeline.Milestones {
				md.WriteString(fmt.Sprintf("#### %s\n\n", milestone.Name))
				if milestone.Description != "" {
					md.WriteString(fmt.Sprintf("%s\n\n", milestone.Description))
				}
				md.WriteString(fmt.Sprintf("**Target Date:** %s\n", milestone.TargetDate))

				if len(milestone.Dependencies) > 0 {
					md.WriteString("**Dependencies:**\n")
					for _, dep := range milestone.Dependencies {
						md.WriteString(fmt.Sprintf("- %s\n", dep))
					}
				}
				md.WriteString("\n")
			}
		}

		if p.Timeline.LaunchDate != "" {
			md.WriteString(fmt.Sprintf("### Launch Date: %s\n\n", p.Timeline.LaunchDate))
		}
	}

	// Risks and Assumptions
	if p.RisksAndAssumptions != nil {
		md.WriteString("## Risks and Assumptions\n\n")

		if len(p.RisksAndAssumptions.Risks) > 0 {
			md.WriteString("### Risks\n\n")
			md.WriteString("| Risk | Impact | Probability | Mitigation Strategy |\n")
			md.WriteString("|------|--------|-------------|--------------------|\n")
			for _, risk := range p.RisksAndAssumptions.Risks {
				md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					risk.Description, risk.Impact, risk.Probability, risk.MitigationStrategy))
			}
			md.WriteString("\n")
		}

		if len(p.RisksAndAssumptions.Assumptions) > 0 {
			md.WriteString("### Assumptions\n\n")
			for _, assumption := range p.RisksAndAssumptions.Assumptions {
				md.WriteString(fmt.Sprintf("- %s\n", assumption))
			}
			md.WriteString("\n")
		}
	}

	// Out of Scope
	if len(p.OutOfScope) > 0 {
		md.WriteString("## Out of Scope\n\n")
		for _, item := range p.OutOfScope {
			md.WriteString(fmt.Sprintf("- %s\n", item))
		}
		md.WriteString("\n")
	}

	// Appendices
	if p.Appendices != nil {
		md.WriteString("## Appendices\n\n")

		if p.Appendices.ResearchData != "" {
			md.WriteString("### Research Data\n\n")
			md.WriteString(fmt.Sprintf("%s\n\n", p.Appendices.ResearchData))
		}

		if len(p.Appendices.MockupsWireframes) > 0 {
			md.WriteString("### Mockups and Wireframes\n\n")
			for _, mockup := range p.Appendices.MockupsWireframes {
				if mockup.Name != "" {
					md.WriteString(fmt.Sprintf("#### %s\n\n", mockup.Name))
				}
				if mockup.Description != "" {
					md.WriteString(fmt.Sprintf("%s\n\n", mockup.Description))
				}
				if mockup.URL != "" {
					md.WriteString(fmt.Sprintf("[View Mockup](%s)\n\n", mockup.URL))
				}
			}
		}

		if len(p.Appendices.RelatedDocuments) > 0 {
			md.WriteString("### Related Documents\n\n")
			for _, doc := range p.Appendices.RelatedDocuments {
				if doc.Title != "" && doc.URL != "" {
					md.WriteString(fmt.Sprintf("- [%s](%s)", doc.Title, doc.URL))
					if doc.Type != "" {
						md.WriteString(fmt.Sprintf(" (%s)", doc.Type))
					}
					md.WriteString("\n")
				}
			}
			md.WriteString("\n")
		}
	}

	return md.String()
}
