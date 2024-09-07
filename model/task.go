package task

import (
	"time"
)

type Tasks struct {
	Name        string    `json:"name"`
	ScheduledAt time.Time `json:"scheduledAt"`
	StartedAt   time.Time `json:"startedAt"`
	FinishedAt  time.Time `json:"finishedAt"`
}
