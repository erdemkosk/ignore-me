package generator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type GitignoreGenerator struct {
	languageMap map[string]string
}

func NewGitignoreGenerator() *GitignoreGenerator {
	return &GitignoreGenerator{
		languageMap: map[string]string{
			"CPP":     "C++",
			"CSharp":  "C Sharp",
			"Flutter": "Dart",
		},
	}
}

func (g *GitignoreGenerator) CreateGitignore(language string) error {
	// Show spinner while downloading
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "📥 "
	s.Suffix = " Downloading template for " + language
	s.Color("cyan")
	s.Start()

	content, err := g.getTemplate(language)
	s.Stop()
	if err != nil {
		return err
	}

	// Check if file exists
	if _, err := os.Stat(".gitignore"); err == nil {
		prompt := promptui.Prompt{
			Label:     "⚠️  .gitignore already exists. Do you want to overwrite it",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil || strings.ToLower(result) != "y" {
			color.Yellow("Operation cancelled")
			return nil
		}
	}

	// Create progress spinner for file creation
	s = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "💾 "
	s.Suffix = " Creating .gitignore file..."
	s.Color("cyan")
	s.Start()

	err = ioutil.WriteFile(".gitignore", []byte(content), 0644)
	s.Stop()

	return err
}

func (g *GitignoreGenerator) getTemplate(language string) (string, error) {
	if mapped, exists := g.languageMap[language]; exists {
		language = mapped
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", language)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("could not download template: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return "", fmt.Errorf("template not found for %s", language)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read content: %v", err)
	}

	return string(content), nil
}
