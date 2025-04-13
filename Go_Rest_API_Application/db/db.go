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
		panic("could not connect to DB")
	}

	DB.SetMaxOpenConns(10) // you can open 10 DB connection at once
	DB.SetMaxIdleConns(5)  // how many connections need to open when no one is using

	createTables()

}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   email TEXT NOT NULL UNIQUE,
	   password TEXT NOT NULL
	   )
	   `
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create user table")
	}

	createEventsTable := `
	
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)
		`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create event table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  event_id INTEGER,
	  user_id INTEGER,
	  FOREIGN KEY(event_id) REFERENCES events(id),
	  FOREIGN KEY(user_id) REFERENCES users(id)
    )
	`
	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("could not create registration table")
	}
}
