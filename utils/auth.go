package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	//Extract the token from the Authorization header
	token := context.Request.Header.Get("Authorization") //Authorization header is used to pass JWT token for authentication.
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "No token found"})
		return
	}

	//Verify the token and extract the user ID
	userId, err := VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
		return
	}

	//Store the user ID in the context for use in later handlers
	context.Set("userId", userId)

	context.Next() //Call the next handler in the chain (e.g. the actual route handlers).
}
