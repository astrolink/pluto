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
		storage.CreatePlutoFile()

		storage.CreateFolder("migrations")

		storage.CreateMigrationXmlFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
