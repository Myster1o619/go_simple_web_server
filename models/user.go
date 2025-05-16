package models

import (
	"example.com/rest_api/db"
	"example.com/rest_api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (usr *User) Save() error {
	query := `
	INSERT INTO users(email, password) 
	VALUES(?, ?)
	`
	statement, err := db.SqlDatabase.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utils.HashPassword(usr.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(usr.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	usr.ID = id

	return err
}
