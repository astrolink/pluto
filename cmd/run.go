/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"io/fs"
	"strings"

	"github.com/spf13/cobra"

	"github.com/astrolink/pluto/general/pluto"
	"github.com/astrolink/pluto/internal/database/mysql"
	"github.com/astrolink/pluto/internal/database/postgre"
	"github.com/astrolink/pluto/internal/storage"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migrations",
	Long:  `Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		executeRun()

		pluto.RunMigrations()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func executeRun() {
	var files []fs.FileInfo = storage.ReadFiles()

	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".xml") {
			result := storage.ReadXml(file.Name())

			switch result.Database {
			case "postgre":
				postgre.Execute(result, file.Name(), "run")
			case "mysql":
				mysql.Execute(result, file.Name(), "run")
			default:
				mysql.Execute(result, file.Name(), "run")
			}
		}
	}
}
