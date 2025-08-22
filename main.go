package main

import (
	"GoBasicRestAPI/db"
	"GoBasicRestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)    // Register the GET request to /events with the getEvents function to retrieve all events.
	server.POST("/events", createEvent) // Register the POST request to /events with the createEvent function to create a new event.

	server.Run(":8081") //localhost:8081

}

// Function will handle the GET request to /events and takes a context parameter which handles the request and response.
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()     //Returns a slice of all events
	context.JSON(http.StatusOK, events) //Returns a JSON with OK status (code 200) and a slice of events.
}

func createEvent(context *gin.Context) {
	event := models.Event{}

	err := context.ShouldBind(&event) //ShouldBind will bind the request body to the format of the Event struct.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse request"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save() // Save the event using the Save method to add struct to the slice of events.

	context.JSON(http.StatusCreated, gin.H{"Success": "Event created successfully", "Event": event})
}
