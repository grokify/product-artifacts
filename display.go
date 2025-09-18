package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/grokify/product-artifacts/prd"
)

// Display PRD in pretty format
func displayPRDPretty(prdDoc *prd.PRD, section string) {
	if section == "" || section == "header" {
		displayHeader(prdDoc)
	}

	if section == "" || section == "overview" {
		displayOverview(prdDoc)
	}

	if section == "" || section == "objectives" {
		displayObjectives(prdDoc)
	}

	if section == "" || section == "requirements" {
		displayRequirements(prdDoc)
	}

	if section == "" || section == "stories" {
		displayUserStories(prdDoc)
	}

	if section == "" || section == "timeline" {
		displayTimeline(prdDoc)
	}

	if section == "" || section == "technical" {
		displayTechnicalSpecs(prdDoc)
	}

	if section == "" || section == "risks" {
		displayRisks(prdDoc)
	}
}

func displayHeader(prdDoc *prd.PRD) {
	fmt.Printf("═══════════════════════════════════════════════════════════════\n")
	fmt.Printf("📋 %s\n", color.CyanString(prdDoc.Title))
	fmt.Printf("═══════════════════════════════════════════════════════════════\n\n")

	fmt.Printf("🆔 ID: %s\n", prdDoc.ID)
	fmt.Printf("📦 Version: %s\n", prdDoc.Version)
	fmt.Printf("📊 Status: %s\n", getStatusWithColor(prdDoc.Status))
	if prdDoc.Priority != "" {
		fmt.Printf("⚡ Priority: %s\n", getPriorityWithColor(prdDoc.Priority))
	}
	fmt.Printf("👤 Owner: %s (%s)\n", prdDoc.Owner.Name, prdDoc.Owner.Email)
	if prdDoc.Owner.Team != "" {
		fmt.Printf("🏢 Team: %s\n", prdDoc.Owner.Team)
	}
	fmt.Printf("📅 Created: %s\n", prdDoc.CreatedDate)
	if prdDoc.LastUpdated != nil {
		fmt.Printf("🔄 Last Updated: %s\n", prdDoc.LastUpdated.Format("2006-01-02 15:04"))
	}
	fmt.Println()
}

