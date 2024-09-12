package main

import (
	"fmt"

	"github.com/RushinShah22/task-scheduler/pkg/scheduler"
)

func main() {
	var conn scheduler.SchedulerConn

	err := conn.SetupAndStartServer("", "8080")
	if err != nil {
		fmt.Println(err)
	}
}
