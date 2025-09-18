package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/grokify/product-artifacts/prd"
)

// Template commands
var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available PRD templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listTemplates()
	},
}

var templateCreateCmd = &cobra.Command{
	Use:   "create <template-name> <filename>",
	Short: "Create a PRD from a template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return createFromTemplate(args[1], args[0])
	},
}

var templateShowCmd = &cobra.Command{
	Use:   "show <template-name>",
	Short: "Show template details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return showTemplate(args[0])
	},
}

func listTemplates() error {
	fmt.Printf("%s\n", color.CyanString("📋 Available PRD Templates"))
	fmt.Println("═══════════════════════════════════════════════")

	templates := map[string]string{
		"basic":   "Basic PRD template with minimal required fields",
		"feature": "Feature-focused PRD template for new product features",
		"epic":    "Epic-scale PRD template for major product initiatives",
	}

	for name, description := range templates {
		fmt.Printf("• %s - %s\n", color.YellowString(name), description)
	}

	fmt.Println("\nUsage: prd-manager template create <template-name> <filename>")
	return nil
}

func showTemplate(templateName string) error {
	templates := map[string]*prd.PRD{
		"basic":   createBasicTemplate(),
		"feature": createFeatureTemplate(),
		"epic":    createEpicTemplate(),
	}

	template, exists := templates[templateName]
	if !exists {
		return fmt.Errorf("template '%s' not found", templateName)
	}

	fmt.Printf("%s: %s\n", color.CyanString("Template"), templateName)
	displayPRDPretty(template, "")
	return nil
}

func createBasicTemplate() *prd.PRD {
	return &prd.PRD{
		ID:          "PRD-BASIC-TEMPLATE",
		Title:       "[TEMPLATE] Basic Product Feature",
		Version:     "1.0.0",
		CreatedDate: time.Now().Format("2006-01-02"),
		Owner: prd.Owner{
			Name:  "[OWNER_NAME]",
			Email: "[OWNER_EMAIL]",
			Team:  "[TEAM_NAME]",
		},
		Status:   "draft",
		Priority: "medium",
		Overview: prd.Overview{
			ProblemStatement: "Define the specific problem or opportunity that this product/feature addresses. What user pain points are we solving?",
			SolutionSummary:  "Provide a high-level description of the proposed solution. How will this solve the identified problem?",
			TargetAudience:   "Describe the primary users or customers who will benefit from this product/feature.",
		},
		Objectives: prd.Objectives{
			BusinessGoals: []string{
				"[PRIMARY_BUSINESS_GOAL] - e.g., Increase user retention by 15%",
				"[SECONDARY_BUSINESS_GOAL] - e.g., Reduce support tickets by 20%",
			},
			SuccessMetrics: []prd.SuccessMetric{
				{
					Metric:            "Primary Success Metric",
					Target:            "[TARGET_VALUE] - e.g., 15% increase",
					MeasurementMethod: "How will this be measured? - e.g., Monthly active users",
				},
			},
		},
		Requirements: prd.Requirements{
			Functional: []prd.FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "Define the first functional requirement",
					Priority:    "must_have",
				},
				{
					ID:          "FR-002",
					Description: "Define the second functional requirement",
					Priority:    "should_have",
				},
			},
			NonFunctional: []prd.NonFunctionalRequirement{
				{
					ID:          "NFR-001",
					Category:    "performance",
					Description: "Define performance requirements - e.g., Page load time < 2 seconds",
				},
			},
		},
	}
}

