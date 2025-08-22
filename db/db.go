package db

import (
	"database/sql"
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
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not configure table!")
	}
}
