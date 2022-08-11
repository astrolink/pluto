package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	env "pluto/internal/env"

	"github.com/charmbracelet/lipgloss"
)

var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
var green = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7CFC00"))

func Execute(result map[string]interface{}, file string, cmd string) {
	if cmd == "" {
		cmd = "run"
	}

	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	fmt.Println(green.Render("Executing: " + file))

	if Check(file) && cmd == "run" {
		_, execErr := db.Exec(result[cmd].(string))
		if execErr != nil {
			Log(file, 0, "There was an error, please check the log")
			fmt.Println(red.Render("There was an error running a migration: " + file))
			fmt.Println(red.Render(execErr.Error()))
			os.Exit(1)
		}

		Log(file, 1, "Migration executed successfully")
	} else {
		if cmd == "run" {
			fmt.Println(red.Render("Migration already executed"))
		}
	}

	db.Close()
}

func Log(file string, success int, message string) {
	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	_, execErr := db.Exec("CREATE TABLE IF NOT EXISTS `pluto_logs` (`id` int(11) NOT NULL AUTO_INCREMENT, `date` datetime NOT NULL, `file` varchar(255) NOT NULL, `success` tinyint(1) NOT NULL DEFAULT '0', `message` text NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB;")
	if execErr != nil {
		fmt.Println(red.Render(execErr.Error()))
		os.Exit(1)
	}

	_, Err := db.Exec("INSERT INTO `pluto_logs` (`date`, `file`, `success`, `message`) VALUES (NOW(), '" + file + "', " + strconv.Itoa(success) + ", '" + message + "');")
	if Err != nil {
		fmt.Println(red.Render(Err.Error()))
		os.Exit(1)
	}

	db.Close()
}

func Check(file string) bool {
	var config string = env.GetMySQlConfig()
	var checked bool = false

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	var col string
	sqlStatement := "SELECT id FROM `pluto_logs` WHERE (`file` = '" + file + "') AND (success = 1) LIMIT 1;"
	row := db.QueryRow(sqlStatement)
	err2 := row.Scan(&col)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			checked = true
		}
	}

	db.Close()

	return checked
}