func createFeatureTemplate() *prd.PRD {
	return &prd.PRD{
		ID:          "PRD-FEATURE-TEMPLATE",
		Title:       "[TEMPLATE] New Product Feature",
		Version:     "1.0.0",
		CreatedDate: time.Now().Format("2006-01-02"),
		Owner: prd.Owner{
			Name:  "[PRODUCT_MANAGER_NAME]",
			Email: "[PM_EMAIL]",
			Team:  "[PRODUCT_TEAM]",
		},
		Stakeholders: []prd.Stakeholder{
			{
				Name: "[ENGINEERING_LEAD]",
				Role: "stakeholder",
				Team: "Engineering",
			},
			{
				Name: "[DESIGN_LEAD]",
				Role: "contributor",
				Team: "Design",
			},
		},
		Status:   "draft",
		Priority: "high",
		Overview: prd.Overview{
			ProblemStatement: "Users are struggling with [SPECIFIC_PROBLEM]. Current data shows [SUPPORTING_METRICS]. This is impacting [BUSINESS_IMPACT].",
			SolutionSummary:  "We will build [FEATURE_NAME] that allows users to [KEY_CAPABILITY]. This will solve the problem by [SOLUTION_MECHANISM].",
			TargetAudience:   "Primary: [PRIMARY_USER_SEGMENT] - [CHARACTERISTICS]\nSecondary: [SECONDARY_USER_SEGMENT] - [CHARACTERISTICS]",
			MarketContext:    "Competitive analysis shows [COMPETITOR_LANDSCAPE]. Market opportunity: [MARKET_SIZE/OPPORTUNITY].",
		},
		Objectives: prd.Objectives{
			BusinessGoals: []string{
				"Increase user engagement metrics",
				"Reduce user friction in core workflow",
				"Drive revenue growth through improved conversion",
			},
			SuccessMetrics: []prd.SuccessMetric{
				{
					Metric:            "Feature Adoption Rate",
					Target:            "60% of active users within 3 months",
					MeasurementMethod: "Analytics tracking of feature usage",
				},
				{
					Metric:            "Task Completion Time",
					Target:            "30% reduction in average task time",
					MeasurementMethod: "User session analytics and A/B testing",
				},
			},
		},
		UserPersonas: []prd.UserPersona{
			{
				Name:        "Primary User Persona",
				Description: "[PERSONA_NAME] - [AGE_RANGE] - [ROLE/TITLE] - [KEY_CHARACTERISTICS]",
				Goals: []string{
					"[PRIMARY_GOAL]",
					"[SECONDARY_GOAL]",
				},
				PainPoints: []string{
					"[MAIN_PAIN_POINT]",
					"[SECONDARY_PAIN_POINT]",
				},
			},
		},
		UserStories: []prd.UserStory{
			{
				ID:    "US-001",
				Story: "As a [USER_TYPE], I want to [DESIRED_ACTION] so that [BENEFIT/VALUE]",
				AcceptanceCriteria: []string{
					"Given [PRECONDITION], when [ACTION], then [EXPECTED_RESULT]",
					"[ADDITIONAL_CRITERIA]",
				},
				Priority:       "must_have",
				EffortEstimate: "[STORY_POINTS]",
			},
		},
		Requirements: prd.Requirements{
			Functional: []prd.FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "System shall provide [CORE_FUNCTIONALITY]",
					Priority:    "must_have",
				},
				{
					ID:          "FR-002",
					Description: "System shall support [SECONDARY_FUNCTIONALITY]",
					Priority:    "should_have",
				},
				{
					ID:           "FR-003",
					Description:  "System shall integrate with [EXTERNAL_SYSTEM]",
					Priority:     "must_have",
					Dependencies: []string{"FR-001"},
				},
			},
			NonFunctional: []prd.NonFunctionalRequirement{
				{
					ID:                 "NFR-001",
					Category:           "performance",
					Description:        "Feature response time shall be under 500ms for 95% of requests",
					AcceptanceCriteria: "Load testing shows 95th percentile < 500ms",
				},
				{
					ID:                 "NFR-002",
					Category:           "usability",
					Description:        "Feature shall be accessible to users with disabilities",
					AcceptanceCriteria: "WCAG 2.1 Level AA compliance verified",
				},
			},
		},
		TechnicalSpecifications: &prd.TechnicalSpecifications{
			ArchitectureOverview: "Feature will be implemented as [ARCHITECTURE_PATTERN]. Integration points: [INTEGRATION_DETAILS].",
			TechnologyStack: &prd.TechnologyStack{
				Frontend: []string{"[FRONTEND_TECH]", "[UI_FRAMEWORK]"},
				Backend:  []string{"[BACKEND_TECH]", "[API_FRAMEWORK]"},
				Database: []string{"[DATABASE_TYPE]"},
			},
			SecurityConsiderations: []string{
				"Data encryption for sensitive information",
				"Authentication and authorization controls",
				"Input validation and sanitization",
			},
		},
		Timeline: &prd.Timeline{
			Milestones: []prd.Milestone{
				{
					Name:        "Design & Planning Complete",
					Description: "UI/UX designs approved, technical design finalized",
					TargetDate:  "[DATE]",
				},
				{
					Name:         "MVP Development Complete",
					Description:  "Core functionality implemented and tested",
					TargetDate:   "[DATE]",
					Dependencies: []string{"Design & Planning Complete"},
				},
				{
					Name:         "Beta Release",
					Description:  "Feature available to beta users for testing",
					TargetDate:   "[DATE]",
					Dependencies: []string{"MVP Development Complete"},
				},
			},
			LaunchDate: "[LAUNCH_DATE]",
		},
		RisksAndAssumptions: &prd.RisksAndAssumptions{
			Risks: []prd.Risk{
				{
					Description:        "Technical complexity may lead to delays",
					Impact:             "medium",
					Probability:        "medium",
					MitigationStrategy: "Conduct technical spike, break into smaller phases",
				},
				{
					Description:        "User adoption may be lower than expected",
					Impact:             "high",
					Probability:        "low",
					MitigationStrategy: "Conduct user research, A/B test different approaches",
				},
			},
			Assumptions: []string{
				"Users are familiar with similar features in other products",
				"Current infrastructure can support the additional load",
				"Third-party integrations will remain stable",
			},
		},
		OutOfScope: []string{
			"[FEATURE_NOT_INCLUDED] - planned for future release",
			"[ANOTHER_EXCLUSION] - requires separate project",
		},
	}
}

