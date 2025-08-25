package models

import (
	"GoBasicRestAPI/db"
	"GoBasicRestAPI/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	u.ID = userID
	return err
}

func (u *User) ValidateCreds() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordValid {
		return errors.New("invalid credentials") // Return an error if the credentials are invalid.
	}
	return nil
}
