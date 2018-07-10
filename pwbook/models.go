package pwbook

import (
	"time"
)

// Entry holds information about name, password pair
type Entry struct {
	Key        string `storm:"id"`
	Value      string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// ModifiedSince returns days elapsed since modifiedAt
func (e *Entry) ModifiedSince() int {
	d := time.Since(e.ModifiedAt)
	return int(d.Hours() / 24)
}
