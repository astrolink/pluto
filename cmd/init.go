/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	file "pluto/common"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files:`,
	Run: func(cmd *cobra.Command, args []string) {
		file.CopyFile("docs/examples/base.yml", "pluto.yml")

		file.CreateFolder("migrations")

		file.CopyFile("docs/examples/000001_create_users_table.json", "migrations/000001_create_users_table.json")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
