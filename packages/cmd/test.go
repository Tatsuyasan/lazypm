package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewTestCommand(pmFlag *string) *cobra.Command {
	var verbose bool
	var coverage bool
	
	cmd := createManagerCommand(pmFlag, CommandConfig{
		Use:   "test [pattern]",
		Short: "Run tests using the appropriate package manager",
		Args:  cobra.MaximumNArgs(1),
		RunFunc: func(pm models.PackageManager, args []string) error {
			pattern := ""
			if len(args) > 0 {
				pattern = args[0]
			}

			return runTests(pm, pattern, verbose, coverage)
		},
	})
	
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Run tests in verbose mode")
	cmd.Flags().BoolVarP(&coverage, "coverage", "c", false, "Run tests with coverage")
	
	return cmd
}

func runTests(pm models.PackageManager, pattern string, verbose, coverage bool) error {
	fmt.Printf("Running tests with %s", pm.Name())
	if pattern != "" {
		fmt.Printf(" (pattern: %s)", pattern)
	}
	fmt.Println()

	switch pm.Name() {
	case "npm":
		return runNpmTests(pattern, verbose, coverage)
	case "go":
		return runGoTests(pattern, verbose, coverage)
	default:
		return fmt.Errorf("test command not supported for package manager: %s", pm.Name())
	}
}

func runNpmTests(pattern string, verbose, coverage bool) error {
	args := []string{"test"}
	
	if pattern != "" {
		args = append(args, "--", pattern)
	}
	
	if verbose {
		args = append(args, "--verbose")
	}
	
	if coverage {
		args = append(args, "--coverage")
	}
	
	cmd := exec.Command("npm", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("npm test failed: %w", err)
	}
	
	fmt.Println("✅ npm tests completed successfully!")
	return nil
}

func runGoTests(pattern string, verbose, coverage bool) error {
	args := []string{"test"}
	
	if verbose {
		args = append(args, "-v")
	}
	
	if coverage {
		args = append(args, "-cover")
	}
	
	if pattern != "" {
		args = append(args, "-run", pattern)
	}
	
	args = append(args, "./...")
	
	cmd := exec.Command("go", args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go test failed: %w", err)
	}
	
	fmt.Println("✅ Go tests completed successfully!")
	return nil
}