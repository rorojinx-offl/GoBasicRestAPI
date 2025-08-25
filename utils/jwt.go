package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "deeznutz" //Secret key to sign the token. In a real application, this should be stored securely and not hardcoded.

func GenerateToken(email string, userID int64) (string, error) {
	//Claims means the payload of the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), //2 hours expiration time in the Unix time format
	}) //Using HS256 algorithm to sign the token and MapClaims to create a map of claims.

	return token.SignedString([]byte(secretKey)) //The secret key is converted to a byte slice before signing the token to avoid any type mismatch issues.
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //Type check to ensure the signing method is HMAC
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) //Type check to ensure the claims are of type MapClaims
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	//email := claims["email"].(string)  //Accessing the email claim from the token with a type assertion
	userId := int64(claims["userId"].(float64)) //Accessing the userId claim from the token with a type assertion

	return userId, nil
}
