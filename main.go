package main

import (
	"GoBasicRestAPI/db"
	"GoBasicRestAPI/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RouteRegistrar(server) //Connecting the routes defined in the routes package to the Gin server instance.

	err := server.Run(":8081") //localhost:8081
	if err != nil {
		panic(fmt.Sprintf("Failed to bootsrap the server: %v", err))
	}
}
