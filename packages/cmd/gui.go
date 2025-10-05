package cmd

import (
	"github.com/Tatsuyasan/lazyPm/packages/gui"
	"github.com/spf13/cobra"
)

func init() {
	// RootCmd.PersistentFlags().StringP("manager", "m", "", "Force the package manager (e.g., npm, go)")
}

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Launch the GUI",
	Run: func(cmd *cobra.Command, args []string) {
		gui.RunGui()
	},
}
