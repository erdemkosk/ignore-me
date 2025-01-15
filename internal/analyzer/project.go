package analyzer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"time"

	"github.com/briandowns/spinner"
)

type ProjectAnalyzer struct {
	currentDir string
	fileMap    map[string]bool
	extensions map[string]bool
}

func NewProjectAnalyzer() *ProjectAnalyzer {
	return &ProjectAnalyzer{
		fileMap:    make(map[string]bool),
		extensions: make(map[string]bool),
	}
}

func (pa *ProjectAnalyzer) Analyze() []string {
	var suggestions []string
	currentDir, err := os.Getwd()
	if err != nil {
		return suggestions
	}
	pa.currentDir = currentDir

	// Create a spinner for analysis
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "🔍 "
	s.Suffix = " Analyzing project structure..."
	s.Color("cyan")
	s.Start()
	defer s.Stop()

	if err := pa.scanDirectory(); err != nil {
		return suggestions
	}

	return pa.detectTechnologies()
}

func (pa *ProjectAnalyzer) scanDirectory() error {
	files, err := os.ReadDir(pa.currentDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		pa.fileMap[file.Name()] = true
	}

	return filepath.Walk(pa.currentDir, pa.processFile)
}

func (pa *ProjectAnalyzer) processFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}
	if !info.IsDir() {
		ext := strings.ToLower(filepath.Ext(path))
		if ext != "" {
			pa.extensions[ext] = true
		}
	}
	return nil
}

func (pa *ProjectAnalyzer) detectTechnologies() []string {
	var suggestions []string
	rootFiles := pa.getRootFiles()
	rootExtensions := pa.getRootExtensions()

	// Go project detection
	if rootFiles["go.mod"] || rootFiles["go.sum"] || rootExtensions[".go"] {
		suggestions = append(suggestions, "Go")
	}

	// Python project detection
	if rootFiles["requirements.txt"] || rootFiles["setup.py"] ||
		rootFiles["Pipfile"] || rootExtensions[".py"] {
		suggestions = append(suggestions, "Python")
	}

	// Node.js/JavaScript project detection
	if rootFiles["package.json"] || rootFiles["node_modules"] ||
		rootExtensions[".js"] || rootExtensions[".ts"] {
		suggestions = append(suggestions, "Node")
	}

	// Java project detection
	if rootFiles["pom.xml"] || rootFiles["build.gradle"] ||
		rootExtensions[".java"] || rootFiles["gradlew"] {
		suggestions = append(suggestions, "Java")
	}

	// Ruby project detection
	if pa.fileMap["Gemfile"] || pa.fileMap["Rakefile"] ||
		pa.extensions[".rb"] || pa.fileMap[".ruby-version"] {
		suggestions = append(suggestions, "Ruby")
	}

	// Rust project detection
	if pa.fileMap["Cargo.toml"] || pa.fileMap["Cargo.lock"] ||
		pa.extensions[".rs"] {
		suggestions = append(suggestions, "Rust")
	}

	// C project detection
	if pa.extensions[".c"] || pa.extensions[".h"] {
		suggestions = append(suggestions, "C")
	}

	// C++ project detection
	if pa.extensions[".cpp"] || pa.extensions[".hpp"] ||
		pa.extensions[".cc"] || pa.extensions[".cxx"] {
		suggestions = append(suggestions, "CPP")
	}

	// C# project detection
	if pa.extensions[".cs"] || pa.fileMap["*.csproj"] ||
		pa.fileMap["*.sln"] {
		suggestions = append(suggestions, "CSharp")
	}

	// PHP project detection
	if pa.fileMap["composer.json"] || pa.extensions[".php"] ||
		pa.fileMap["artisan"] {
		suggestions = append(suggestions, "PHP")
	}

	// Swift project detection
	if pa.extensions[".swift"] || pa.fileMap["Package.swift"] ||
		pa.fileMap["*.xcodeproj"] || pa.fileMap["*.xcworkspace"] {
		suggestions = append(suggestions, "Swift")
	}

	// Kotlin project detection
	if pa.extensions[".kt"] || pa.extensions[".kts"] ||
		pa.fileMap["*.gradle.kts"] {
		suggestions = append(suggestions, "Kotlin")
	}

	// Dart/Flutter project detection
	if pa.fileMap["pubspec.yaml"] || pa.extensions[".dart"] {
		suggestions = append(suggestions, "Flutter")
	}

	// React project detection
	if pa.fileExists("src/App.js") || pa.fileExists("src/App.tsx") ||
		(pa.fileMap["package.json"] && pa.containsInFile("package.json", "react")) {
		suggestions = append(suggestions, "React")
	}

	// Vue project detection
	if pa.fileExists("src/App.vue") ||
		(pa.fileMap["package.json"] && pa.containsInFile("package.json", "vue")) {
		suggestions = append(suggestions, "Vue")
	}

	// Angular project detection
	if pa.fileMap["angular.json"] ||
		(pa.fileMap["package.json"] && pa.containsInFile("package.json", "angular")) {
		suggestions = append(suggestions, "Angular")
	}

	return suggestions
}

// Helper method to check if file exists
func (pa *ProjectAnalyzer) fileExists(path string) bool {
	_, err := os.Stat(filepath.Join(pa.currentDir, path))
	return err == nil
}

// Helper method to check if file contains string
func (pa *ProjectAnalyzer) containsInFile(filename, search string) bool {
	content, err := ioutil.ReadFile(filepath.Join(pa.currentDir, filename))
	if err != nil {
		return false
	}
	return strings.Contains(string(content), search)
}

// Sadece kök dizindeki dosyaları kontrol et
func (pa *ProjectAnalyzer) getRootFiles() map[string]bool {
	rootFiles := make(map[string]bool)
	files, err := os.ReadDir(pa.currentDir)
	if err != nil {
		return rootFiles
	}

	for _, file := range files {
		if !file.IsDir() {
			rootFiles[file.Name()] = true
		} else {
			// Dizin ise direkt dizin adını da ekle
			rootFiles[file.Name()] = true
		}
	}
	return rootFiles
}

// Sadece kök dizindeki uzantıları kontrol et
func (pa *ProjectAnalyzer) getRootExtensions() map[string]bool {
	rootExts := make(map[string]bool)
	files, err := os.ReadDir(pa.currentDir)
	if err != nil {
		return rootExts
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext != "" {
				rootExts[ext] = true
			}
		}
	}
	return rootExts
}
