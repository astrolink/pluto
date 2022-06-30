/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>

*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files:`,
	Run: func(cmd *cobra.Command, args []string) {
		createFile()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createFile() {
	file, err := os.Create("pluto.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
