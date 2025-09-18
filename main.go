package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prd-manager",
	Short: "A comprehensive Product Requirements Document (PRD) management tool",
	Long: `PRD Manager is a CLI tool for creating, managing, and validating
Product Requirements Documents. It provides interactive wizards,
validation, templating, and comprehensive PRD lifecycle management.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add all subcommands
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(viewCmd)
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(exportCmd)
}

// Create command
var createCmd = &cobra.Command{
	Use:   "create [filename]",
	Short: "Create a new PRD document",
	Long:  `Create a new PRD document using an interactive wizard or from a template.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")
		template, _ := cmd.Flags().GetString("template")

		filename := "new_prd.json"
		if len(args) > 0 {
			filename = args[0]
		}

		if interactive {
			return createInteractivePRD(filename)
		} else if template != "" {
			return createFromTemplate(filename, template)
		} else {
			return createBasicPRD(filename)
		}
	},
}

// List command
var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List all PRD documents",
	Long:  `List all PRD documents in the specified directory or current directory.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}
		return listPRDs(dir)
	},
}

// View command
var viewCmd = &cobra.Command{
	Use:   "view <filename>",
	Short: "View a PRD document",
	Long:  `View the contents of a PRD document in various formats.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		section, _ := cmd.Flags().GetString("section")
		return viewPRD(args[0], format, section)
	},
}

// Edit command
var editCmd = &cobra.Command{
	Use:   "edit <filename>",
	Short: "Edit a PRD document",
	Long:  `Edit specific sections of a PRD document interactively.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		section, _ := cmd.Flags().GetString("section")
		return editPRD(args[0], section)
	},
}

// Validate command
var validateCmd = &cobra.Command{
	Use:   "validate <filename>",
	Short: "Validate a PRD document",
	Long:  `Validate a PRD document against the schema and business rules.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		strict, _ := cmd.Flags().GetBool("strict")
		return validatePRD(args[0], strict)
	},
}

// Template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage PRD templates",
	Long:  `Create, list, and manage PRD templates for different product types.`,
}

// Status command
var statusCmd = &cobra.Command{
	Use:   "status <filename>",
	Short: "Show PRD status and statistics",
	Long:  `Display comprehensive status information and statistics about a PRD.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return showPRDStatus(args[0])
	},
}

// Export command
var exportCmd = &cobra.Command{
	Use:   "export <filename>",
	Short: "Export PRD to different formats",
	Long:  `Export a PRD document to various formats like Markdown, HTML, or PDF.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		output, _ := cmd.Flags().GetString("output")
		return exportPRD(args[0], format, output)
	},
}

func init() {
	// Create command flags
	createCmd.Flags().BoolP("interactive", "i", false, "Use interactive wizard")
	createCmd.Flags().StringP("template", "t", "", "Create from template (basic, feature, epic)")

	// View command flags
	viewCmd.Flags().StringP("format", "f", "pretty", "Output format (pretty, json, table)")
	viewCmd.Flags().StringP("section", "s", "", "View specific section (overview, requirements, etc.)")

	// Edit command flags
	editCmd.Flags().StringP("section", "s", "", "Edit specific section")

	// Validate command flags
	validateCmd.Flags().BoolP("strict", "", false, "Use strict validation mode")

	// Export command flags
	exportCmd.Flags().StringP("format", "f", "markdown", "Export format (markdown, html, pdf)")
	exportCmd.Flags().StringP("output", "o", "", "Output filename")

	// Template subcommands
	templateCmd.AddCommand(templateListCmd)
	templateCmd.AddCommand(templateCreateCmd)
	templateCmd.AddCommand(templateShowCmd)
}
