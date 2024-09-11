package main

import (
	"fmt"
	"time"

	"github.com/RushinShah22/task-scheduler/pkg/scheduler"
)

func main() {

	task := scheduler.CommandRequest{
		Command: "pwd", ScheduledAt: time.Now(),
	}
	result, err := scheduler.InsertTaskIntoDB(&task)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
