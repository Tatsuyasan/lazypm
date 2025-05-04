package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/helpers"
	"github.com/Tatsuyasan/lazyPm/internal/pkgman"
	"github.com/spf13/cobra"
)

func NewRunCommand(pmFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [script]",
		Short: "Run a script using the appropriate package manager",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			script := args[0]
			scriptArgs := args[1:]

			return helpers.WithManager(*pmFlag, func(pm pkgman.PackageManager) error {
				fmt.Printf("Running script '%s' with %s\n", script, pm.Name())
				return pm.RunScript(script, scriptArgs)
			})
		},
	}
	return cmd
}
