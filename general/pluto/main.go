package pluto

import (
	"io/fs"
	"strings"

	"github.com/astrolink/pluto/internal/database/mysql"
	"github.com/astrolink/pluto/internal/database/postgre"
	"github.com/astrolink/pluto/internal/storage"
)

func RunMigrations() bool {
	ExecuteRun()

	return true
}

func ExecuteRun() {
	var files []fs.FileInfo = storage.ReadFiles()

	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".xml") {
			result := storage.ReadXml(file.Name())

			switch result.Database {
			case "postgre":
				postgre.Execute(result, file.Name(), "run")
			case "mysql":
				mysql.Execute(result, file.Name(), "run")
			default:
				mysql.Execute(result, file.Name(), "run")
			}
		}
	}
}
