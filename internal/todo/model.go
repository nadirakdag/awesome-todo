package todo

import "time"

type Task struct {
	ID          string
	Description string
	Status      bool
	CreatedAt   time.Time
}
