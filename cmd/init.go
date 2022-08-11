/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	storage "pluto/internal/storage"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files:`,
	Run: func(cmd *cobra.Command, args []string) {
		storage.CopyFile("docs/examples/base.yml", "pluto.yml")

		storage.CreateFolder("migrations")

		storage.CopyFile("docs/examples/000001_create_users_table.json", "migrations/000001_create_users_table.json")

		storage.CopyFile("docs/examples/000002_create_settings_table.json", "migrations/000002_create_settings_table.json")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
