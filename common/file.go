package common

import (
	"fmt"
	"io"
	"log"
	"os"
)

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

	x := os.IsNotExist(err)

	return x
}

func CopyFile(src, dst string) (int64, error) {
	exist := FileExist(dst)
	if !exist {
		return 0, nil
	}

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

func CreateFolder(folderName string) {
	os.Mkdir(folderName, 0755)
}