func createEpicTemplate() *prd.PRD {
	return &prd.PRD{
		ID:          "PRD-EPIC-TEMPLATE",
		Title:       "[TEMPLATE] Major Product Initiative",
		Version:     "1.0.0",
		CreatedDate: time.Now().Format("2006-01-02"),
		Owner: prd.Owner{
			Name:  "[SENIOR_PM_NAME]",
			Email: "[SENIOR_PM_EMAIL]",
			Team:  "[PRODUCT_ORGANIZATION]",
		},
		Stakeholders: []prd.Stakeholder{
			{
				Name: "[EXECUTIVE_SPONSOR]",
				Role: "approver",
				Team: "Leadership",
			},
			{
				Name: "[ENGINEERING_DIRECTOR]",
				Role: "stakeholder",
				Team: "Engineering",
			},
			{
				Name: "[DESIGN_DIRECTOR]",
				Role: "stakeholder",
				Team: "Design",
			},
			{
				Name: "[MARKETING_LEAD]",
				Role: "contributor",
				Team: "Marketing",
			},
		},
		Status:   "review",
		Priority: "critical",
		Overview: prd.Overview{
			ProblemStatement: "Market research indicates [MARKET_OPPORTUNITY]. Current product limitations: [CURRENT_LIMITATIONS]. Customer feedback shows: [CUSTOMER_FEEDBACK]. Business impact: [REVENUE_IMPACT].",
			SolutionSummary:  "Launch comprehensive [PRODUCT_AREA] initiative including [MAJOR_COMPONENTS]. This multi-phase project will [TRANSFORMATION_GOAL].",
			TargetAudience:   "Primary: [MARKET_SEGMENT] worth $[MARKET_SIZE]\nSecondary: [EXPANSION_SEGMENT]\nTertiary: [FUTURE_SEGMENT]",
			MarketContext:    "Market size: $[SIZE]. Growth rate: [RATE]%. Key competitors: [COMPETITORS]. Our differentiation: [UNIQUE_VALUE_PROP].",
		},
		Objectives: prd.Objectives{
			BusinessGoals: []string{
				"Drive $[REVENUE_TARGET] in additional annual revenue",
				"Expand market share in [TARGET_SEGMENT] by [PERCENTAGE]%",
				"Establish leadership position in [PRODUCT_CATEGORY]",
				"Improve customer satisfaction scores by [IMPROVEMENT]",
			},
			SuccessMetrics: []prd.SuccessMetric{
				{
					Metric:            "Revenue Impact",
					Target:            "$[AMOUNT] ARR within 12 months",
					MeasurementMethod: "Financial reporting and customer analytics",
				},
				{
					Metric:            "Market Share",
					Target:            "[PERCENTAGE]% of [MARKET_SEGMENT]",
					MeasurementMethod: "Third-party market research and internal data",
				},
				{
					Metric:            "Customer Satisfaction",
					Target:            "NPS score improvement of [POINTS] points",
					MeasurementMethod: "Quarterly customer surveys",
				},
			},
			OKRs: []prd.OKR{
				{
					Objective: "Establish market-leading [PRODUCT_CAPABILITY]",
					KeyResults: []string{
						"Launch [CAPABILITY] to 100% of customers",
						"Achieve [METRIC] improvement in [MEASUREMENT]",
						"Secure [NUMBER] strategic customer wins",
					},
				},
				{
					Objective: "Drive significant revenue growth",
					KeyResults: []string{
						"Generate $[AMOUNT] in new revenue",
						"Increase average contract value by [PERCENTAGE]%",
						"Improve customer retention to [PERCENTAGE]%",
					},
				},
			},
		},
		UserPersonas: []prd.UserPersona{
			{
				Name:        "Enterprise Decision Maker",
				Description: "C-level or VP responsible for [DOMAIN] strategy and budget decisions",
				Goals: []string{
					"Improve [BUSINESS_OUTCOME] for organization",
					"Reduce operational costs and complexity",
					"Stay competitive in rapidly evolving market",
				},
				PainPoints: []string{
					"Current solutions don't scale with business growth",
					"Integration challenges with existing systems",
					"Lack of comprehensive analytics and insights",
				},
			},
			{
				Name:        "Power User",
				Description: "Daily user who leverages advanced features for [USE_CASE]",
				Goals: []string{
					"Increase productivity and efficiency",
					"Access advanced capabilities and customization",
					"Streamline complex workflows",
				},
				PainPoints: []string{
					"Feature limitations prevent optimal workflows",
					"Time-consuming manual processes",
					"Inadequate reporting and analysis tools",
				},
			},
		},
		Requirements: prd.Requirements{
			Functional: []prd.FunctionalRequirement{
				{
					ID:          "FR-001",
					Description: "Core platform foundation supporting [SCALE] concurrent users",
					Priority:    "must_have",
				},
				{
					ID:          "FR-002",
					Description: "Advanced analytics and reporting engine",
					Priority:    "must_have",
				},
				{
					ID:          "FR-003",
					Description: "Enterprise-grade security and compliance features",
					Priority:    "must_have",
				},
				{
					ID:          "FR-004",
					Description: "API ecosystem for third-party integrations",
					Priority:    "should_have",
				},
				{
					ID:          "FR-005",
					Description: "Mobile application with offline capabilities",
					Priority:    "should_have",
				},
			},
			NonFunctional: []prd.NonFunctionalRequirement{
				{
					ID:                 "NFR-001",
					Category:           "scalability",
					Description:        "System shall support [NUMBER] concurrent users with [RESPONSE_TIME] response time",
					AcceptanceCriteria: "Load testing demonstrates sustained performance under peak load",
				},
				{
					ID:                 "NFR-002",
					Category:           "security",
					Description:        "Achieve [COMPLIANCE_STANDARD] certification",
					AcceptanceCriteria: "Third-party audit confirms compliance",
				},
				{
					ID:                 "NFR-003",
					Category:           "reliability",
					Description:        "System uptime of 99.9% with disaster recovery < 4 hours",
					AcceptanceCriteria: "SLA monitoring and incident response testing",
				},
			},
		},
		TechnicalSpecifications: &prd.TechnicalSpecifications{
			ArchitectureOverview: "Microservices architecture with [CLOUD_PROVIDER] infrastructure. Event-driven design with [MESSAGE_BROKER]. Multi-tenant SaaS platform.",
			TechnologyStack: &prd.TechnologyStack{
				Frontend:       []string{"[FRONTEND_FRAMEWORK]", "[MOBILE_PLATFORM]", "[WEB_FRAMEWORK]"},
				Backend:        []string{"[BACKEND_LANGUAGE]", "[API_FRAMEWORK]", "[MESSAGE_QUEUE]"},
				Database:       []string{"[PRIMARY_DB]", "[ANALYTICS_DB]", "[CACHE_LAYER]"},
				Infrastructure: []string{"[CLOUD_PROVIDER]", "[CONTAINER_PLATFORM]", "[MONITORING_STACK]"},
			},
			SecurityConsiderations: []string{
				"End-to-end encryption for all data transmission",
				"Role-based access control with enterprise SSO",
				"Regular security audits and penetration testing",
				"Compliance with [REGULATORY_REQUIREMENTS]",
				"Data residency and sovereignty requirements",
			},
		},
		Timeline: &prd.Timeline{
			Milestones: []prd.Milestone{
				{
					Name:        "Phase 1: Foundation",
					Description: "Core platform infrastructure and basic functionality",
					TargetDate:  "[Q1_DATE]",
				},
				{
					Name:         "Phase 2: Advanced Features",
					Description:  "Analytics, integrations, and enterprise features",
					TargetDate:   "[Q2_DATE]",
					Dependencies: []string{"Phase 1: Foundation"},
				},
				{
					Name:         "Phase 3: Scale & Polish",
					Description:  "Performance optimization, mobile app, and market launch",
					TargetDate:   "[Q3_DATE]",
					Dependencies: []string{"Phase 2: Advanced Features"},
				},
			},
			LaunchDate: "[GENERAL_AVAILABILITY_DATE]",
		},
		RisksAndAssumptions: &prd.RisksAndAssumptions{
			Risks: []prd.Risk{
				{
					Description:        "Technical complexity may exceed current team capabilities",
					Impact:             "high",
					Probability:        "medium",
					MitigationStrategy: "Hire additional senior engineers, engage technical consultants",
				},
				{
					Description:        "Market conditions may change during development",
					Impact:             "high",
					Probability:        "low",
					MitigationStrategy: "Quarterly market reviews, flexible roadmap planning",
				},
				{
					Description:        "Competitive response may impact differentiation",
					Impact:             "medium",
					Probability:        "high",
					MitigationStrategy: "Accelerate unique features, build patent portfolio",
				},
			},
			Assumptions: []string{
				"Market demand will remain stable throughout development",
				"Key talent will be available for hiring",
				"Technology choices will remain viable for 3+ years",
				"Regulatory environment will not significantly change",
				"Customer budget allocation for [CATEGORY] will increase",
			},
		},
		OutOfScope: []string{
			"International localization - Phase 4 initiative",
			"Advanced AI/ML capabilities - separate R&D project",
			"Acquisition of external companies or technologies",
			"On-premises deployment options - cloud-first strategy",
		},
		Appendices: &prd.Appendices{
			ResearchData: "Market research conducted by [FIRM]. Customer interviews: [NUMBER] participants. Competitive analysis: [DATE].",
			RelatedDocuments: []prd.RelatedDocument{
				{
					Title: "Market Research Report",
					Type:  "market_research",
				},
				{
					Title: "Technical Architecture Design",
					Type:  "technical_spec",
				},
				{
					Title: "Financial Business Case",
					Type:  "business_case",
				},
			},
		},
	}
}
