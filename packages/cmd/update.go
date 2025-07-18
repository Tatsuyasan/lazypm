package cmd

import (
	"fmt"
	"os/exec"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewUpdateCommand(pmFlag *string) *cobra.Command {
	var all bool
	var dev bool
	
	cmd := createManagerCommand(pmFlag, CommandConfig{
		Use:   "update [package]",
		Short: "Update dependencies using the appropriate package manager",
		Args:  cobra.MaximumNArgs(1),
		RunFunc: func(pm models.PackageManager, args []string) error {
			packageName := ""
			if len(args) > 0 {
				packageName = args[0]
			}

			return updateDependencies(pm, packageName, all, dev)
		},
	})
	
	cmd.Flags().BoolVarP(&all, "all", "a", false, "Update all dependencies")
	cmd.Flags().BoolVarP(&dev, "dev", "d", false, "Include dev dependencies")
	
	return cmd
}

func updateDependencies(pm models.PackageManager, packageName string, all, dev bool) error {
	fmt.Printf("Updating dependencies with %s", pm.Name())
	if packageName != "" {
		fmt.Printf(" (package: %s)", packageName)
	}
	fmt.Println()

	switch pm.Name() {
	case "npm":
		return updateNpmDependencies(packageName, all, dev)
	case "go":
		return updateGoDependencies(packageName, all)
	default:
		return fmt.Errorf("update command not supported for package manager: %s", pm.Name())
	}
}

func updateNpmDependencies(packageName string, all, dev bool) error {
	if packageName != "" {
		// Update specific package
		args := []string{"update", packageName}
		if dev {
			args = append(args, "--save-dev")
		}
		
		cmd := exec.Command("npm", args...)
		cmd.Stdout = nil
		cmd.Stderr = nil
		
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("npm update failed for %s: %w", packageName, err)
		}
		
		fmt.Printf("✅ Updated %s successfully!\n", packageName)
	} else {
		// Update all packages
		cmd := exec.Command("npm", "update")
		cmd.Stdout = nil
		cmd.Stderr = nil
		
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("npm update failed: %w", err)
		}
		
		fmt.Println("✅ All npm dependencies updated successfully!")
	}
	
	return nil
}

func updateGoDependencies(packageName string, all bool) error {
	if packageName != "" {
		// Update specific package
		cmd := exec.Command("go", "get", "-u", packageName)
		cmd.Stdout = nil
		cmd.Stderr = nil
		
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("go get failed for %s: %w", packageName, err)
		}
		
		fmt.Printf("✅ Updated %s successfully!\n", packageName)
	} else {
		// Update all packages
		cmd := exec.Command("go", "get", "-u", "./...")
		cmd.Stdout = nil
		cmd.Stderr = nil
		
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("go get -u failed: %w", err)
		}
		
		// Run go mod tidy to clean up
		tidyCmd := exec.Command("go", "mod", "tidy")
		tidyCmd.Stdout = nil
		tidyCmd.Stderr = nil
		
		if err := tidyCmd.Run(); err != nil {
			fmt.Printf("Warning: go mod tidy failed: %v\n", err)
		}
		
		fmt.Println("✅ All Go dependencies updated successfully!")
	}
	
	return nil
}