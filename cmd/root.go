package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/internal/helpers"
	"github.com/Tatsuyasan/lazyPm/internal/models"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var pmFlag string

	cmd := &cobra.Command{
		Use:   "lpm",
		Short: "A CLI-agnostic wrapper for package managers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return helpers.WithManager(pmFlag, func(manager models.PackageManager) error {
				fmt.Println("package manager detected :", manager.Name())
				return nil
			})
		},
	}

	// flags for command root
	cmd.Flags().StringVarP(&pmFlag, "manager", "m", "", "Force the package manager (e.g., npm, go)")

	// add subcommands here to root command
	cmd.AddCommand(NewInstallCommand(&pmFlag))
	cmd.AddCommand(NewRunCommand(&pmFlag))
	cmd.AddCommand(NewListCommand(&pmFlag))

	return cmd
}
