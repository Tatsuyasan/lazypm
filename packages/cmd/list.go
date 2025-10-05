package cmd

//
// import (
// 	"github.com/Tatsuyasan/lazyPm/packages/models"
// 	"github.com/spf13/cobra"
// )
//
// func NewListCommand(pmFlag *string) *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "list",
// 		Short: "List scripts or dependencies",
// 	}
//
// 	scriptsCmd := createManagerCommand(pmFlag, CommandConfig{
// 		Use:   "scripts",
// 		Short: "List available scripts",
// 		Args:  cobra.NoArgs,
// 		RunFunc: func(pm models.PackageManager, args []string) error {
// 			scripts, err := pm.ListScripts()
// 			if err != nil {
// 				return err
// 			}
// 			PrintList(scripts)
// 			return nil
// 		},
// 	})
//
// 	depsCmd := createManagerCommand(pmFlag, CommandConfig{
// 		Use:   "deps",
// 		Short: "List project dependencies",
// 		Args:  cobra.NoArgs,
// 		RunFunc: func(pm models.PackageManager, args []string) error {
// 			deps, err := pm.ListDependencies()
// 			if err != nil {
// 				return err
// 			}
// 			PrintList(deps)
// 			return nil
// 		},
// 	})
//
// 	cmd.AddCommand(scriptsCmd)
// 	cmd.AddCommand(depsCmd)
//
// 	return cmd
// }
