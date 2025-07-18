package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/spf13/cobra"
)

type CommandConfig struct {
	Use     string
	Short   string
	Aliases []string
	Args    cobra.PositionalArgs
	RunFunc func(pm models.PackageManager, args []string) error
}

func createManagerCommand(pmFlag *string, config CommandConfig) *cobra.Command {
	return &cobra.Command{
		Use:     config.Use,
		Short:   config.Short,
		Aliases: config.Aliases,
		Args:    config.Args,
		RunE: func(cmd *cobra.Command, args []string) error {
			return helpers.WithManager(*pmFlag, func(pm models.PackageManager) error {
				return config.RunFunc(pm, args)
			})
		},
	}
}

func PrintManagerAction(pm models.PackageManager, action string) {
	fmt.Printf("%s with %s\n", action, pm.Name())
}

func PrintList(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}