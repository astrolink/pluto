package storage

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func ReadFiles() []fs.FileInfo {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(pwd + "/migrations/")
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func Pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd
}

func ReadJson(name string) map[string]interface{} {
	jsonFile, err := os.Open(Pwd() + "/migrations/" + name)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

type PlutoXml struct {
	XMLName     xml.Name `xml:"pluto"`
	Database    string   `xml:"database"`
	Run         string   `xml:"run"`
	Rollback    string   `xml:"rollback"`
	Description string   `xml:"description"`
}

func ReadXml(name string) PlutoXml {
	xmlFile, err := os.Open(Pwd() + "/migrations/" + name)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var result PlutoXml
	xml.Unmarshal(byteValue, &result)

	return result
}

func CreatePlutoFile() {
	exist := FileExist("pluto.yml")
	if !exist {
		fmt.Println(exist)
		os.Exit(1)
	}

	file := []byte("mysql:\n" +
		"  host: \"127.0.0.1\"\n" +
		"  port: 3306\n" +
		"  database: \"api\"\n" +
		"  username: \"root\"\n" +
		"  password: \"secret\"\n" +
		"\n" +
		"postgre:\n" +
		"  host: \"127.0.0.1\"\n" +
		"  port: 5432\n" +
		"  database: \"base\"\n" +
		"  username: \"postgres\"\n" +
		"  password: \"\"\n" +
		"\n" +
		"log: \"mysql\"\n" +
		"source: \"api\"")
	err := os.WriteFile("pluto.yml", file, 0644)
	if err != nil {
		panic(err)
	}
}

func CreateMigrationXmlFile() {
	exist := FileExist("migrations/000001_create_users_table.xml")
	if !exist {
		fmt.Println(exist)
		os.Exit(1)
	}

	file := []byte("<pluto>\n" +
		"    <database>\n" +
		"        mysql\n" +
		"    </database>\n" +
		"    <run>\n" +
		"        CREATE TABLE users (name VARCHAR(20),email VARCHAR(20),created_at DATE);\n" +
		"    </run>\n" +
		"    <rollback>\n" +
		"        DROP TABLE users;\n" +
		"    </rollback>\n" +
		"<pluto>\n")
	err := os.WriteFile("migrations/000001_create_users_table.xml", file, 0644)
	if err != nil {
		panic(err)
	}
}

func RemoveExtensionFromFile(file string) string {
	file = strings.Replace(file, ".json", "", 1)
	file = strings.Replace(file, ".xml", "", 1)
	return file
}
