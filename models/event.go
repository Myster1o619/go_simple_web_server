package models

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

func (evt *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, date, user_id) 
	VALUES(?, ?, ?, ?, ?)
	`
	statement, err := db.SqlDatabase.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(evt.Name, evt.Description, evt.Location, evt.Date, evt.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	evt.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.SqlDatabase.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// statement, err := db.SqlDatabase.Prepare(query)

	// if err != nil {
	// 	return nil, err
	// }

	// defer statement.Close()

	// _, err = statement.Exec(query)

	// if err != nil {
	// 	return nil, err
	// }

	var events = []Event{}

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description,
			&event.Location, &event.Date, &event.UserID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.SqlDatabase.QueryRow(query, id)

	var event Event

	if err := row.Scan(&event.ID, &event.Name, &event.Description,
		&event.Location, &event.Date, &event.UserID); err != nil {
		return nil, err
	}

	return &event, nil
}
