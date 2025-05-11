package cmd

import (
	"github.com/Tatsuyasan/lazyPm/packages/gui"
	"github.com/spf13/cobra"
)

func NewGUICommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gui",
		Short: "Launch the GUI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return gui.RunGUI()
		},
	}
}
