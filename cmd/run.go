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
	Short: "To run all your pending migrations",
	Long:  `When executing the run, all migrations that have not yet been executed will be executed`,
	Run: func(cmd *cobra.Command, args []string) {
		pluto.RunMigrations()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
