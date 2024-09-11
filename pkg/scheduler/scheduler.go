package scheduler

import (
	"context"
	"fmt"
	"time"

	task "github.com/RushinShah22/task-scheduler/model"
	db "github.com/RushinShah22/task-scheduler/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// CommandRequest represents the structure of the request body
type CommandRequest struct {
	Command     string    `json:"command"`
	ScheduledAt time.Time `json:"scheduled_at"` // ISO 8601 format
}

var schedulerDB db.DB

func connSchedulerDB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	schedulerDB.SetupDB()
	if err := schedulerDB.ConnectDB(ctx); err != nil {
		return err
	}
	return nil
}

func InsertTaskIntoDB(t *CommandRequest) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if !schedulerDB.IsConnected {
		if err := connSchedulerDB(ctx); err != nil {
			return nil, err
		}
	}

	collection := schedulerDB.Client.Database("task-scheduler").Collection("Task")

	res, err := collection.InsertOne(ctx, task.Task{Command: t.Command, ScheduledAt: t.ScheduledAt})

	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
