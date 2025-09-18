package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/grokify/product-artifacts/prd"
)

// Interactive PRD creation wizard
func createInteractivePRD(filename string) error {
	fmt.Println(color.CyanString("🚀 Welcome to the PRD Creation Wizard!"))
	fmt.Println("Let's create a comprehensive Product Requirements Document.")

	reader := bufio.NewReader(os.Stdin)

	// Basic Information
	fmt.Println(color.YellowString("📋 Basic Information"))
	fmt.Print("PRD ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Product/Feature Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Version (e.g., 1.0.0): ")
	version, _ := reader.ReadString('\n')
	version = strings.TrimSpace(version)

	// Owner Information
	fmt.Println(color.YellowString("\n👤 Owner Information"))
	fmt.Print("Owner Name: ")
	ownerName, _ := reader.ReadString('\n')
	ownerName = strings.TrimSpace(ownerName)

	fmt.Print("Owner Email: ")
	ownerEmail, _ := reader.ReadString('\n')
	ownerEmail = strings.TrimSpace(ownerEmail)

	fmt.Print("Owner Team: ")
	ownerTeam, _ := reader.ReadString('\n')
	ownerTeam = strings.TrimSpace(ownerTeam)

	// Status and Priority
	fmt.Println(color.YellowString("\n📊 Status & Priority"))
	status := selectFromOptions("Status", []string{"draft", "review", "approved", "in_development", "completed"})
	priority := selectFromOptions("Priority", []string{"critical", "high", "medium", "low"})

	// Overview
	fmt.Println(color.YellowString("\n📝 Product Overview"))
	fmt.Print("Problem Statement: ")
	problemStatement, _ := reader.ReadString('\n')
	problemStatement = strings.TrimSpace(problemStatement)

	fmt.Print("Solution Summary: ")
	solutionSummary, _ := reader.ReadString('\n')
	solutionSummary = strings.TrimSpace(solutionSummary)

	fmt.Print("Target Audience: ")
	targetAudience, _ := reader.ReadString('\n')
	targetAudience = strings.TrimSpace(targetAudience)

	// Business Goals
	fmt.Println(color.YellowString("\n🎯 Business Goals"))
	businessGoals := collectMultipleInputs("Business Goal", 3)

	// Functional Requirements
	fmt.Println(color.YellowString("\n⚙️ Functional Requirements"))
	functionalReqs := collectFunctionalRequirements()

	// Create PRD object
	now := time.Now()
	prdDoc := &prd.PRD{
		ID:          id,
		Title:       title,
		Version:     version,
		CreatedDate: now.Format("2006-01-02"),
		LastUpdated: &now,
		Owner: prd.Owner{
			Name:  ownerName,
			Email: ownerEmail,
			Team:  ownerTeam,
		},
		Status:   status,
		Priority: priority,
		Overview: prd.Overview{
			ProblemStatement: problemStatement,
			SolutionSummary:  solutionSummary,
			TargetAudience:   targetAudience,
		},
		Objectives: prd.Objectives{
			BusinessGoals: businessGoals,
		},
		Requirements: prd.Requirements{
			Functional: functionalReqs,
		},
	}

	// Validate and save
	if err := prdDoc.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := prdDoc.SaveToFile(filename); err != nil {
		return fmt.Errorf("failed to save PRD: %w", err)
	}

	fmt.Printf(color.GreenString("\n✅ PRD successfully created: %s\n"), filename)
	return nil
}

// Create basic PRD with minimal input
func createBasicPRD(filename string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("PRD Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Your Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	now := time.Now()
	prdDoc := &prd.PRD{
		ID:          fmt.Sprintf("PRD-%d", now.Unix()),
		Title:       title,
		Version:     "1.0.0",
		CreatedDate: now.Format("2006-01-02"),
		LastUpdated: &now,
		Owner: prd.Owner{
			Name:  name,
			Email: email,
		},
		Status: "draft",
		Overview: prd.Overview{
			ProblemStatement: "TODO: Define the problem this product/feature solves",
			SolutionSummary:  "TODO: Describe the proposed solution",
		},
		Objectives: prd.Objectives{
			BusinessGoals: []string{"TODO: Add business goals"},
		},
		Requirements: prd.Requirements{
			Functional: []prd.FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "TODO: Define functional requirements",
					Priority:    "must_have",
				},
			},
		},
	}

	if err := prdDoc.SaveToFile(filename); err != nil {
		return fmt.Errorf("failed to save PRD: %w", err)
	}

	fmt.Printf(color.GreenString("✅ Basic PRD template created: %s\n"), filename)
	fmt.Println("Use 'prd-manager edit' to add more details.")
	return nil
}

