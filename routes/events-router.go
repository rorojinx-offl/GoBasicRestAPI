package routes

import (
	"GoBasicRestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Function will handle the GET request to /events and takes a context parameter which handles the request and response.
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents() //Returns a slice of all events
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not retrieve events"})
		return
	}

	context.JSON(http.StatusOK, events) //Returns a JSON with OK status (code 200) and a slice of events.
}

func getSingleEvent(context *gin.Context) {
	idString := context.Param("id")                    //Retrieve the ID parameter from the URL.
	eventId, err := strconv.ParseInt(idString, 10, 64) //Convert the ID string to an integer.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not retrieve event"})
		return
	}

	context.JSON(http.StatusOK, event)
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

	err = event.Save() // Save the event using the Save method to add struct to the slice of events.
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not create events"})
	}

	context.JSON(http.StatusCreated, gin.H{"Success": "Event created successfully", "Event": event})
}

func updateEvent(context *gin.Context) {
	idString := context.Param("id")                    //Retrieve the ID parameter from the URL.
	eventId, err := strconv.ParseInt(idString, 10, 64) //Convert the ID string to an integer.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse event ID"})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not retrieve event"})
		return
	}

	updatedEvent := models.Event{}
	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse request"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not update event"})
	}

	context.JSON(http.StatusOK, gin.H{"Success": "Event updated successfully", "Event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	idString := context.Param("id")                    //Retrieve the ID parameter from the URL.
	eventId, err := strconv.ParseInt(idString, 10, 64) //Convert the ID string to an integer.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not retrieve event"})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Success": "Event deleted successfully"})
}
