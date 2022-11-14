/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/astrolink/pluto/general/pluto"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migrations",
	Long:  `Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		pluto.RunMigrations()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
