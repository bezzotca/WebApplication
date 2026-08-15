package entity

import "time"

// User is a pure domain entity — no DB tags, no JSON tags, no framework code.
type User struct {
	ID        int64
	Email     string
	Name      string
	CreatedAt time.Time
}
