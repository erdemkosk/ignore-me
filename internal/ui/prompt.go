package ui

import (
	"strings"
	"text/template"

	"github.com/erdemkosk/ignore-me/internal/config"
	"github.com/manifoldco/promptui"
)

type PromptData struct {
	Value       string
	Aliases     []string
	Suggestions []string
}

func CreateLanguagePrompt(suggestions []string) *promptui.Select {
	var items []PromptData

	// Dilleri hazırla
	for _, lang := range config.Languages {
		items = append(items, PromptData{
			Value:       lang.Name,
			Aliases:     lang.Aliases,
			Suggestions: suggestions,
		})
	}

	return &promptui.Select{
		Label: "Select Language/Framework",
		Items: items,
		Size:  10,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U0001F539 {{ if containsAny .Value .Aliases .Suggestions }}⭐ {{end}}{{ .Value | cyan }}",
			Inactive: "  {{ if containsAny .Value .Aliases .Suggestions }}⭐ {{end}}{{ .Value | white }}",
			Selected: "\U0001F4E6 {{ .Value | green | bold }} has been selected",
			Details: `
--------- Language/Framework ----------
{{ "Name:" | faint }}	{{ .Value }}
{{ if .Aliases }}{{ "Aliases:" | faint }}	{{ join .Aliases ", " }}{{ end }}
{{ "Template:" | faint }}	{{ .Value }} template will be downloaded from GitHub
`,
			FuncMap: template.FuncMap{
				"containsAny": func(value string, aliases []string, suggestions []string) bool {
					// Önce tam eşleşme kontrolü
					for _, v := range suggestions {
						if v == value {
							return true
						}
					}
					// Sonra alias kontrolü
					for _, alias := range aliases {
						for _, v := range suggestions {
							if v == alias {
								return true
							}
						}
					}
					return false
				},
				"join":  strings.Join,
				"cyan":  promptui.Styler(promptui.FGCyan),
				"white": promptui.Styler(promptui.FGWhite),
				"green": promptui.Styler(promptui.FGGreen),
				"bold":  promptui.Styler(promptui.FGBold),
				"faint": promptui.Styler(promptui.FGFaint),
			},
		},
		Searcher: func(input string, index int) bool {
			item := items[index]
			input = strings.ToLower(input)

			// Ana isimde ara
			if strings.Contains(strings.ToLower(item.Value), input) {
				return true
			}

			// Alias'larda ara
			for _, alias := range item.Aliases {
				if strings.Contains(strings.ToLower(alias), input) {
					return true
				}
			}

			return false
		},
	}
}
