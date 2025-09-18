package main

import (
	"fmt"
	"os"

	"github.com/grokify/product-artifacts/prd"
)

func main() {
	runDemo()
}

// Demo function to showcase PRD capabilities
func runDemo() {
	fmt.Println("🚀 PRD Manager - Demonstration")
	fmt.Println("===============================")

	// 1. Create a sample PRD using the Go structs
	fmt.Println("1. Creating a sample PRD...")
	samplePRD := createSamplePRD()

	// 2. Validate the PRD
	fmt.Println("2. Validating PRD...")
	if err := samplePRD.Validate(); err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ PRD validation passed")

	// 3. Save to file
	filename := "demo_prd.json"
	fmt.Printf("3. Saving PRD to %s...\n", filename)
	if err := samplePRD.SaveToFile(filename); err != nil {
		fmt.Printf("❌ Failed to save: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ PRD saved successfully\n")

	// 4. Load from file
	fmt.Printf("4. Loading PRD from %s...\n", filename)
	loadedPRD, err := prd.LoadFromFile(filename)
	if err != nil {
		fmt.Printf("❌ Failed to load: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ PRD loaded successfully")

	// 5. Display PRD information
	fmt.Println("\n5. PRD Summary:")
	fmt.Println("===============")
	displayPRDSummary(loadedPRD)

	// 6. Convert to JSON
	fmt.Println("\n6. JSON Export Sample:")
	fmt.Println("======================")
	jsonStr, err := loadedPRD.ToJSON()
	if err != nil {
		fmt.Printf("❌ Failed to convert to JSON: %v\n", err)
		os.Exit(1)
	}

	// Show first 500 characters of JSON
	if len(jsonStr) > 500 {
		fmt.Printf("%s...\n", jsonStr[:500])
	} else {
		fmt.Println(jsonStr)
	}

	fmt.Println("\n🎉 Demo completed successfully!")
	fmt.Printf("📄 Full PRD saved as: %s\n", filename)
}

func createSamplePRD() *prd.PRD {
	return &prd.PRD{
		ID:          "PRD-DEMO-2024",
		Title:       "Smart Task Management Feature",
		Version:     "1.0.0",
		CreatedDate: "2024-01-15",
		Owner: prd.Owner{
			Name:  "Alex Johnson",
			Email: "alex.johnson@company.com",
			Team:  "Product Team",
		},
		Status:   "approved",
		Priority: "high",
		Overview: prd.Overview{
			ProblemStatement: "Users struggle to prioritize and manage their daily tasks effectively, leading to decreased productivity and missed deadlines.",
			SolutionSummary:  "Implement an AI-powered task management system that automatically prioritizes tasks, suggests optimal scheduling, and provides intelligent reminders.",
			TargetAudience:   "Professionals and teams who manage multiple projects and tasks daily",
			MarketContext:    "Task management market is growing 15% annually with increasing demand for intelligent automation",
		},
		Objectives: prd.Objectives{
			BusinessGoals: []string{
				"Increase user engagement by 30%",
				"Reduce task completion time by 25%",
				"Improve user retention rate by 20%",
			},
			SuccessMetrics: []prd.SuccessMetric{
				{
					Metric:            "Daily Active Users",
					Target:            "50% increase within 6 months",
					MeasurementMethod: "Analytics dashboard tracking",
				},
				{
					Metric:            "Task Completion Rate",
					Target:            "85% completion rate",
					MeasurementMethod: "User behavior analytics",
				},
			},
		},
		UserPersonas: []prd.UserPersona{
			{
				Name:        "Busy Professional",
				Description: "Mid-level manager juggling multiple projects and deadlines",
				Goals: []string{
					"Optimize daily schedule",
					"Never miss important deadlines",
					"Reduce time spent on task management",
				},
				PainPoints: []string{
					"Too many competing priorities",
					"Difficulty estimating task duration",
					"Forgetting important tasks",
				},
			},
		},
		UserStories: []prd.UserStory{
			{
				ID:    "US-001",
				Story: "As a busy professional, I want the system to automatically prioritize my tasks so that I can focus on the most important work first",
				AcceptanceCriteria: []string{
					"System analyzes task deadlines, importance, and dependencies",
					"Tasks are automatically reordered based on priority algorithm",
					"User can override automatic prioritization if needed",
				},
				Priority:       "must_have",
				EffortEstimate: "8 story points",
			},
			{
				ID:    "US-002",
				Story: "As a team member, I want intelligent scheduling suggestions so that I can optimize my daily schedule",
				AcceptanceCriteria: []string{
					"System considers my calendar availability",
					"Suggests optimal time blocks for different types of tasks",
					"Provides estimated completion times",
				},
				Priority:       "should_have",
				EffortEstimate: "5 story points",
			},
		},
		Requirements: prd.Requirements{
			Functional: []prd.FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "System shall provide automatic task prioritization based on deadline, importance, and dependencies",
					Priority:    "must_have",
				},
				{
					ID:           "FR-002",
					Description:  "System shall integrate with calendar applications to suggest optimal scheduling",
					Priority:     "should_have",
					Dependencies: []string{"FR-001"},
				},
				{
					ID:          "FR-003",
					Description: "System shall provide intelligent reminders and notifications",
					Priority:    "must_have",
				},
			},
			NonFunctional: []prd.NonFunctionalRequirement{
				{
					ID:                 "NFR-001",
					Category:           "performance",
					Description:        "Task prioritization algorithm must complete within 2 seconds",
					AcceptanceCriteria: "Load testing shows 95th percentile response time < 2 seconds",
				},
				{
					ID:                 "NFR-002",
					Category:           "usability",
					Description:        "Interface must be intuitive with minimal learning curve",
					AcceptanceCriteria: "User testing shows 90% task completion rate without training",
				},
			},
		},
		TechnicalSpecifications: &prd.TechnicalSpecifications{
			ArchitectureOverview: "Microservices architecture with machine learning pipeline for intelligent prioritization and scheduling",
			TechnologyStack: &prd.TechnologyStack{
				Frontend:       []string{"React", "TypeScript", "Tailwind CSS"},
				Backend:        []string{"Node.js", "Express", "Python ML services"},
				Database:       []string{"PostgreSQL", "Redis"},
				Infrastructure: []string{"AWS", "Docker", "Kubernetes"},
			},
			SecurityConsiderations: []string{
				"End-to-end encryption for sensitive task data",
				"OAuth 2.0 authentication with major calendar providers",
				"GDPR compliance for user data handling",
			},
		},
		Timeline: &prd.Timeline{
			Milestones: []prd.Milestone{
				{
					Name:        "MVP Development",
					Description: "Core task management and basic prioritization",
					TargetDate:  "2024-03-15",
				},
				{
					Name:         "AI Integration",
					Description:  "Machine learning algorithm for intelligent prioritization",
					TargetDate:   "2024-04-30",
					Dependencies: []string{"MVP Development"},
				},
				{
					Name:         "Calendar Integration",
					Description:  "Integration with Google Calendar, Outlook, and Apple Calendar",
					TargetDate:   "2024-06-01",
					Dependencies: []string{"AI Integration"},
				},
			},
			LaunchDate: "2024-07-15",
		},
		RisksAndAssumptions: &prd.RisksAndAssumptions{
			Risks: []prd.Risk{
				{
					Description:        "Machine learning algorithm may not achieve target accuracy",
					Impact:             "medium",
					Probability:        "low",
					MitigationStrategy: "Extensive training data collection and algorithm tuning",
				},
				{
					Description:        "Calendar API limitations may restrict functionality",
					Impact:             "high",
					Probability:        "medium",
					MitigationStrategy: "Early API testing and fallback solutions",
				},
			},
			Assumptions: []string{
				"Users are willing to grant calendar access permissions",
				"Third-party APIs will remain stable during development",
				"User adoption of AI-suggested prioritization will be high",
			},
		},
		OutOfScope: []string{
			"Project management features (Gantt charts, resource allocation)",
			"Team collaboration tools (chat, file sharing)",
			"Mobile app (planned for Phase 2)",
		},
	}
}

