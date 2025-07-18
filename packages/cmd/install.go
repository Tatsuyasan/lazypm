package cmd

import (
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewInstallCommand(pmFlag *string) *cobra.Command {
	return createManagerCommand(pmFlag, CommandConfig{
		Use:     "install",
		Aliases: []string{"i", "add"},
		Short:   "Install project dependencies using the detected or forced package manager",
		Args:    cobra.ArbitraryArgs,
		RunFunc: func(pm models.PackageManager, args []string) error {
			PrintManagerAction(pm, "Installing dependencies")
			return pm.Install(args)
		},
	})
}
