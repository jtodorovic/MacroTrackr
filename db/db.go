package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to the database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		dob DATETIME NOT NULL,
		weight INTEGER NOT NULL,
		gender TEXT NOT NULL
	)
	`
	createGoalsTable := `
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		calories_target INTEGER NOT NULL,
		protein_target INTEGER NOT NULL,
		carbs_target INTEGER NOT NULL,
		fats_target INTEGER NOT NULL,
		user_id INTEGER NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create Users table.")
	}

	_, err2 := DB.Exec(createGoalsTable)
	if err2 != nil {
		panic("Could not create Goals table.")
	}
}
