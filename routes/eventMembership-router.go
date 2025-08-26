package routes

import (
	"GoBasicRestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")                          // Retrieve the authenticated user's ID from the context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Parse the event ID from the URL parameter
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not retrieve event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Success": "User registered for event successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")                          // Retrieve the authenticated user's ID from the context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Parse the event ID from the URL parameter

	event := models.Event{}
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not cancel registration for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Success": "Cancelled registration for event successfully"})
}
