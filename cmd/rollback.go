/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"

	"github.com/astrolink/pluto/general/pluto"
	"github.com/astrolink/pluto/internal/storage"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback migrations",
	Long:  `When performing rollback, migrations are rolled back`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var files []fs.DirEntry = storage.ReadFiles()

		if (args[0] != "all") && (args[0] != "step=-1") {
			fmt.Println("Invalid argument use all or step=-1")
			os.Exit(1)
		}

		pluto.RunRollback(files, args)
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
