package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewRunCommand(pmFlag *string) *cobra.Command {
	return createManagerCommand(pmFlag, CommandConfig{
		Use:   "run [script]",
		Short: "Run a script using the appropriate package manager",
		Args:  cobra.MinimumNArgs(1),
		RunFunc: func(pm models.PackageManager, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("script name is required")
			}
			
			script := args[0]
			scriptArgs := args[1:]
			
			fmt.Printf("Running script '%s' with %s\n", script, pm.Name())
			return pm.RunScript(script, scriptArgs)
		},
	})
}
