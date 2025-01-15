package cmd

import (
	"github.com/erdemkosk/ignore-me/internal/analyzer"
	"github.com/erdemkosk/ignore-me/internal/config"
	"github.com/erdemkosk/ignore-me/internal/generator"
	"github.com/erdemkosk/ignore-me/internal/ui"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ignore-me",
	Short: "A beautiful gitignore generator",
	Long: `Ignore Me is a CLI tool that helps you create .gitignore files
for your projects with ease. It supports multiple programming languages
and frameworks.`,
	Run: func(cmd *cobra.Command, args []string) {
		showLanguages()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func showLanguages() {
	headerColor := color.New(color.FgCyan, color.Bold)
	headerColor.Println("\n🚀 Ignore Me - Gitignore Generator")
	headerColor.Println("====================================")

	// Analyze project
	projectAnalyzer := analyzer.NewProjectAnalyzer()
	suggestions := projectAnalyzer.Analyze()

	if len(suggestions) > 0 {
		color.Yellow("\n📝 Project Analysis Results:")
		color.Yellow("Found these technologies in your project:")
		for _, s := range suggestions {
			color.Yellow("  • %s", s)
		}
		color.Yellow("\n💡 Tip: Recommended templates are marked with ⭐ in the list below")
		color.Yellow("")
	}

	// Show prompt and get selection
	prompt := ui.CreateLanguagePrompt(suggestions)
	index, _, err := prompt.Run()
	if err != nil {
		color.Red("❌ Error: %v", err)
		return
	}

	selectedLang := config.Languages[index].GitHubName

	// Generate gitignore
	gen := generator.NewGitignoreGenerator()
	if err := gen.CreateGitignore(selectedLang); err != nil {
		color.Red("❌ Error: %v", err)
		return
	}

	successColor := color.New(color.FgGreen, color.Bold)
	successColor.Printf("\n✨ Success! .gitignore file created for %s!\n\n", config.Languages[index].Name)
}
