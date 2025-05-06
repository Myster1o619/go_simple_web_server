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
