package routes

import (
	"GoBasicRestAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	user := models.User{}

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse request"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Success": "User created successfully"})
}
