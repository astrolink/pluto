/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"io/fs"
	"strings"

	"github.com/spf13/cobra"

	"pluto/internal/database/mysql"
	"pluto/internal/database/postgre"
	"pluto/internal/storage"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migrations",
	Long:  `Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		var files []fs.FileInfo = storage.ReadFiles()

		for _, file := range files {
			if !file.IsDir() && strings.Contains(file.Name(), ".json") {
				var result map[string]interface{} = storage.ReadJson(file.Name())

				switch result["database"] {
				case "postgre":
					postgre.Execute(result, file.Name(), "run")
				case "mysql":
					mysql.Execute(result, file.Name(), "run")
				default:
					mysql.Execute(result, file.Name(), "run")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
