package db

import (
	"database/sql"
	"fmt"

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

	DB.Exec("PRAGMA foreign_keys = ON;")

	createTables()
}

func createTables() {
	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		dob DATETIME NOT NULL,
		weight INTEGER NOT NULL,
		height INTEGER NOT NULL,
		gender TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	createDailyGoalsTableQuery := `
	CREATE TABLE IF NOT EXISTS daily_goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		calories_target INTEGER NOT NULL,
		protein_target INTEGER NOT NULL,
		carbs_target INTEGER NOT NULL,
		fats_target INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	createFoodLogsTableQuery := `
	CREATE TABLE IF NOT EXISTS food_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		calories INTEGER NOT NULL,
		protein INTEGER NOT NULL,
		carbs INTEGER NOT NULL,
		fats INTEGER NOT NULL,
		meal_type TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	queries := []string{createUsersTableQuery, createDailyGoalsTableQuery, createFoodLogsTableQuery}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			panic(fmt.Sprintf("Could not create table; error: %v", err))
		}
	}
}
