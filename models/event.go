package models

// data logic to go here

import (
	"time"
)

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	Date        time.Time
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
