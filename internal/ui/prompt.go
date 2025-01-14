package ui

import (
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
)

type PromptData struct {
	Value       string
	Suggestions []string
}

func CreateLanguagePrompt(items []PromptData) *promptui.Select {
	return &promptui.Select{
		Label: "Select Language/Framework",
		Items: items,
		Size:  10,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U0001F539 {{ if contains .Value .Suggestions }}⭐ {{end}}{{ .Value | cyan }} ({{ .Value | cyan }})",
			Inactive: "  {{ if contains .Value .Suggestions }}⭐ {{end}}{{ .Value | white }}",
			Selected: "\U0001F4E6 {{ .Value | green | bold }} has been selected",
			Details: `
--------- Language/Framework ----------
{{ "Name:" | faint }}	{{ .Value }}
{{ "Template:" | faint }}	{{ .Value }} template will be downloaded from GitHub
`,
			FuncMap: template.FuncMap{
				"contains": func(s string, suggestions []string) bool {
					for _, v := range suggestions {
						if v == s {
							return true
						}
					}
					return false
				},
				"cyan":  promptui.Styler(promptui.FGCyan),
				"white": promptui.Styler(promptui.FGWhite),
				"green": promptui.Styler(promptui.FGGreen),
				"bold":  promptui.Styler(promptui.FGBold),
				"faint": promptui.Styler(promptui.FGFaint),
			},
		},
		Searcher: func(input string, index int) bool {
			lang := strings.ToLower(items[index].Value)
			input = strings.ToLower(input)
			return strings.Contains(lang, input)
		},
	}
}
