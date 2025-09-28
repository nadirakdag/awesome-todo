package todo

import (
	"time"

	"github.com/mergestat/timediff"
)

type Task struct {
	ID          string
	Description string
	Status      bool
	CreatedAt   time.Time
}

func (t *Task) GetTimeDiff() string {
	return timediff.TimeDiff(t.CreatedAt)
}
