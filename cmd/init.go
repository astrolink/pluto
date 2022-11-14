/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"fmt"

	storage "github.com/astrolink/pluto/internal/storage"
	"github.com/charmbracelet/lipgloss"

	"github.com/spf13/cobra"
)

var green = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7CFC00"))

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files 000001_created_users.xml`,
	Run: func(cmd *cobra.Command, args []string) {
		storage.CreatePlutoFile()

		storage.CreateFolder("migrations")

		storage.CreateMigrationXmlFile()

		fmt.Println(green.Render("Configuration file created and migrations folder started 🎉"))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
