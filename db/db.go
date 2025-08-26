package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" //Import with underscore so that we don't use it directly, but we have SQL lib use its functionalities
)

var DB *sql.DB

// Function to initialize the database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "db/api.db") //Open a database connection to the SQLite database file named api.db
	if err != nil {
		panic("Could not connect to the database")
	}

	DB.SetMaxOpenConns(10) //Configure connection pool to allow a maximum of 25 open connections to the database.
	DB.SetMaxIdleConns(5)  //Set the maximum number of idle connections to 5.

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	    	id INTEGER PRIMARY KEY AUTOINCREMENT,
	    	email TEXT NOT NULL UNIQUE,
	    	password TEXT NOT NULL
	)
`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(fmt.Sprintf("Could not configure users table! %v", err))
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER,
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not configure events table! %v", err))
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    event_id INTEGER,
	    user_id INTEGER,
	    FOREIGN KEY(event_id) REFERENCES events(id),
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not configure registrations table! %v", err))
	}
}