func displayPRDSummary(prdDoc *prd.PRD) {
	fmt.Printf("📋 Title: %s\n", prdDoc.Title)
	fmt.Printf("🆔 ID: %s\n", prdDoc.ID)
	fmt.Printf("📦 Version: %s\n", prdDoc.Version)
	fmt.Printf("📊 Status: %s\n", prdDoc.Status)
	fmt.Printf("⚡ Priority: %s\n", prdDoc.Priority)
	fmt.Printf("👤 Owner: %s (%s)\n", prdDoc.Owner.Name, prdDoc.Owner.Email)
	fmt.Printf("📅 Created: %s\n", prdDoc.CreatedDate)

	fmt.Println("\n📈 Statistics:")
	fmt.Printf("• Business Goals: %d\n", len(prdDoc.Objectives.BusinessGoals))
	fmt.Printf("• Success Metrics: %d\n", len(prdDoc.Objectives.SuccessMetrics))
	fmt.Printf("• User Personas: %d\n", len(prdDoc.UserPersonas))
	fmt.Printf("• User Stories: %d\n", len(prdDoc.UserStories))
	fmt.Printf("• Functional Requirements: %d\n", len(prdDoc.Requirements.Functional))
	fmt.Printf("• Non-Functional Requirements: %d\n", len(prdDoc.Requirements.NonFunctional))

	if prdDoc.Timeline != nil {
		fmt.Printf("• Milestones: %d\n", len(prdDoc.Timeline.Milestones))
		fmt.Printf("• Launch Date: %s\n", prdDoc.Timeline.LaunchDate)
	}

	if prdDoc.RisksAndAssumptions != nil {
		fmt.Printf("• Identified Risks: %d\n", len(prdDoc.RisksAndAssumptions.Risks))
		fmt.Printf("• Key Assumptions: %d\n", len(prdDoc.RisksAndAssumptions.Assumptions))
	}
}
