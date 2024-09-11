package task

import (
	"time"
)

type Task struct {
	Command     string    `json:"name"`
	ScheduledAt time.Time `json:"scheduledAt"`
	PickedAt    time.Time `json:"pickedAt"`
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
	FailedAt    time.Time `json:"failedAt"`
}
