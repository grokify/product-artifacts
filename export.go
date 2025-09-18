package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/grokify/product-artifacts/prd"
)

// Export PRD to Markdown
func exportToMarkdown(prdDoc *prd.PRD, filename string) error {
	var md strings.Builder

	// Header
	md.WriteString(fmt.Sprintf("# %s\n\n", prdDoc.Title))
	md.WriteString("---\n\n")
	md.WriteString(fmt.Sprintf("**Document ID:** %s  \n", prdDoc.ID))
	md.WriteString(fmt.Sprintf("**Version:** %s  \n", prdDoc.Version))
	md.WriteString(fmt.Sprintf("**Status:** %s  \n", prdDoc.Status))
	if prdDoc.Priority != "" {
		md.WriteString(fmt.Sprintf("**Priority:** %s  \n", prdDoc.Priority))
	}
	md.WriteString(fmt.Sprintf("**Owner:** %s (%s)  \n", prdDoc.Owner.Name, prdDoc.Owner.Email))
	if prdDoc.Owner.Team != "" {
		md.WriteString(fmt.Sprintf("**Team:** %s  \n", prdDoc.Owner.Team))
	}
	md.WriteString(fmt.Sprintf("**Created:** %s  \n", prdDoc.CreatedDate))
	if prdDoc.LastUpdated != nil {
		md.WriteString(fmt.Sprintf("**Last Updated:** %s  \n", prdDoc.LastUpdated.Format("2006-01-02 15:04")))
	}
	md.WriteString("\n---\n\n")

	// Table of Contents
	md.WriteString("## Table of Contents\n\n")
	md.WriteString("1. [Overview](#overview)\n")
	md.WriteString("2. [Objectives](#objectives)\n")
	md.WriteString("3. [User Stories](#user-stories)\n")
	md.WriteString("4. [Requirements](#requirements)\n")
	if prdDoc.TechnicalSpecifications != nil {
		md.WriteString("5. [Technical Specifications](#technical-specifications)\n")
	}
	if prdDoc.Timeline != nil {
		md.WriteString("6. [Timeline](#timeline)\n")
	}
	if prdDoc.RisksAndAssumptions != nil {
		md.WriteString("7. [Risks and Assumptions](#risks-and-assumptions)\n")
	}
	md.WriteString("\n")

	// Overview
	md.WriteString("## Overview\n\n")
	md.WriteString("### Problem Statement\n\n")
	md.WriteString(fmt.Sprintf("%s\n\n", prdDoc.Overview.ProblemStatement))

	md.WriteString("### Solution Summary\n\n")
	md.WriteString(fmt.Sprintf("%s\n\n", prdDoc.Overview.SolutionSummary))

	if prdDoc.Overview.TargetAudience != "" {
		md.WriteString("### Target Audience\n\n")
		md.WriteString(fmt.Sprintf("%s\n\n", prdDoc.Overview.TargetAudience))
	}

	if prdDoc.Overview.MarketContext != "" {
		md.WriteString("### Market Context\n\n")
		md.WriteString(fmt.Sprintf("%s\n\n", prdDoc.Overview.MarketContext))
	}

	// Objectives
	md.WriteString("## Objectives\n\n")

	if len(prdDoc.Objectives.BusinessGoals) > 0 {
		md.WriteString("### Business Goals\n\n")
		for _, goal := range prdDoc.Objectives.BusinessGoals {
			md.WriteString(fmt.Sprintf("- %s\n", goal))
		}
		md.WriteString("\n")
	}

	if len(prdDoc.Objectives.SuccessMetrics) > 0 {
		md.WriteString("### Success Metrics\n\n")
		md.WriteString("| Metric | Target | Measurement Method |\n")
		md.WriteString("|--------|--------|-----------------|\n")
		for _, metric := range prdDoc.Objectives.SuccessMetrics {
			method := metric.MeasurementMethod
			if method == "" {
				method = "TBD"
			}
			md.WriteString(fmt.Sprintf("| %s | %s | %s |\n", metric.Metric, metric.Target, method))
		}
		md.WriteString("\n")
	}

	// User Stories
	if len(prdDoc.UserStories) > 0 {
		md.WriteString("## User Stories\n\n")
		for _, story := range prdDoc.UserStories {
			md.WriteString(fmt.Sprintf("### %s", story.ID))
			if story.Priority != "" {
				md.WriteString(fmt.Sprintf(" [%s]", story.Priority))
			}
			if story.EffortEstimate != "" {
				md.WriteString(fmt.Sprintf(" (%s)", story.EffortEstimate))
			}
			md.WriteString("\n\n")
			md.WriteString(fmt.Sprintf("**User Story:** %s\n\n", story.Story))

			if len(story.AcceptanceCriteria) > 0 {
				md.WriteString("**Acceptance Criteria:**\n\n")
				for _, criteria := range story.AcceptanceCriteria {
					md.WriteString(fmt.Sprintf("- %s\n", criteria))
				}
				md.WriteString("\n")
			}
		}
	}

	// Requirements
	md.WriteString("## Requirements\n\n")

	if len(prdDoc.Requirements.Functional) > 0 {
		md.WriteString("### Functional Requirements\n\n")
		md.WriteString("| ID | Description | Priority | Dependencies |\n")
		md.WriteString("|----|-------------|----------|-------------|\n")
		for _, req := range prdDoc.Requirements.Functional {
			deps := strings.Join(req.Dependencies, ", ")
			if deps == "" {
				deps = "None"
			}
			priority := req.Priority
			if priority == "" {
				priority = "TBD"
			}
			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", req.ID, req.Description, priority, deps))
		}
		md.WriteString("\n")
	}

	if len(prdDoc.Requirements.NonFunctional) > 0 {
		md.WriteString("### Non-Functional Requirements\n\n")
		md.WriteString("| ID | Category | Description | Acceptance Criteria |\n")
		md.WriteString("|----|----------|-------------|-------------------|\n")
		for _, req := range prdDoc.Requirements.NonFunctional {
			criteria := req.AcceptanceCriteria
			if criteria == "" {
				criteria = "TBD"
			}
			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", req.ID, req.Category, req.Description, criteria))
		}
		md.WriteString("\n")
	}

	// Technical Specifications
	if prdDoc.TechnicalSpecifications != nil {
		md.WriteString("## Technical Specifications\n\n")

		if prdDoc.TechnicalSpecifications.ArchitectureOverview != "" {
			md.WriteString("### Architecture Overview\n\n")
			md.WriteString(fmt.Sprintf("%s\n\n", prdDoc.TechnicalSpecifications.ArchitectureOverview))
		}

		if prdDoc.TechnicalSpecifications.TechnologyStack != nil {
			stack := prdDoc.TechnicalSpecifications.TechnologyStack
			md.WriteString("### Technology Stack\n\n")

			if len(stack.Frontend) > 0 {
				md.WriteString(fmt.Sprintf("**Frontend:** %s\n\n", strings.Join(stack.Frontend, ", ")))
			}
			if len(stack.Backend) > 0 {
				md.WriteString(fmt.Sprintf("**Backend:** %s\n\n", strings.Join(stack.Backend, ", ")))
			}
			if len(stack.Database) > 0 {
				md.WriteString(fmt.Sprintf("**Database:** %s\n\n", strings.Join(stack.Database, ", ")))
			}
			if len(stack.Infrastructure) > 0 {
				md.WriteString(fmt.Sprintf("**Infrastructure:** %s\n\n", strings.Join(stack.Infrastructure, ", ")))
			}
		}

		if len(prdDoc.TechnicalSpecifications.SecurityConsiderations) > 0 {
			md.WriteString("### Security Considerations\n\n")
			for _, consideration := range prdDoc.TechnicalSpecifications.SecurityConsiderations {
				md.WriteString(fmt.Sprintf("- %s\n", consideration))
			}
			md.WriteString("\n")
		}
	}

	// Timeline
	if prdDoc.Timeline != nil {
		md.WriteString("## Timeline\n\n")

		if prdDoc.Timeline.LaunchDate != "" {
			md.WriteString(fmt.Sprintf("**Target Launch Date:** %s\n\n", prdDoc.Timeline.LaunchDate))
		}

		if len(prdDoc.Timeline.Milestones) > 0 {
			md.WriteString("### Milestones\n\n")
			md.WriteString("| Milestone | Target Date | Description | Dependencies |\n")
			md.WriteString("|-----------|-------------|-------------|-------------|\n")
			for _, milestone := range prdDoc.Timeline.Milestones {
				deps := strings.Join(milestone.Dependencies, ", ")
				if deps == "" {
					deps = "None"
				}
				description := milestone.Description
				if description == "" {
					description = "TBD"
				}
				md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", milestone.Name, milestone.TargetDate, description, deps))
			}
			md.WriteString("\n")
		}
	}

	// Risks and Assumptions
	if prdDoc.RisksAndAssumptions != nil {
		md.WriteString("## Risks and Assumptions\n\n")

		if len(prdDoc.RisksAndAssumptions.Risks) > 0 {
			md.WriteString("### Risks\n\n")
			md.WriteString("| Risk | Impact | Probability | Mitigation Strategy |\n")
			md.WriteString("|------|--------|-------------|-------------------|\n")
			for _, risk := range prdDoc.RisksAndAssumptions.Risks {
				mitigation := risk.MitigationStrategy
				if mitigation == "" {
					mitigation = "TBD"
				}
				md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", risk.Description, risk.Impact, risk.Probability, mitigation))
			}
			md.WriteString("\n")
		}

		if len(prdDoc.RisksAndAssumptions.Assumptions) > 0 {
			md.WriteString("### Assumptions\n\n")
			for _, assumption := range prdDoc.RisksAndAssumptions.Assumptions {
				md.WriteString(fmt.Sprintf("- %s\n", assumption))
			}
			md.WriteString("\n")
		}
	}

	// Out of Scope
	if len(prdDoc.OutOfScope) > 0 {
		md.WriteString("## Out of Scope\n\n")
		for _, item := range prdDoc.OutOfScope {
			md.WriteString(fmt.Sprintf("- %s\n", item))
		}
		md.WriteString("\n")
	}

	// Footer
	md.WriteString("---\n\n")
	md.WriteString(fmt.Sprintf("*Document generated on %s*\n", time.Now().Format("2006-01-02 15:04")))

	// Write to file
	if err := os.WriteFile(filename, []byte(md.String()), 0644); err != nil {
		return fmt.Errorf("failed to write markdown file: %w", err)
	}

	fmt.Printf(color.GreenString("✅ PRD exported to Markdown: %s\n"), filename)
	return nil
}

