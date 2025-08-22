package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14) //Generate hased password from a byte-casted plaintext using bcrypt with a cost of 14.
	return string(bytes), err                                        //Return the hashed password as a string.
}
