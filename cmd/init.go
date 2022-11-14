/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"fmt"

	storage "github.com/astrolink/pluto/internal/storage"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files 000001_created_users.xml`,
	Run: func(cmd *cobra.Command, args []string) {
		storage.CreatePlutoFile()

		storage.CreateFolder("migrations")

		storage.CreateMigrationXmlFile()

		fmt.Println("Configuration file created and migrations folder started 🎉")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