// Export PRD to HTML
func exportToHTML(prdDoc *prd.PRD, filename string) error {
	var html strings.Builder

	// HTML Header
	html.WriteString("<!DOCTYPE html>\n")
	html.WriteString("<html lang=\"en\">\n")
	html.WriteString("<head>\n")
	html.WriteString("    <meta charset=\"UTF-8\">\n")
	html.WriteString("    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	html.WriteString(fmt.Sprintf("    <title>%s - PRD</title>\n", prdDoc.Title))
	html.WriteString("    <style>\n")
	html.WriteString(getHTMLCSS())
	html.WriteString("    </style>\n")
	html.WriteString("</head>\n")
	html.WriteString("<body>\n")

	// Header
	html.WriteString("    <div class=\"header\">\n")
	html.WriteString(fmt.Sprintf("        <h1>%s</h1>\n", prdDoc.Title))
	html.WriteString("        <div class=\"metadata\">\n")
	html.WriteString(fmt.Sprintf("            <span class=\"badge\">%s</span>\n", prdDoc.Status))
	if prdDoc.Priority != "" {
		html.WriteString(fmt.Sprintf("            <span class=\"badge priority-%s\">%s</span>\n", prdDoc.Priority, prdDoc.Priority))
	}
	html.WriteString("        </div>\n")
	html.WriteString("    </div>\n\n")

	// Document Info
	html.WriteString("    <div class=\"doc-info\">\n")
	html.WriteString("        <div class=\"info-grid\">\n")
	html.WriteString(fmt.Sprintf("            <div><strong>ID:</strong> %s</div>\n", prdDoc.ID))
	html.WriteString(fmt.Sprintf("            <div><strong>Version:</strong> %s</div>\n", prdDoc.Version))
	html.WriteString(fmt.Sprintf("            <div><strong>Owner:</strong> %s</div>\n", prdDoc.Owner.Name))
	html.WriteString(fmt.Sprintf("            <div><strong>Created:</strong> %s</div>\n", prdDoc.CreatedDate))
	html.WriteString("        </div>\n")
	html.WriteString("    </div>\n\n")

	// Overview
	html.WriteString("    <section class=\"section\">\n")
	html.WriteString("        <h2>Overview</h2>\n")
	html.WriteString("        <div class=\"subsection\">\n")
	html.WriteString("            <h3>Problem Statement</h3>\n")
	html.WriteString(fmt.Sprintf("            <p>%s</p>\n", prdDoc.Overview.ProblemStatement))
	html.WriteString("        </div>\n")
	html.WriteString("        <div class=\"subsection\">\n")
	html.WriteString("            <h3>Solution Summary</h3>\n")
	html.WriteString(fmt.Sprintf("            <p>%s</p>\n", prdDoc.Overview.SolutionSummary))
	html.WriteString("        </div>\n")
	html.WriteString("    </section>\n\n")

	// Objectives
	html.WriteString("    <section class=\"section\">\n")
	html.WriteString("        <h2>Objectives</h2>\n")
	if len(prdDoc.Objectives.BusinessGoals) > 0 {
		html.WriteString("        <div class=\"subsection\">\n")
		html.WriteString("            <h3>Business Goals</h3>\n")
		html.WriteString("            <ul>\n")
		for _, goal := range prdDoc.Objectives.BusinessGoals {
			html.WriteString(fmt.Sprintf("                <li>%s</li>\n", goal))
		}
		html.WriteString("            </ul>\n")
		html.WriteString("        </div>\n")
	}
	html.WriteString("    </section>\n\n")

	// Requirements
	html.WriteString("    <section class=\"section\">\n")
	html.WriteString("        <h2>Requirements</h2>\n")
	if len(prdDoc.Requirements.Functional) > 0 {
		html.WriteString("        <div class=\"subsection\">\n")
		html.WriteString("            <h3>Functional Requirements</h3>\n")
		html.WriteString("            <table>\n")
		html.WriteString("                <thead>\n")
		html.WriteString("                    <tr><th>ID</th><th>Description</th><th>Priority</th></tr>\n")
		html.WriteString("                </thead>\n")
		html.WriteString("                <tbody>\n")
		for _, req := range prdDoc.Requirements.Functional {
			priority := req.Priority
			if priority == "" {
				priority = "TBD"
			}
			html.WriteString(fmt.Sprintf("                    <tr><td>%s</td><td>%s</td><td>%s</td></tr>\n", req.ID, req.Description, priority))
		}
		html.WriteString("                </tbody>\n")
		html.WriteString("            </table>\n")
		html.WriteString("        </div>\n")
	}
	html.WriteString("    </section>\n\n")

	// Footer
	html.WriteString("    <footer>\n")
	html.WriteString(fmt.Sprintf("        <p>Generated on %s</p>\n", time.Now().Format("January 2, 2006 at 3:04 PM")))
	html.WriteString("    </footer>\n")

	html.WriteString("</body>\n")
	html.WriteString("</html>\n")

	// Write to file
	if err := os.WriteFile(filename, []byte(html.String()), 0644); err != nil {
		return fmt.Errorf("failed to write HTML file: %w", err)
	}

	fmt.Printf(color.GreenString("✅ PRD exported to HTML: %s\n"), filename)
	return nil
}

func getHTMLCSS() string {
	return `
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            color: #333;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 2rem;
            border-radius: 10px;
            margin-bottom: 2rem;
            text-align: center;
        }
        .header h1 {
            margin: 0;
            font-size: 2.5rem;
        }
        .metadata {
            margin-top: 1rem;
        }
        .badge {
            display: inline-block;
            padding: 0.25rem 0.75rem;
            background: rgba(255,255,255,0.2);
            border-radius: 20px;
            font-size: 0.875rem;
            margin: 0 0.5rem;
            text-transform: uppercase;
        }
        .priority-critical { background: #e74c3c; }
        .priority-high { background: #e67e22; }
        .priority-medium { background: #f39c12; }
        .priority-low { background: #27ae60; }
        .doc-info {
            background: #f8f9fa;
            padding: 1.5rem;
            border-radius: 8px;
            margin-bottom: 2rem;
        }
        .info-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 1rem;
        }
        .section {
            margin-bottom: 3rem;
        }
        .section h2 {
            color: #2c3e50;
            border-bottom: 3px solid #3498db;
            padding-bottom: 0.5rem;
        }
        .subsection {
            margin: 2rem 0;
        }
        .subsection h3 {
            color: #34495e;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin: 1rem 0;
        }
        th, td {
            text-align: left;
            padding: 0.75rem;
            border-bottom: 1px solid #ddd;
        }
        th {
            background: #f4f4f4;
            font-weight: 600;
        }
        tr:hover {
            background: #f9f9f9;
        }
        ul, ol {
            padding-left: 2rem;
        }
        li {
            margin: 0.5rem 0;
        }
        footer {
            margin-top: 3rem;
            padding-top: 2rem;
            border-top: 1px solid #eee;
            text-align: center;
            color: #666;
            font-size: 0.9rem;
        }
    `
}
