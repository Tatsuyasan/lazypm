package cmd

import (
	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i", "add"},
	Short:   "Install project dependencies using the detected or forced package manager",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.GetContext()

		ctx.Manager.Install(args)
	},
}
