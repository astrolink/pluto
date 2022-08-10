/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"

	storage "pluto/internal/storage"

	env "pluto/internal/env"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migrations",
	Long:  `Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		files := storage.ReadFiles()

		for _, file := range files {
			if !file.IsDir() {
				jsonFile, err := os.Open(storage.Pwd() + "/migrations/" + file.Name())

				if err != nil {
					fmt.Println(err)
				}

				defer jsonFile.Close()

				byteValue, _ := io.ReadAll(jsonFile)

				var result map[string]interface{}
				json.Unmarshal([]byte(byteValue), &result)

				var config string

				switch result["database"] {
				case "mysql":
					config = env.GetMySQlConfig()
				case "postgre":
				default:
					config = env.GetPostgreConfig()
				}

				db, err := sql.Open("mysql", config)
				db.SetConnMaxLifetime(time.Minute * 1)

				if err != nil {
					log.Fatal(err)
				}

				_, execErr := db.Exec(result["run"].(string))
				if execErr != nil {
					log.Fatal(execErr)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
