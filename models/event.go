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
	e.ID = id  // Set the ID of the event to the last inserted ID from the database
	return err //The latest error encountered, if any, during the execution of the function

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

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id) // QueryRow is used when we expect only one row to be returned and the second argument is the value to replace the placeholder in the query.

	event := Event{}
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err //Nil can be returned for a struct as the return type is a pointer to the struct, not a copy of the struct. To return a copy, use Event
	}

	return &event, nil
}

func (e *Event) UpdateEvent() error {
	query := `
	UPDATE events SET name=?, description=?, location=?, dateTime=?
	WHERE id=?
`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}
