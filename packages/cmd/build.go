package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewBuildCommand(pmFlag *string) *cobra.Command {
	return createManagerCommand(pmFlag, CommandConfig{
		Use:   "build [target]",
		Short: "Build the project using the appropriate package manager",
		Args:  cobra.MaximumNArgs(1),
		RunFunc: func(pm models.PackageManager, args []string) error {
			target := ""
			if len(args) > 0 {
				target = args[0]
			}

			return buildProject(pm, target)
		},
	})
}

func buildProject(pm models.PackageManager, target string) error {
	fmt.Printf("Building project with %s", pm.Name())
	if target != "" {
		fmt.Printf(" (target: %s)", target)
	}
	fmt.Println()

	switch pm.Name() {
	case "npm":
		return buildNpmProject(target)
	case "go":
		return buildGoProject(target)
	default:
		return fmt.Errorf("build command not supported for package manager: %s", pm.Name())
	}
}

func buildNpmProject(target string) error {
	// Check if build script exists in package.json
	cmd := exec.Command("npm", "run", "build")
	if target != "" {
		cmd = exec.Command("npm", "run", "build", "--", target)
	}
	
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	if err := cmd.Run(); err != nil {
		// If build script doesn't exist, try webpack or other common build tools
		fmt.Println("No build script found, trying alternative build methods...")
		
		// Try webpack
		webpackCmd := exec.Command("npx", "webpack")
		if err := webpackCmd.Run(); err == nil {
			fmt.Println("✅ Built successfully with webpack!")
			return nil
		}
		
		return fmt.Errorf("build failed: %w", err)
	}
	
	fmt.Println("✅ npm build completed successfully!")
	return nil
}

func buildGoProject(target string) error {
	var cmd *exec.Cmd
	
	if target != "" {
		cmd = exec.Command("go", "build", "-o", target)
	} else {
		cmd = exec.Command("go", "build")
	}
	
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go build failed: %w", err)
	}
	
	fmt.Println("✅ Go build completed successfully!")
	return nil
}