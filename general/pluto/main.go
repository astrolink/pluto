package pluto

import (
	"io/fs"
	"sort"
	"strings"

	"github.com/astrolink/pluto/internal/database/mysql"
	"github.com/astrolink/pluto/internal/database/postgre"
	"github.com/astrolink/pluto/internal/storage"
)

func RunMigrations() bool {
	ExecuteRun()

	return true
}

func RunRollback(files []fs.DirEntry, args []string) bool {
	ExecuteRollback(files, args)

	return true
}

func ExecuteRun() {
	var files []fs.DirEntry = storage.ReadFiles()

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

func ExecuteRollback(files []fs.DirEntry, args []string) {
	sort.Slice(files, func(i, j int) bool { return files[i].Name() > files[j].Name() })

	for _, file := range files {
		if mysql.CheckRollback(file.Name()) {
			if !file.IsDir() && strings.Contains(file.Name(), ".xml") {
				var result storage.PlutoXml = storage.ReadXml(file.Name())

				switch result.Database {
				case "postgre":
					mysql.Rollback(result, file.Name(), args[0])
				case "mysql":
					mysql.Rollback(result, file.Name(), args[0])
				default:
					mysql.Rollback(result, file.Name(), args[0])
				}

				if args[0] == "step=-1" {
					break
				}
			}
		}
	}
}

func ExecuteRestart() {
	mysql.RecreatePlutoTable()
}
