package routes

import "github.com/gin-gonic/gin"

// RouteRegistrar This function can be used to register routes defined in this package to a Gin router instance.
func RouteRegistrar(server *gin.Engine) {
	server.GET("/events", getEvents)          // Register the GET request to /events with the getEvents function to retrieve all events.
	server.GET("/events/:id", getSingleEvent) //Using :id can allow for dynamic ID parameters in the URL to get individual events.
	server.POST("/events", createEvent)       // Register the POST request to /events with the createEvent function to create a new event.
	server.PUT("/events/:id", updateEvent)    // Register the PUT request to /events/:id with the updateEvent function to update an existing event.
	server.DELETE("/events/:id", deleteEvent) // Register the DELETE request to /events/:id with the deleteEvent function to delete an event.
}
