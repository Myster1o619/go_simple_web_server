package models

// data logic to go here

import (
	"time"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (evt *Event) Save() {
	// later: save to database
	events = append(events, *evt)
}

func GetAllEvents() []Event {
	return events
}
