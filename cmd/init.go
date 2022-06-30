/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>

*/
package cmd

import (
	"pluto/common"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files:`,
	Run: func(cmd *cobra.Command, args []string) {

		// CreateFile("pluto.yml")

		common.CopyFile("docs/examples/base.yml", "pluto.yml")

		// CreateFile("migrations/001_users.json")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
