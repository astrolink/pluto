package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	env "github.com/astrolink/pluto/internal/env"
	"github.com/astrolink/pluto/internal/storage"

	"github.com/charmbracelet/lipgloss"
)

var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
var green = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7CFC00"))
var orange = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500"))

func Execute(result storage.PlutoXml, file string, cmd string, batch int) {
	if cmd == "" {
		cmd = "run"
	}

	file = storage.RemoveExtensionFromFile(file)

	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	if Check(file) && cmd == "run" {
		_, execErr := db.Exec(result.Run)
		if execErr != nil {
			fmt.Println(red.Render("There was an error running a migration: " + file))
			fmt.Println(red.Render(execErr.Error()))
			os.Exit(1)
		}

		fmt.Println(green.Render("Migration " + file + " executed successfully"))

		Log(file, 1, "Migration executed successfully", result, batch)
	} else {
		if cmd == "run" {
			fmt.Println(red.Render("Migration already executed"))
		}
	}

	db.Close()
}

func Log(file string, success int, message string, xml storage.PlutoXml, batch int) {
	var config string = env.GetMySQlConfig()
	var source string = env.GetSource()
	var author string = strings.Trim(xml.Author, " \n")
	var description string = strings.Trim(xml.Description, " \n")
	file = storage.RemoveExtensionFromFile(file)

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	CreatePlutoTable(db)

	_, Err := db.Exec(
		"INSERT INTO pluto_logs (date, batch, source, file, success, message, author, description) VALUES (NOW(), " + strconv.Itoa(batch) + ", '" + source + "', '" + file + "', " + strconv.Itoa(success) + ", '" + message + "', '" + author + "', '" + description + "');",
	)
	if Err != nil {
		fmt.Println(red.Render(Err.Error()))
		os.Exit(1)
	}

	db.Close()
}

func Check(file string) bool {
	var config string = env.GetMySQlConfig()
	var checked bool = false
	var source string = env.GetSource()
	file = storage.RemoveExtensionFromFile(file)

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	CreatePlutoTable(db)

	var col string
	sqlStatement := "SELECT id FROM pluto_logs WHERE (file = '" + file + "') AND (source = '" + source + "') AND (success = 1) LIMIT 1;"
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

func CheckRollback(file string) bool {
	var config string = env.GetMySQlConfig()
	var checked bool = false
	var source string = env.GetSource()
	file = storage.RemoveExtensionFromFile(file)

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	CreatePlutoTable(db)

	var col string
	sqlStatement := "SELECT id FROM pluto_logs WHERE (file = '" + file + "') AND (source = '" + source + "') AND (success = 1) LIMIT 1;"
	row := db.QueryRow(sqlStatement)
	err2 := row.Scan(&col)
	if err2 == nil {
		checked = true
	}

	db.Close()
	return checked
}

func Rollback(result storage.PlutoXml, file string, step string) {
	var config string = env.GetMySQlConfig()
	var source string = env.GetSource()
	file = storage.RemoveExtensionFromFile(file)

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	_, execErr := db.Exec(result.Rollback)
	if execErr != nil {
		fmt.Println(red.Render("There was an error running a rollback: " + file))
		fmt.Println(red.Render(execErr.Error()))
		os.Exit(1)
	}

	_, delErr := db.Exec("DELETE FROM pluto_logs WHERE (file = '" + file + "') AND (source = '" + source + "');")
	if delErr != nil {
		fmt.Println(red.Render("There was an error deleting rollback: " + file))
		fmt.Println(red.Render(delErr.Error()))
		os.Exit(1)
	}

	db.Close()

	fmt.Println(orange.Render("Rollback " + file + " executed successfully"))

	db.Close()
}

func CreatePlutoTable(db *sql.DB) {
	_, execErr := db.Exec(`
		CREATE TABLE IF NOT EXISTS pluto_logs
			(
				id int(11) NOT NULL AUTO_INCREMENT,
				source varchar(255) NOT NULL,
				batch int(11) NOT NULL DEFAULT 1,
				date datetime NOT NULL,
				file varchar(255) NOT NULL,
				success tinyint(1) NOT NULL DEFAULT 0,
				message varchar(255) NOT NULL,
				author varchar(255),
				description text,
				PRIMARY KEY (id),
				UNIQUE KEY idx_id (id),
				UNIQUE KEY idx_file (file)
			) ENGINE=InnoDB;`,
	)
	if execErr != nil {
		fmt.Println(red.Render(execErr.Error()))
		os.Exit(1)
	}
}

func RecreatePlutoTable() {
	var config string = env.GetMySQlConfig()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	var col string
	sqlStatement := `DROP TABLE pluto_logs;`
	row := db.QueryRow(sqlStatement)
	err2 := row.Scan(&col)
	if err2 == nil {
		log.Fatal(err2)
	}

	CreatePlutoTable(db)

	db.Close()
}

func GetBatch(reverse bool) int {
	var config string = env.GetMySQlConfig()
	var source string = env.GetSource()

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	var total int
	if err := db.QueryRow("SELECT batch FROM pluto_logs WHERE (source = '" + source + "') ORDER BY batch DESC LIMIT 1;").Scan(&total); err != nil {
		if err == sql.ErrNoRows {
			return 1
		}
		return 1
	}

	if reverse {
		total = total + 1
	}

	db.Close()
	return total
}

func TestConnection() bool {
	var config string = env.GetMySQlConfig()
	var checked bool = false

	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)

	var col string
	sqlStatement := "SELECT NOW() FROM DUAL;"
	row := db.QueryRow(sqlStatement)
	err2 := row.Scan(&col)
	if err2 == nil {
		checked = true
		fmt.Println(green.Render("Database connection successful!"))
	} else {
		fmt.Println(red.Render("There was a problem with your connection check the configuration file"))
		log.Fatal(err2)
	}

	db.Close()

	return checked
}
