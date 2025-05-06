package models

// data logic to go here

import (
	"time"

	"example.com/rest_api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (evt *Event) Save() error {
	query := `
	INSERT INTO TABLE events(name, description, location, date, user_id) 
	VALUES(?, ?, ?, ?, ?)
	`
	statement, err := db.SqlDatabase.Prepare(query)

	if err != nil {
		return err
	}

	result, err := statement.Exec(evt.Name, evt.Description, evt.Location, evt.Date, evt.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	evt.ID = id
	
	return nil
}

func GetAllEvents() []Event {
	return events
}
