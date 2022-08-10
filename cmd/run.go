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
	"time"

	"gopkg.in/yaml.v3"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"

	file "pluto/common"
)

type Connection struct {
	Connection struct {
		Drive    string `yaml:"drive"`
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
		files := file.ReadFiles()

		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name(), file.IsDir())

				pwd, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				jsonFile, err := os.Open(pwd + "/migrations/" + file.Name())

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
				mysql := connection.Connection.Username + ":@tcp(127.0.0.1:3306)/astrolink"

				db, err := sql.Open("mysql", mysql)
				db.SetConnMaxLifetime(time.Minute * 1)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("SQL: %v\n", result["run"])

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

	fmt.Printf("Result: %v\n", yamlConfig.Connection.Drive)

	return yamlConfig
}
