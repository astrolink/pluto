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

	"gopkg.in/yaml.v3"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
)

type Connection struct {
	Connection struct {
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

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

		// Yaml
		connection := readYml()

		// MySQL
		// db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/api")
		mysql := connection.Connection.Username + ":secret@tcp(127.0.0.1:3306)/api"

		db, err := sql.Open("mysql", mysql)

		if err != nil {
			log.Fatal(err)
		}

		db.QueryRow(result["run"].(string))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func readYml() Connection {
	fileName := "pluto.yml"

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
	}

	var yamlConfig Connection
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	// fmt.Printf("Result: %v\n", yamlConfig.Connection)

	return yamlConfig
}
