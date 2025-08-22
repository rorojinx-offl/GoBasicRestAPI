package models

import (
	"GoBasicRestAPI/db"
	"time"
)

type Event struct {
	ID          int64     //Int64 to match the type returned by LastInsertId
	Name        string    `binding:"required"` // Binding tag to ensure Name is required and mapped to the JSON body request
	Description string    `binding:"required"` // Binding tag to ensure Description is required and mapped to the JSON body request
	Location    string    `binding:"required"` // Binding tag to ensure Location is required and mapped to the JSON body request
	DateTime    time.Time `binding:"required"` // Binding tag to ensure DateTime is required and mapped to the JSON body request
	UserID      int
}

var events = []Event{} // Slice to hold all events

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description,location,dateTime,user_id) 
	VALUES (?,?,?,?,?)` // SQL query to insert a new event into the events table and placeholders as values to be inserted to avoid SQLi attacks
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() //Ensure the statement is closed only after execution of function to free up resources

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id // Set the ID of the event to the last inserted ID from the database
	return err

	//events = append(events, *e) //Append the event to the slice
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events" // SQL query to select all events from the events table
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil // Returns all events stored in the slice
}
