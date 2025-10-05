// Package cmd defines the CLI commands for lazyPm using Cobra.
//
// This package contains all the commands, flags, and configuration for
// the lazyPm command-line interface. Each file typically defines one
// command or a related set of commands.
package cmd

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.PersistentFlags().StringP("manager", "m", "", "Force the package manager (e.g., npm, go)")

	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(scriptsCmd)
	RootCmd.AddCommand(guiCmd)
}

var RootCmd = &cobra.Command{
	Use:   "lpm",
	Short: "A CLI-agnostic wrapper for package managers",
	Long:  "LazyPM is a CLI-agnostic wrapper that provides a unified interface for package managers like npm, go modules, and more.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		managerNameFlag, _ := cmd.Flags().GetString("manager")

		ctx := context.GetContext()
		pm := helpers.PackageManagerName(managerNameFlag)

		ctx.Manager, _ = helpers.GetPackageManager(pm)
		fmt.Printf("Package manager: %s\n", ctx.Manager.Name())
	},
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Root command executed")
	},
}

// juste make a wrapper that call the package manager with the args ?

// var RootCmd = &cobra.Command{
// 	Use:   "lpm",
// 	Short: "A CLI-agnostic wrapper for package managers",
// 	Long:  "LazyPM is a CLI-agnostic wrapper that provides a unified interface for package managers like npm, go modules, and more.",
// 	PersistentPreRun: func(cmd *cobra.Command, args []string) {
// 		managerNameFlag, _ := cmd.Flags().GetString("manager")
//
// 		ctx := context.GetContext()
// 		pm := helpers.PackageManagerName(managerNameFlag)
//
// 		ctx.Manager, _ = helpers.GetPackageManager(pm)
// 		fmt.Printf("Package manager: %s\n", ctx.Manager.Name())
// 	},
// 	Run: func(cmd *cobra.Command, args []string) {
// 		ctx := context.GetContext()
//
// 		log.Println("Executing root command with args:", args)
// 		cmdt := exec.Command(ctx.Manager.Name(), args...)
// 		cmdt.Dir = "."
// 		cmdt.Stdout = os.Stdout
// 		cmdt.Stderr = os.Stderr
// 		cmdt.Run()
//
// 		// fmt.Println("Root command executed")
// 	},
// }
