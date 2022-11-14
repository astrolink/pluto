/*
Copyright © 2022 ROGER SOUZA <rogersilvasouza@hotmail.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"math"
	"strconv"
	"strings"

	storage "github.com/astrolink/pluto/internal/storage"
	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make a new migration",
	Long:  `Create a new migration on folder, add new file with sequence`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var files []fs.FileInfo = storage.ReadFiles()
		var lastId string = "000000"
		var lastIdArray []string
		var newMigration string

		for _, file := range files {
			if !file.IsDir() && strings.Contains(file.Name(), ".xml") {
				lastId = file.Name()
			}
		}

		if strings.Contains(lastId, "_") {
			lastIdArray = strings.Split(lastId, "_")

			newLastId, err := strconv.Atoi(lastIdArray[0])
			if err != nil {
				log.Fatal(err)
			}

			newLastId = newLastId + 1
			var newLastIdInt string = strconv.Itoa(newLastId)
			var newLastIdStr string = StrPad(newLastIdInt, 6, "0", "LEFT")
			newMigration = newLastIdStr + "_" + args[0] + ".xml"
		}

		fmt.Println(green.Render("Migration " + newMigration + " created successfully"))
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

func StrPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	}

	return output
}
