package config

type Language struct {
	Name       string   // Görüntülenecek isim
	Aliases    []string // Alternatif isimler
	GitHubName string   // GitHub template adı
}

var Languages = []Language{
	{
		Name:       "Go",
		Aliases:    []string{"Golang"},
		GitHubName: "Go",
	},
	{
		Name:       "Python",
		Aliases:    []string{"py", "python3"},
		GitHubName: "Python",
	},
	{
		Name:       "Node.js",
		Aliases:    []string{"Node", "JavaScript", "JS"},
		GitHubName: "Node",
	},
	{
		Name:       "Java",
		Aliases:    []string{"JVM"},
		GitHubName: "Java",
	},
	{
		Name:       "Ruby",
		Aliases:    []string{"rb"},
		GitHubName: "Ruby",
	},
	{
		Name:       "Rust",
		Aliases:    []string{"rs"},
		GitHubName: "Rust",
	},
	{
		Name:       "C",
		Aliases:    []string{"gcc"},
		GitHubName: "C",
	},
	{
		Name:       "C++",
		Aliases:    []string{"cpp", "CPP"},
		GitHubName: "C++",
	},
	{
		Name:       "C#",
		Aliases:    []string{"csharp", "CSharp", "cs"},
		GitHubName: "CSharp",
	},
	{
		Name:       "PHP",
		Aliases:    []string{"php"},
		GitHubName: "PHP",
	},
	{
		Name:       "Swift",
		Aliases:    []string{"swift"},
		GitHubName: "Swift",
	},
	{
		Name:       "Kotlin",
		Aliases:    []string{"kt"},
		GitHubName: "Kotlin",
	},
	{
		Name:       "Dart",
		Aliases:    []string{"Flutter"},
		GitHubName: "Dart",
	},
	{
		Name:       "React",
		Aliases:    []string{"reactjs", "react.js"},
		GitHubName: "React",
	},
	{
		Name:       "Vue",
		Aliases:    []string{"vuejs", "vue.js"},
		GitHubName: "Vue",
	},
	{
		Name:       "Angular",
		Aliases:    []string{"ng", "angular.js", "angularjs"},
		GitHubName: "Angular",
	},
}
