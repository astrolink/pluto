/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/astrolink/pluto/general/pluto"

	"github.com/c-bata/go-prompt"
)

var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart recreate pluto_logs table, restart project",
	Long:  `This command restarts the entire base very carefully when using this command`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(red.Render("Please confirm that you want to recreate the pluto base, this action is not reversible (Yes/No)"))
		var choosed string = prompt.Input("> ", completer)

		if choosed == "Yes" {
			pluto.ExecuteRestart()
			fmt.Println(red.Render("Logs of pluto is recreated"))
		}
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{
			Text:        "yes",
			Description: "Your base will be recreated and all your migration history will be lost",
		},
		{
			Text:        "no",
			Description: "Nothing will be changed",
		},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
