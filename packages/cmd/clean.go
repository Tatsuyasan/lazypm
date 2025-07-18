package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewCleanCommand(pmFlag *string) *cobra.Command {
	var cache bool
	var force bool
	
	cmd := createManagerCommand(pmFlag, CommandConfig{
		Use:   "clean",
		Short: "Clean build artifacts and cache using the appropriate package manager",
		Args:  cobra.NoArgs,
		RunFunc: func(pm models.PackageManager, args []string) error {
			return cleanProject(pm, cache, force)
		},
	})
	
	cmd.Flags().BoolVarP(&cache, "cache", "c", false, "Also clean cache")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force clean without confirmation")
	
	return cmd
}

func cleanProject(pm models.PackageManager, cache, force bool) error {
	fmt.Printf("Cleaning project with %s", pm.Name())
	if cache {
		fmt.Print(" (including cache)")
	}
	fmt.Println()

	switch pm.Name() {
	case "npm":
		return cleanNpmProject(cache, force)
	case "go":
		return cleanGoProject(cache, force)
	default:
		return fmt.Errorf("clean command not supported for package manager: %s", pm.Name())
	}
}

func cleanNpmProject(cache, force bool) error {
	cleaned := []string{}
	
	// Remove node_modules
	if _, err := os.Stat("node_modules"); err == nil {
		shouldRemove := force
		if !force {
			fmt.Print("Remove node_modules? (y/N): ")
			var response string
			fmt.Scanln(&response)
			shouldRemove = response == "y" || response == "Y"
		}
		
		if shouldRemove {
			if err := os.RemoveAll("node_modules"); err != nil {
				return fmt.Errorf("failed to remove node_modules: %w", err)
			}
			cleaned = append(cleaned, "node_modules")
		}
	}
	
	// Remove package-lock.json
	if _, err := os.Stat("package-lock.json"); err == nil {
		if err := os.Remove("package-lock.json"); err != nil {
			return fmt.Errorf("failed to remove package-lock.json: %w", err)
		}
		cleaned = append(cleaned, "package-lock.json")
	}
	
	// Remove common build directories
	buildDirs := []string{"dist", "build", "out", ".next", ".nuxt"}
	for _, dir := range buildDirs {
		if _, err := os.Stat(dir); err == nil {
			if err := os.RemoveAll(dir); err != nil {
				fmt.Printf("Warning: failed to remove %s: %v\n", dir, err)
			} else {
				cleaned = append(cleaned, dir)
			}
		}
	}
	
	if cache {
		// Clean npm cache
		cmd := exec.Command("npm", "cache", "clean", "--force")
		cmd.Stdout = nil
		cmd.Stderr = nil
		
		if err := cmd.Run(); err != nil {
			fmt.Printf("Warning: npm cache clean failed: %v\n", err)
		} else {
			cleaned = append(cleaned, "npm cache")
		}
	}
	
	if len(cleaned) > 0 {
		fmt.Println("✅ Cleaned successfully!")
		fmt.Println("Removed:")
		for _, item := range cleaned {
			fmt.Printf("  - %s\n", item)
		}
	} else {
		fmt.Println("Nothing to clean")
	}
	
	return nil
}

func cleanGoProject(cache, force bool) error {
	cleaned := []string{}
	
	// Clean go build cache
	cmd := exec.Command("go", "clean")
	cmd.Stdout = nil
	cmd.Stderr = nil
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go clean failed: %w", err)
	}
	cleaned = append(cleaned, "build cache")
	
	// Clean go mod cache if requested
	if cache {
		modCacheCmd := exec.Command("go", "clean", "-modcache")
		modCacheCmd.Stdout = nil
		modCacheCmd.Stderr = nil
		
		if err := modCacheCmd.Run(); err != nil {
			fmt.Printf("Warning: go clean -modcache failed: %v\n", err)
		} else {
			cleaned = append(cleaned, "module cache")
		}
	}
	
	// Remove common build artifacts
	buildFiles := []string{"main", "app", "cmd", "*.exe"}
	for _, pattern := range buildFiles {
		// This is a simplified approach - in a real implementation,
		// you'd want to use filepath.Glob to match patterns
		if _, err := os.Stat(pattern); err == nil {
			if err := os.Remove(pattern); err != nil {
				fmt.Printf("Warning: failed to remove %s: %v\n", pattern, err)
			} else {
				cleaned = append(cleaned, pattern)
			}
		}
	}
	
	fmt.Println("✅ Cleaned successfully!")
	fmt.Println("Removed:")
	for _, item := range cleaned {
		fmt.Printf("  - %s\n", item)
	}
	
	return nil
}