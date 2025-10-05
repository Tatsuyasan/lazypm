package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/spf13/cobra"
)

func init() {
	// RootCmd.PersistentFlags().StringP("manager", "m", "", "Force the package manager (e.g., npm, go)")
}

var scriptsCmd = &cobra.Command{
	Use:   "scripts",
	Short: "List available scripts",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.GetContext()

		scripts, _ := ctx.Manager.ListScripts()

		fmt.Println(scripts) // Affiche le slice tel quel
	},
}