// Create from template
func createFromTemplate(filename, templateType string) error {
	templates := map[string]*prd.PRD{
		"basic":   createBasicTemplate(),
		"feature": createFeatureTemplate(),
		"epic":    createEpicTemplate(),
	}

	template, exists := templates[templateType]
	if !exists {
		return fmt.Errorf("template '%s' not found. Available: basic, feature, epic", templateType)
	}

	now := time.Now()
	template.LastUpdated = &now
	template.CreatedDate = now.Format("2006-01-02")

	if err := template.SaveToFile(filename); err != nil {
		return fmt.Errorf("failed to save PRD: %w", err)
	}

	fmt.Printf(color.GreenString("✅ PRD created from '%s' template: %s\n"), templateType, filename)
	return nil
}

// List PRDs in directory
func listPRDs(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		fmt.Println("No PRD files found in directory.")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("File", "ID", "Title", "Status", "Owner", "Last Updated")

	for _, file := range files {
		prdDoc, err := prd.LoadFromFile(file)
		if err != nil {
			continue // Skip invalid files
		}

		lastUpdated := "N/A"
		if prdDoc.LastUpdated != nil {
			lastUpdated = prdDoc.LastUpdated.Format("2006-01-02 15:04")
		}

		err = table.Append([]string{
			filepath.Base(file),
			prdDoc.ID,
			truncateString(prdDoc.Title, 30),
			prdDoc.Status,
			prdDoc.Owner.Name,
			lastUpdated,
		})
		if err != nil {
			return err
		}
	}

	fmt.Println(color.CyanString("📋 PRD Documents"))
	return table.Render()
}

// View PRD content
func viewPRD(filename, format, section string) error {
	prdDoc, err := prd.LoadFromFile(filename)
	if err != nil {
		return err
	}

	switch format {
	case "json":
		jsonStr, err := prdDoc.ToJSON()
		if err != nil {
			return err
		}
		fmt.Println(jsonStr)
	case "table":
		return displayPRDTable(prdDoc, section)
	default: // pretty
		displayPRDPretty(prdDoc, section)
	}

	return nil
}

// Edit PRD
func editPRD(filename, section string) error {
	prdDoc, err := prd.LoadFromFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf(color.CyanString("📝 Editing PRD: %s\n"), prdDoc.Title)

	if section == "" {
		section = selectFromOptions("Section to edit", []string{
			"basic", "overview", "objectives", "requirements", "timeline", "risks",
		})
	}

	switch section {
	case "basic":
		editBasicInfo(prdDoc)
	case "overview":
		editOverview(prdDoc)
	case "objectives":
		editObjectives(prdDoc)
	case "requirements":
		editRequirements(prdDoc)
	default:
		return fmt.Errorf("section '%s' not supported for editing", section)
	}

	prdDoc.UpdateLastModified()

	if err := prdDoc.SaveToFile(filename); err != nil {
		return fmt.Errorf("failed to save PRD: %w", err)
	}

	fmt.Printf(color.GreenString("✅ PRD updated: %s\n"), filename)
	return nil
}

// Validate PRD
func validatePRD(filename string, strict bool) error {
	prdDoc, err := prd.LoadFromFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf(color.CyanString("🔍 Validating PRD: %s\n"), prdDoc.Title)

	// Basic validation
	if err := prdDoc.Validate(); err != nil {
		fmt.Printf(color.RedString("❌ Validation failed: %v\n"), err)
		return err
	}

	// Additional checks for strict mode
	warnings := []string{}
	if strict {
		if len(prdDoc.UserStories) == 0 {
			warnings = append(warnings, "No user stories defined")
		}
		if prdDoc.Timeline == nil {
			warnings = append(warnings, "No timeline specified")
		}
		if len(prdDoc.Requirements.NonFunctional) == 0 {
			warnings = append(warnings, "No non-functional requirements")
		}
	}

	fmt.Println(color.GreenString("✅ PRD validation passed"))

	if len(warnings) > 0 {
		fmt.Println(color.YellowString("\n⚠️ Warnings:"))
		for _, warning := range warnings {
			fmt.Printf("  • %s\n", warning)
		}
	}

	return nil
}

