package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewInitCommand(pmFlag *string) *cobra.Command {
	return createManagerCommand(pmFlag, CommandConfig{
		Use:   "init [project-name]",
		Short: "Initialize a new project with the appropriate package manager",
		Args:  cobra.MaximumNArgs(1),
		RunFunc: func(pm models.PackageManager, args []string) error {
			projectName := ""
			if len(args) > 0 {
				projectName = args[0]
			} else {
				// Use current directory name as project name
				cwd, err := os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to get current directory: %w", err)
				}
				projectName = filepath.Base(cwd)
			}

			return initProject(pm, projectName)
		},
	})
}

func initProject(pm models.PackageManager, projectName string) error {
	fmt.Printf("Initializing project '%s' with %s\n", projectName, pm.Name())

	switch pm.Name() {
	case "npm":
		return initNpmProject(projectName)
	case "go":
		return initGoProject(projectName)
	default:
		return fmt.Errorf("init command not supported for package manager: %s", pm.Name())
	}
}

func initNpmProject(projectName string) error {
	fmt.Println("Creating package.json...")
	
	// Create basic package.json content
	packageJsonContent := fmt.Sprintf(`{
  "name": "%s",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "node index.js",
    "dev": "node index.js"
  },
  "keywords": [],
  "author": "",
  "license": "ISC"
}`, projectName)

	// Write package.json file
	if err := os.WriteFile("package.json", []byte(packageJsonContent), 0644); err != nil {
		return fmt.Errorf("failed to create package.json: %w", err)
	}

	// Create basic index.js file
	indexJsContent := `console.log("Hello from " + require('./package.json').name);`
	if err := os.WriteFile("index.js", []byte(indexJsContent), 0644); err != nil {
		return fmt.Errorf("failed to create index.js: %w", err)
	}

	fmt.Println("✅ npm project initialized successfully!")
	fmt.Println("Files created:")
	fmt.Println("  - package.json")
	fmt.Println("  - index.js")
	
	return nil
}

func initGoProject(projectName string) error {
	fmt.Println("Creating go.mod...")
	
	// Create go.mod content
	goModContent := fmt.Sprintf(`module %s

go 1.21
`, projectName)

	// Write go.mod file
	if err := os.WriteFile("go.mod", []byte(goModContent), 0644); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}

	// Create basic main.go file
	mainGoContent := fmt.Sprintf(`package main

import "fmt"

func main() {
	fmt.Println("Hello from %s!")
}
`, projectName)
	
	if err := os.WriteFile("main.go", []byte(mainGoContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	fmt.Println("✅ Go project initialized successfully!")
	fmt.Println("Files created:")
	fmt.Println("  - go.mod")
	fmt.Println("  - main.go")
	
	return nil
}