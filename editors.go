package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"github.com/grokify/product-artifacts/prd"
)

// Edit basic information
func editBasicInfo(prdDoc *prd.PRD) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Current Title: %s\n", color.CyanString(prdDoc.Title))
	fmt.Print("New Title (or press Enter to keep current): ")
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		prdDoc.Title = strings.TrimSpace(input)
	}

	fmt.Printf("Current Version: %s\n", color.CyanString(prdDoc.Version))
	fmt.Print("New Version (or press Enter to keep current): ")
	if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
		prdDoc.Version = strings.TrimSpace(input)
	}

	fmt.Printf("Current Status: %s\n", color.CyanString(prdDoc.Status))
	if confirmChange("Do you want to change the status? (y/n): ") {
		prdDoc.Status = selectFromOptions("New Status", []string{
			"draft", "review", "approved", "in_development", "completed", "archived",
		})
	}

	fmt.Printf("Current Priority: %s\n", color.CyanString(prdDoc.Priority))
	if confirmChange("Do you want to change the priority? (y/n): ") {
		prdDoc.Priority = selectFromOptions("New Priority", []string{
			"critical", "high", "medium", "low",
		})
	}

	fmt.Println(color.GreenString("✅ Basic information updated"))
}

// Edit overview section
func editOverview(prdDoc *prd.PRD) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Current Problem Statement:\n%s\n", color.CyanString(prdDoc.Overview.ProblemStatement))
	if confirmChange("Do you want to update the problem statement? (y/n): ") {
		fmt.Print("New Problem Statement: ")
		if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
			prdDoc.Overview.ProblemStatement = strings.TrimSpace(input)
		}
	}

	fmt.Printf("Current Solution Summary:\n%s\n", color.CyanString(prdDoc.Overview.SolutionSummary))
	if confirmChange("Do you want to update the solution summary? (y/n): ") {
		fmt.Print("New Solution Summary: ")
		if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
			prdDoc.Overview.SolutionSummary = strings.TrimSpace(input)
		}
	}

	fmt.Printf("Current Target Audience:\n%s\n", color.CyanString(prdDoc.Overview.TargetAudience))
	if confirmChange("Do you want to update the target audience? (y/n): ") {
		fmt.Print("New Target Audience: ")
		if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
			prdDoc.Overview.TargetAudience = strings.TrimSpace(input)
		}
	}

	fmt.Println(color.GreenString("✅ Overview updated"))
}

// Edit objectives
func editObjectives(prdDoc *prd.PRD) {
	fmt.Println(color.YellowString("📝 Editing Objectives"))

	// Edit business goals
	fmt.Println("\nCurrent Business Goals:")
	for i, goal := range prdDoc.Objectives.BusinessGoals {
		fmt.Printf("  %d. %s\n", i+1, goal)
	}

	if confirmChange("Do you want to modify business goals? (y/n): ") {
		action := selectFromOptions("Action", []string{"add", "remove", "replace_all"})

		switch action {
		case "add":
			newGoals := collectMultipleInputs("New Business Goal", 3)
			prdDoc.Objectives.BusinessGoals = append(prdDoc.Objectives.BusinessGoals, newGoals...)

		case "remove":
			if len(prdDoc.Objectives.BusinessGoals) > 0 {
				idx := selectGoalToRemove(prdDoc.Objectives.BusinessGoals)
				if idx >= 0 && idx < len(prdDoc.Objectives.BusinessGoals) {
					prdDoc.Objectives.BusinessGoals = append(
						prdDoc.Objectives.BusinessGoals[:idx],
						prdDoc.Objectives.BusinessGoals[idx+1:]...)
				}
			}

		case "replace_all":
			newGoals := collectMultipleInputs("Business Goal", 5)
			if len(newGoals) > 0 {
				prdDoc.Objectives.BusinessGoals = newGoals
			}
		}
	}

	// Edit success metrics
	fmt.Println("\nCurrent Success Metrics:")
	for i, metric := range prdDoc.Objectives.SuccessMetrics {
		fmt.Printf("  %d. %s: %s\n", i+1, metric.Metric, metric.Target)
	}

	if confirmChange("Do you want to modify success metrics? (y/n): ") {
		action := selectFromOptions("Action", []string{"add", "remove", "replace_all"})

		switch action {
		case "add":
			newMetrics := collectSuccessMetrics(3)
			prdDoc.Objectives.SuccessMetrics = append(prdDoc.Objectives.SuccessMetrics, newMetrics...)

		case "remove":
			if len(prdDoc.Objectives.SuccessMetrics) > 0 {
				idx := selectMetricToRemove(prdDoc.Objectives.SuccessMetrics)
				if idx >= 0 && idx < len(prdDoc.Objectives.SuccessMetrics) {
					prdDoc.Objectives.SuccessMetrics = append(
						prdDoc.Objectives.SuccessMetrics[:idx],
						prdDoc.Objectives.SuccessMetrics[idx+1:]...)
				}
			}

		case "replace_all":
			newMetrics := collectSuccessMetrics(5)
			if len(newMetrics) > 0 {
				prdDoc.Objectives.SuccessMetrics = newMetrics
			}
		}
	}

	fmt.Println(color.GreenString("✅ Objectives updated"))
}

