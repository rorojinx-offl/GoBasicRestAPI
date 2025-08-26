package routes

import (
	"GoBasicRestAPI/utils"
	"github.com/gin-gonic/gin"
)

// RouteRegistrar This function can be used to register routes defined in this package to a Gin router instance.
func RouteRegistrar(server *gin.Engine) {
	//Event routes
	server.GET("/events", getEvents)          // Register the GET request to /events with the getEvents function to retrieve all events.
	server.GET("/events/:id", getSingleEvent) //Using :id can allow for dynamic ID parameters in the URL to get individual events.

	//Protected routes - only authenticated users can create events
	protected := server.Group("/")                           // Create a route group for protected routes
	protected.Use(utils.Authenticate)                        // Apply the authentication middleware to the group
	protected.POST("/events", createEvent)                   // Register the POST request to /events with the createEvent function to create a new event.
	protected.PUT("/events/:id", updateEvent)                // Register the PUT request to /events/:id with the updateEvent function to update an existing event.
	protected.DELETE("/events/:id", deleteEvent)             // Register the DELETE request to /events/:id with the deleteEvent function to delete an event.
	protected.POST("/events/:id/register", registerForEvent) //Register the POST request to /events/:id/register with the registerForEvent function to register the authenticated user for the specified event.
	protected.DELETE("/events/:id/register", cancelRegistration)

	//User routes
	server.POST("/signup", signup) // Register the POST request to /signup with the signup function to create a new user.
	server.POST("/login", login)   // Register the POST request to /login with the login function to authenticate a user.
}
