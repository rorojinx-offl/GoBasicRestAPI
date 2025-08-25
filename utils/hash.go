package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14) //Generate hased password from a byte-casted plaintext using bcrypt with a cost of 14.
	return string(bytes), err                                        //Return the hashed password as a string.
}

func CheckPasswordHash(plaintext, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext)) //Compare a hashed password with its possible plaintext equivalent.
	return err == nil                                                     //Return true if they match, false otherwise.
}