func displayOverview(prdDoc *prd.PRD) {
	fmt.Printf("%s\n", color.YellowString("📝 OVERVIEW"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	fmt.Printf("🎯 Problem Statement:\n%s\n\n", wrapText(prdDoc.Overview.ProblemStatement, 60))
	fmt.Printf("💡 Solution Summary:\n%s\n\n", wrapText(prdDoc.Overview.SolutionSummary, 60))

	if prdDoc.Overview.TargetAudience != "" {
		fmt.Printf("👥 Target Audience:\n%s\n\n", wrapText(prdDoc.Overview.TargetAudience, 60))
	}

	if prdDoc.Overview.MarketContext != "" {
		fmt.Printf("📈 Market Context:\n%s\n\n", wrapText(prdDoc.Overview.MarketContext, 60))
	}
}

func displayObjectives(prdDoc *prd.PRD) {
	fmt.Printf("%s\n", color.YellowString("🎯 OBJECTIVES"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	if len(prdDoc.Objectives.BusinessGoals) > 0 {
		fmt.Printf("🏢 Business Goals:\n")
		for i, goal := range prdDoc.Objectives.BusinessGoals {
			fmt.Printf("  %d. %s\n", i+1, goal)
		}
		fmt.Println()
	}

	if len(prdDoc.Objectives.SuccessMetrics) > 0 {
		fmt.Printf("📊 Success Metrics:\n")
		for _, metric := range prdDoc.Objectives.SuccessMetrics {
			fmt.Printf("  • %s: %s\n", color.CyanString(metric.Metric), metric.Target)
			if metric.MeasurementMethod != "" {
				fmt.Printf("    Method: %s\n", metric.MeasurementMethod)
			}
		}
		fmt.Println()
	}

	if len(prdDoc.Objectives.OKRs) > 0 {
		fmt.Printf("🎯 OKRs:\n")
		for i, okr := range prdDoc.Objectives.OKRs {
			fmt.Printf("  %s %s\n", color.CyanString(fmt.Sprintf("O%d:", i+1)), okr.Objective)
			for j, kr := range okr.KeyResults {
				fmt.Printf("    %s %s\n", color.GreenString(fmt.Sprintf("KR%d.%d:", i+1, j+1)), kr)
			}
		}
		fmt.Println()
	}
}

func displayRequirements(prdDoc *prd.PRD) {
	fmt.Printf("%s\n", color.YellowString("⚙️ REQUIREMENTS"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	if len(prdDoc.Requirements.Functional) > 0 {
		fmt.Printf("🔧 Functional Requirements:\n")
		for _, req := range prdDoc.Requirements.Functional {
			priority := ""
			if req.Priority != "" {
				priority = fmt.Sprintf(" [%s]", getPriorityWithColor(req.Priority))
			}
			fmt.Printf("  • %s%s\n", color.CyanString(req.ID), priority)
			fmt.Printf("    %s\n", wrapText(req.Description, 58))
			if len(req.Dependencies) > 0 {
				fmt.Printf("    Dependencies: %s\n", strings.Join(req.Dependencies, ", "))
			}
		}
		fmt.Println()
	}

	if len(prdDoc.Requirements.NonFunctional) > 0 {
		fmt.Printf("🛡️ Non-Functional Requirements:\n")
		for _, req := range prdDoc.Requirements.NonFunctional {
			fmt.Printf("  • %s [%s]\n", color.CyanString(req.ID), req.Category)
			fmt.Printf("    %s\n", wrapText(req.Description, 58))
			if req.AcceptanceCriteria != "" {
				fmt.Printf("    Criteria: %s\n", req.AcceptanceCriteria)
			}
		}
		fmt.Println()
	}
}

func displayUserStories(prdDoc *prd.PRD) {
	if len(prdDoc.UserStories) == 0 {
		return
	}

	fmt.Printf("%s\n", color.YellowString("📖 USER STORIES"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	for _, story := range prdDoc.UserStories {
		priority := ""
		if story.Priority != "" {
			priority = fmt.Sprintf(" [%s]", getPriorityWithColor(story.Priority))
		}
		estimate := ""
		if story.EffortEstimate != "" {
			estimate = fmt.Sprintf(" (%s)", story.EffortEstimate)
		}

		fmt.Printf("📝 %s%s%s\n", color.CyanString(story.ID), priority, estimate)
		fmt.Printf("   %s\n", story.Story)

		if len(story.AcceptanceCriteria) > 0 {
			fmt.Printf("   Acceptance Criteria:\n")
			for _, criteria := range story.AcceptanceCriteria {
				fmt.Printf("   ✓ %s\n", criteria)
			}
		}
		fmt.Println()
	}
}

func displayTimeline(prdDoc *prd.PRD) {
	if prdDoc.Timeline == nil {
		return
	}

	fmt.Printf("%s\n", color.YellowString("📅 TIMELINE"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	if prdDoc.Timeline.LaunchDate != "" {
		fmt.Printf("🚀 Launch Date: %s\n\n", color.GreenString(prdDoc.Timeline.LaunchDate))
	}

	if len(prdDoc.Timeline.Milestones) > 0 {
		fmt.Printf("🏁 Milestones:\n")
		for _, milestone := range prdDoc.Timeline.Milestones {
			fmt.Printf("  • %s - %s\n", color.CyanString(milestone.TargetDate), milestone.Name)
			if milestone.Description != "" {
				fmt.Printf("    %s\n", milestone.Description)
			}
			if len(milestone.Dependencies) > 0 {
				fmt.Printf("    Dependencies: %s\n", strings.Join(milestone.Dependencies, ", "))
			}
		}
		fmt.Println()
	}
}

func displayTechnicalSpecs(prdDoc *prd.PRD) {
	if prdDoc.TechnicalSpecifications == nil {
		return
	}

	fmt.Printf("%s\n", color.YellowString("🔧 TECHNICAL SPECIFICATIONS"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	if prdDoc.TechnicalSpecifications.ArchitectureOverview != "" {
		fmt.Printf("🏗️ Architecture:\n%s\n\n", wrapText(prdDoc.TechnicalSpecifications.ArchitectureOverview, 60))
	}

	if prdDoc.TechnicalSpecifications.TechnologyStack != nil {
		stack := prdDoc.TechnicalSpecifications.TechnologyStack
		fmt.Printf("💻 Technology Stack:\n")

		if len(stack.Frontend) > 0 {
			fmt.Printf("  Frontend: %s\n", strings.Join(stack.Frontend, ", "))
		}
		if len(stack.Backend) > 0 {
			fmt.Printf("  Backend: %s\n", strings.Join(stack.Backend, ", "))
		}
		if len(stack.Database) > 0 {
			fmt.Printf("  Database: %s\n", strings.Join(stack.Database, ", "))
		}
		if len(stack.Infrastructure) > 0 {
			fmt.Printf("  Infrastructure: %s\n", strings.Join(stack.Infrastructure, ", "))
		}
		fmt.Println()
	}

	if len(prdDoc.TechnicalSpecifications.SecurityConsiderations) > 0 {
		fmt.Printf("🔒 Security Considerations:\n")
		for _, consideration := range prdDoc.TechnicalSpecifications.SecurityConsiderations {
			fmt.Printf("  • %s\n", consideration)
		}
		fmt.Println()
	}
}

func displayRisks(prdDoc *prd.PRD) {
	if prdDoc.RisksAndAssumptions == nil {
		return
	}

	fmt.Printf("%s\n", color.YellowString("⚠️ RISKS & ASSUMPTIONS"))
	fmt.Printf("─────────────────────────────────────────────────────────────\n")

	if len(prdDoc.RisksAndAssumptions.Risks) > 0 {
		fmt.Printf("⚠️ Risks:\n")
		for _, risk := range prdDoc.RisksAndAssumptions.Risks {
			impactColor := color.New(color.FgYellow)
			if risk.Impact == "high" || risk.Impact == "critical" {
				impactColor = color.New(color.FgRed)
			} else if risk.Impact == "low" {
				impactColor = color.New(color.FgGreen)
			}

			fmt.Printf("  • %s [Impact: %s, Probability: %s]\n",
				risk.Description,
				impactColor.Sprint(risk.Impact),
				risk.Probability)

			if risk.MitigationStrategy != "" {
				fmt.Printf("    Mitigation: %s\n", risk.MitigationStrategy)
			}
		}
		fmt.Println()
	}

	if len(prdDoc.RisksAndAssumptions.Assumptions) > 0 {
		fmt.Printf("📋 Assumptions:\n")
		for _, assumption := range prdDoc.RisksAndAssumptions.Assumptions {
			fmt.Printf("  • %s\n", assumption)
		}
		fmt.Println()
	}
}

// Display PRD in table format
func displayPRDTable(prdDoc *prd.PRD, section string) {
	switch section {
	case "requirements":
		displayRequirementsTable(prdDoc)
	case "stories":
		displayUserStoriesTable(prdDoc)
	case "milestones":
		displayMilestonesTable(prdDoc)
	default:
		displayOverviewTable(prdDoc)
	}
}

func displayRequirementsTable(prdDoc *prd.PRD) {
	if len(prdDoc.Requirements.Functional) == 0 {
		fmt.Println("No functional requirements found.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description", "Priority", "Dependencies"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetColWidth(50)

	for _, req := range prdDoc.Requirements.Functional {
		deps := strings.Join(req.Dependencies, ", ")
		if deps == "" {
			deps = "None"
		}

		table.Append([]string{
			req.ID,
			truncateString(req.Description, 40),
			req.Priority,
			deps,
		})
	}

	fmt.Println(color.CyanString("Functional Requirements"))
	table.Render()
}

func displayUserStoriesTable(prdDoc *prd.PRD) {
	if len(prdDoc.UserStories) == 0 {
		fmt.Println("No user stories found.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Story", "Priority", "Estimate"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetColWidth(40)

	for _, story := range prdDoc.UserStories {
		table.Append([]string{
			story.ID,
			truncateString(story.Story, 35),
			story.Priority,
			story.EffortEstimate,
		})
	}

	fmt.Println(color.CyanString("User Stories"))
	table.Render()
}

func displayMilestonesTable(prdDoc *prd.PRD) {
	if prdDoc.Timeline == nil || len(prdDoc.Timeline.Milestones) == 0 {
		fmt.Println("No milestones found.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Target Date", "Description", "Dependencies"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetColWidth(30)

	for _, milestone := range prdDoc.Timeline.Milestones {
		deps := strings.Join(milestone.Dependencies, ", ")
		if deps == "" {
			deps = "None"
		}

		table.Append([]string{
			milestone.Name,
			milestone.TargetDate,
			truncateString(milestone.Description, 25),
			deps,
		})
	}

	fmt.Println(color.CyanString("Project Milestones"))
	table.Render()
}

func displayOverviewTable(prdDoc *prd.PRD) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetColWidth(50)

	table.Append([]string{"ID", prdDoc.ID})
	table.Append([]string{"Title", prdDoc.Title})
	table.Append([]string{"Version", prdDoc.Version})
	table.Append([]string{"Status", prdDoc.Status})
	table.Append([]string{"Priority", prdDoc.Priority})
	table.Append([]string{"Owner", fmt.Sprintf("%s (%s)", prdDoc.Owner.Name, prdDoc.Owner.Email)})
	table.Append([]string{"Created", prdDoc.CreatedDate})

	if prdDoc.LastUpdated != nil {
		table.Append([]string{"Last Updated", prdDoc.LastUpdated.Format("2006-01-02 15:04")})
	}

	fmt.Println(color.CyanString("PRD Overview"))
	table.Render()
}

// Utility functions
func wrapText(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var result strings.Builder
	words := strings.Fields(text)
	lineLength := 0

	for i, word := range words {
		if i > 0 && lineLength+len(word)+1 > width {
			result.WriteString("\n")
			lineLength = 0
		}

		if lineLength > 0 {
			result.WriteString(" ")
			lineLength++
		}

		result.WriteString(word)
		lineLength += len(word)
	}

	return result.String()
}
