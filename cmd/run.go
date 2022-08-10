/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"

	storage "pluto/internal/storage"

	env "pluto/internal/env"

	"github.com/charmbracelet/lipgloss"
)

var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migrations",
	Long:  `Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		files := storage.ReadFiles()

		for _, file := range files {
			if !file.IsDir() {
				var result map[string]interface{} = storage.ReadJson(file.Name())
				var config string

				switch result["database"] {
				case "mysql":
					config = env.GetMySQlConfig()
				case "postgre":
				default:
					config = env.GetPostgreConfig()
				}

				db, err := sql.Open("mysql", config)
				if err != nil {
					log.Fatal(err)
				}

				db.SetConnMaxLifetime(time.Minute * 1)

				_, execErr := db.Exec(result["run"].(string))
				if execErr != nil {
					fmt.Println(red.Render("There was an error running a migration: " + file.Name()))
					fmt.Println(red.Render(execErr.Error()))
					os.Exit(1)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
