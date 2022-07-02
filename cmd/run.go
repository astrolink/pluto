/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>

*/
package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run migrations",
	Long:  `Long Description`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jsonFile, err := os.Open(pwd + "/migrations/000001_create_users_table.json")

		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		// MySQL
		db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/api")

		if err != nil {
			log.Fatal(err)
		}

		db.QueryRow(result["run"].(string))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
