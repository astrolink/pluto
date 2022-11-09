/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/astrolink/pluto/internal/database/mysql"
	"github.com/astrolink/pluto/internal/storage"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback migrations",
	Long:  `Long Description`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var files []fs.FileInfo = storage.ReadFiles()

		if (args[0] != "all") && (args[0] != "step=-1") {
			fmt.Println("Invalid argument use all or step=-1")
			os.Exit(1)
		}

		sort.Slice(files, func(i, j int) bool { return files[i].Name() > files[j].Name() })

		for _, file := range files {
			if mysql.CheckRollback(file.Name()) {
				if !file.IsDir() && strings.Contains(file.Name(), ".xml") {
					var result storage.PlutoXml = storage.ReadXml(file.Name())

					switch result.Database {
					case "postgre":
						mysql.Rollback(result, file.Name(), args[0])
					case "mysql":
						mysql.Rollback(result, file.Name(), args[0])
					default:
						mysql.Rollback(result, file.Name(), args[0])
					}

					if args[0] == "step=-1" {
						break
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
