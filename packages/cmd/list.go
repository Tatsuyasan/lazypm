package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

func NewListCommand(pmFlag *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List scripts or dependencies",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "scripts",
		Short: "List available scripts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return helpers.WithManager(*pmFlag, func(pm models.PackageManager) error {
				scripts, err := pm.ListScripts()
				if err != nil {
					return err
				}
				for _, s := range scripts {
					fmt.Println(s)
				}
				return nil
			})
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "deps",
		Short: "List project dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			return helpers.WithManager(*pmFlag, func(pm models.PackageManager) error {
				deps, err := pm.ListDependencies()
				if err != nil {
					return err
				}
				for _, d := range deps {
					fmt.Println(d)
				}
				return nil
			})
		},
	})

	return cmd
}
