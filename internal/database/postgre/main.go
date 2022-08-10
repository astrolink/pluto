package postgre

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	env "pluto/internal/env"

	"github.com/charmbracelet/lipgloss"
)

func Execute(result map[string]interface{}, file string) {
	var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	_, execErr := db.Exec(result["run"].(string))
	if execErr != nil {
		fmt.Println(red.Render("There was an error running a migration: " + file))
		fmt.Println(red.Render(execErr.Error()))
		os.Exit(1)
	}
}
