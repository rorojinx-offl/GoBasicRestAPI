package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"` // Binding tag to ensure Name is required
	Description string    `binding:"required"` // Binding tag to ensure Description is required
	Location    string    `binding:"required"` // Binding tag to ensure Location is required
	DateTime    time.Time `binding:"required"` // Binding tag to ensure DateTime is required
	UserID      int
}

var events = []Event{} // Slice to hold all events

func (e *Event) Save() {
	events = append(events, *e) //Append the event to the slice
}

func GetAllEvents() []Event {
	return events // Returns all events stored in the slice
}
