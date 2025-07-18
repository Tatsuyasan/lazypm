package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var pmFlag string

	cmd := &cobra.Command{
		Use:   "lpm",
		Short: "A CLI-agnostic wrapper for package managers",
		Long:  "LazyPM is a CLI-agnostic wrapper that provides a unified interface for package managers like npm, go modules, and more.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return helpers.WithManager(pmFlag, func(manager models.PackageManager) error {
				fmt.Printf("Package manager detected: %s\n", manager.Name())
				return nil
			})
		},
	}

	cmd.Flags().StringVarP(&pmFlag, "manager", "m", "", "Force the package manager (e.g., npm, go)")

	cmd.AddCommand(NewInstallCommand(&pmFlag))
	cmd.AddCommand(NewRunCommand(&pmFlag))
	cmd.AddCommand(NewListCommand(&pmFlag))
	cmd.AddCommand(NewInitCommand(&pmFlag))
	cmd.AddCommand(NewBuildCommand(&pmFlag))
	cmd.AddCommand(NewTestCommand(&pmFlag))
	cmd.AddCommand(NewUpdateCommand(&pmFlag))
	cmd.AddCommand(NewCleanCommand(&pmFlag))
	cmd.AddCommand(NewGUICommand())

	return cmd
}
