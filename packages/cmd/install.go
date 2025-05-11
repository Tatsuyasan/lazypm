package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewInstallCommand(pmFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install",
		Aliases: []string{"i", "add"},
		Short:   "Install project dependencies using the detected or forced package manager",
		Args:    cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return helpers.WithManager(*pmFlag, func(pm models.PackageManager) error {
				fmt.Println("Installing dependencies with", pm.Name())
				return pm.Install(args)
			})
		},
	}
	return cmd
}
