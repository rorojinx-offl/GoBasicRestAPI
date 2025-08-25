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

func login(context *gin.Context) {
	user := models.User{}

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Could not parse request"})
	}

	err = user.ValidateCreds()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid credentials"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Success": "Login successful"})
}
