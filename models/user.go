package models

import (
	"example.com/rest_api/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (usr *User) Save() error {
	query := `
	INSERT INTO users(id, email, password) 
	VALUES(?, ?, ?)
	`
	statement, err := db.SqlDatabase.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(usr.ID, usr.Email, usr.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	usr.ID = id

	return err
}
