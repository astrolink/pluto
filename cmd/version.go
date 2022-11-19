package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var greenLite = lipgloss.NewStyle().Foreground(lipgloss.Color("#7CFC00"))

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the pluto project version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(greenLite.Render("Version: v0.1.14 \nPluto is installed correctly. Enjoy! 😎"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
