package postgre

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	env "github.com/astrolink/pluto/internal/env"
	"github.com/astrolink/pluto/internal/storage"

	"github.com/charmbracelet/lipgloss"
)

func Execute(result storage.PlutoXml, file string, cmd string, batch int) {
	if cmd == "" {
		cmd = "run"
	}

	var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	_, execErr := db.Exec(result.Run)
	if execErr != nil {
		fmt.Println(red.Render("There was an error running a migration: " + file))
		fmt.Println(red.Render(execErr.Error()))
		os.Exit(1)
	}

	Log()
}

func Log() {
	var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	_, execErr := db.Exec("CREATE TABLE IF NOT EXISTS `log` (`id` int(11) NOT NULL AUTO_INCREMENT, `date` datetime NOT NULL, `message` varchar(255) NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB;")
	if execErr != nil {
		fmt.Println(red.Render(execErr.Error()))
		os.Exit(1)
	}
}