// Show PRD status
func showPRDStatus(filename string) error {
	prdDoc, err := prd.LoadFromFile(filename)
	if err != nil {
		return err
	}

	fmt.Printf(color.CyanString("📊 PRD Status Report: %s\n\n"), prdDoc.Title)

	// Basic info
	fmt.Printf("ID: %s\n", prdDoc.ID)
	fmt.Printf("Version: %s\n", prdDoc.Version)
	fmt.Printf("Status: %s\n", getStatusWithColor(prdDoc.Status))
	fmt.Printf("Priority: %s\n", getPriorityWithColor(prdDoc.Priority))
	fmt.Printf("Owner: %s (%s)\n", prdDoc.Owner.Name, prdDoc.Owner.Email)

	if prdDoc.LastUpdated != nil {
		fmt.Printf("Last Updated: %s\n", prdDoc.LastUpdated.Format("2006-01-02 15:04"))
	}

	// Statistics
	fmt.Printf("\n📈 Statistics:\n")
	fmt.Printf("• Functional Requirements: %d\n", len(prdDoc.Requirements.Functional))
	fmt.Printf("• Non-Functional Requirements: %d\n", len(prdDoc.Requirements.NonFunctional))
	fmt.Printf("• User Stories: %d\n", len(prdDoc.UserStories))
	fmt.Printf("• Business Goals: %d\n", len(prdDoc.Objectives.BusinessGoals))
	fmt.Printf("• Stakeholders: %d\n", len(prdDoc.Stakeholders))

	if prdDoc.Timeline != nil {
		fmt.Printf("• Milestones: %d\n", len(prdDoc.Timeline.Milestones))
	}

	return nil
}

// Export PRD to different formats
func exportPRD(filename, format, output string) error {
	prdDoc, err := prd.LoadFromFile(filename)
	if err != nil {
		return err
	}

	if output == "" {
		ext := map[string]string{
			"markdown": ".md",
			"html":     ".html",
			"pdf":      ".pdf",
		}
		output = strings.TrimSuffix(filename, ".json") + ext[format]
	}

	switch format {
	case "markdown":
		return exportToMarkdown(prdDoc, output)
	case "html":
		return exportToHTML(prdDoc, output)
	default:
		return fmt.Errorf("export format '%s' not supported", format)
	}
}

// Helper functions
func selectFromOptions(prompt string, options []string) string {
	fmt.Printf("%s:\n", prompt)
	for i, option := range options {
		fmt.Printf("  %d. %s\n", i+1, option)
	}
	fmt.Print("Select (1-" + strconv.Itoa(len(options)) + "): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if choice, err := strconv.Atoi(input); err == nil && choice > 0 && choice <= len(options) {
		return options[choice-1]
	}

	return options[0] // Default to first option
}

func collectMultipleInputs(itemType string, maxItems int) []string {
	var items []string
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < maxItems; i++ {
		fmt.Printf("%s %d (or press Enter to finish): ", itemType, i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}
		items = append(items, input)
	}

	return items
}

func collectFunctionalRequirements() []prd.FunctionalRequirement {
	var requirements []prd.FunctionalRequirement
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 5; i++ {
		fmt.Printf("Functional Requirement %d (or press Enter to finish): ", i+1)
		desc, _ := reader.ReadString('\n')
		desc = strings.TrimSpace(desc)

		if desc == "" {
			break
		}

		priority := selectFromOptions("Priority", []string{"must_have", "should_have", "could_have", "wont_have"})

		requirements = append(requirements, prd.FunctionalRequirement{
			ID:          fmt.Sprintf("FR-%03d", i+1),
			Description: desc,
			Priority:    priority,
		})
	}

	return requirements
}

func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}

func getStatusWithColor(status string) string {
	colors := map[string]*color.Color{
		"draft":          color.New(color.FgYellow),
		"review":         color.New(color.FgCyan),
		"approved":       color.New(color.FgGreen),
		"in_development": color.New(color.FgBlue),
		"completed":      color.New(color.FgGreen, color.Bold),
		"archived":       color.New(color.FgHiBlack),
	}

	if c, exists := colors[status]; exists {
		return c.Sprint(status)
	}
	return status
}

func getPriorityWithColor(priority string) string {
	colors := map[string]*color.Color{
		"critical": color.New(color.FgRed, color.Bold),
		"high":     color.New(color.FgRed),
		"medium":   color.New(color.FgYellow),
		"low":      color.New(color.FgGreen),
	}

	if c, exists := colors[priority]; exists {
		return c.Sprint(priority)
	}
	return priority
}
