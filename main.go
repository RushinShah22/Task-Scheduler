package main

import (
	"github.com/RushinShah22/task-scheduler/pkg/scheduler"
)

func main() {
	var conn scheduler.SchedulerConn

	conn.SetupAndStartServer("", "8080")
}