// Edit requirements
func editRequirements(prdDoc *prd.PRD) {
	fmt.Println(color.YellowString("⚙️ Editing Requirements"))

	// Edit functional requirements
	fmt.Println("\nCurrent Functional Requirements:")
	for i, req := range prdDoc.Requirements.Functional {
		fmt.Printf("  %d. %s: %s [%s]\n", i+1, req.ID, truncateString(req.Description, 50), req.Priority)
	}

	if confirmChange("Do you want to modify functional requirements? (y/n): ") {
		action := selectFromOptions("Action", []string{"add", "remove", "edit", "replace_all"})

		switch action {
		case "add":
			newReqs := collectFunctionalRequirements()
			// Update IDs to be unique
			startID := len(prdDoc.Requirements.Functional) + 1
			for i := range newReqs {
				newReqs[i].ID = fmt.Sprintf("FR-%03d", startID+i)
			}
			prdDoc.Requirements.Functional = append(prdDoc.Requirements.Functional, newReqs...)

		case "remove":
			if len(prdDoc.Requirements.Functional) > 0 {
				idx := selectRequirementToRemove(prdDoc.Requirements.Functional)
				if idx >= 0 && idx < len(prdDoc.Requirements.Functional) {
					prdDoc.Requirements.Functional = append(
						prdDoc.Requirements.Functional[:idx],
						prdDoc.Requirements.Functional[idx+1:]...)
				}
			}

		case "edit":
			if len(prdDoc.Requirements.Functional) > 0 {
				idx := selectRequirementToEdit(prdDoc.Requirements.Functional)
				if idx >= 0 && idx < len(prdDoc.Requirements.Functional) {
					editSingleRequirement(&prdDoc.Requirements.Functional[idx])
				}
			}

		case "replace_all":
			newReqs := collectFunctionalRequirements()
			if len(newReqs) > 0 {
				prdDoc.Requirements.Functional = newReqs
			}
		}
	}

	fmt.Println(color.GreenString("✅ Requirements updated"))
}

// Helper functions for editing
func confirmChange(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

func selectGoalToRemove(goals []string) int {
	fmt.Println("Select goal to remove:")
	for i, goal := range goals {
		fmt.Printf("  %d. %s\n", i+1, truncateString(goal, 60))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if choice, err := strconv.Atoi(input); err == nil && choice > 0 && choice <= len(goals) {
		return choice - 1
	}

	return -1
}

func selectMetricToRemove(metrics []prd.SuccessMetric) int {
	fmt.Println("Select metric to remove:")
	for i, metric := range metrics {
		fmt.Printf("  %d. %s: %s\n", i+1, metric.Metric, metric.Target)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if choice, err := strconv.Atoi(input); err == nil && choice > 0 && choice <= len(metrics) {
		return choice - 1
	}

	return -1
}

func selectRequirementToRemove(requirements []prd.FunctionalRequirement) int {
	fmt.Println("Select requirement to remove:")
	for i, req := range requirements {
		fmt.Printf("  %d. %s: %s\n", i+1, req.ID, truncateString(req.Description, 50))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if choice, err := strconv.Atoi(input); err == nil && choice > 0 && choice <= len(requirements) {
		return choice - 1
	}

	return -1
}

func selectRequirementToEdit(requirements []prd.FunctionalRequirement) int {
	fmt.Println("Select requirement to edit:")
	for i, req := range requirements {
		fmt.Printf("  %d. %s: %s\n", i+1, req.ID, truncateString(req.Description, 50))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if choice, err := strconv.Atoi(input); err == nil && choice > 0 && choice <= len(requirements) {
		return choice - 1
	}

	return -1
}

func editSingleRequirement(req *prd.FunctionalRequirement) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Current Description: %s\n", req.Description)
	if confirmChange("Update description? (y/n): ") {
		fmt.Print("New Description: ")
		if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
			req.Description = strings.TrimSpace(input)
		}
	}

	fmt.Printf("Current Priority: %s\n", req.Priority)
	if confirmChange("Update priority? (y/n): ") {
		req.Priority = selectFromOptions("New Priority", []string{
			"must_have", "should_have", "could_have", "wont_have",
		})
	}

	fmt.Printf("Current Dependencies: %v\n", req.Dependencies)
	if confirmChange("Update dependencies? (y/n): ") {
		fmt.Print("Dependencies (comma-separated, or press Enter for none): ")
		if input, _ := reader.ReadString('\n'); strings.TrimSpace(input) != "" {
			deps := strings.Split(strings.TrimSpace(input), ",")
			for i := range deps {
				deps[i] = strings.TrimSpace(deps[i])
			}
			req.Dependencies = deps
		} else {
			req.Dependencies = []string{}
		}
	}
}

func collectSuccessMetrics(maxMetrics int) []prd.SuccessMetric {
	var metrics []prd.SuccessMetric
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < maxMetrics; i++ {
		fmt.Printf("Success Metric %d (or press Enter to finish): ", i+1)
		metricName, _ := reader.ReadString('\n')
		metricName = strings.TrimSpace(metricName)

		if metricName == "" {
			break
		}

		fmt.Print("Target Value: ")
		target, _ := reader.ReadString('\n')
		target = strings.TrimSpace(target)

		fmt.Print("Measurement Method (optional): ")
		method, _ := reader.ReadString('\n')
		method = strings.TrimSpace(method)

		metrics = append(metrics, prd.SuccessMetric{
			Metric:            metricName,
			Target:            target,
			MeasurementMethod: method,
		})
	}

	return metrics
}
