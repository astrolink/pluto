/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files:`,
	Run: func(cmd *cobra.Command, args []string) {
		if FileExist("pluto.yaml") {
			CreateFile("pluto.yaml")
		}

		if FileExist(".migration") {
			CreateFile(".migration")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created file: %s\n", fileName)
	defer file.Close()
}

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	fmt.Println("err:", err)

	x := os.IsNotExist(err)
	fmt.Println("x:", x)

	return x
}
