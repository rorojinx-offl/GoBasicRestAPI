package routes

import (
	"GoBasicRestAPI/models"
	"GoBasicRestAPI/utils"
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

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not authenticate user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Success": "Login successful", "Token": token})
}
