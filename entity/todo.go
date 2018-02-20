package entity

import "time"

// Todo type
type Todo struct {
	ID        string
	Content   string
	Done      bool
	CreatedAt time.Time
}
