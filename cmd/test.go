/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/astrolink/pluto/general/pluto"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Verify pluto is connected on database",
	Long:  `Verify pluto is connected on database`,
	Run: func(cmd *cobra.Command, args []string) {
		pluto.TestConnection()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
