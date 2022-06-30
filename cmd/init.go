/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start pluto in a new project",
	Long:  `Init create a new project with the following files:`,
	Run: func(cmd *cobra.Command, args []string) {

		// CreateFile("pluto.yml")

		CopyFile("docs/examples/base.yml", "pluto.yml")

		// CreateFile("migrations/001_users.json")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func CreateFile(fileName string) {
	if FileExist(fileName) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created file: %s\n", fileName)
		defer file.Close()
	}
}

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	fmt.Println("err:", err)

	x := os.IsNotExist(err)
	fmt.Println("x:", x)

	return x
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
