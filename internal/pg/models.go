// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package pg

import (
	"time"
)

type Activity struct {
	ID        int32
	ClassName string
	Title     string
	Timestamp time.Time
}
