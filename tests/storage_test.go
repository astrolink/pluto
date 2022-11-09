package main

import (
	"log"
	"os"
	"github.com/astrolink/pluto/internal/storage"

	"testing"
)

func TestStorage(t *testing.T) {
	t.Run("Verify file is created", func(t *testing.T) {
		var fileName string = "test.txt"

		os.Remove(fileName)

		storage.CreateFile(fileName)

		_, err := os.Stat(fileName)
		var notExists bool = os.IsNotExist(err)
		if notExists {
			t.Errorf("File %s not created", fileName)
		}

		var error error = os.Remove(fileName)
		if error != nil {
			log.Fatal(error)
		}
	})

	t.Run("Check for existence of the file inside the folder", func(t *testing.T) {
		var filePath string = "tests"
		var fileName string = "test.txt"

		os.Remove(fileName)

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if !storage.FileExist(filePath + "/" + fileName) {
			t.Errorf("File %s dont exists", filePath+"/"+fileName)
		}

		var error error = os.Remove(fileName)
		if error != nil {
			log.Fatal(error)
		}
	})
}
