/*
Copyright © 2022  Phelipe Galiotti
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(green.Render("Version: v0.1.19"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
