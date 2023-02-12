package database

import (
	"time"
)

// Activity represents a record of an active application
type Activity struct {
	ClassName string
	Title     string
	Timestamp time.Time
}
