package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	task "github.com/RushinShah22/task-scheduler/model"
	db "github.com/RushinShah22/task-scheduler/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchedulerConn struct {
	db.DB
	addr string
	port string
}

// CommandRequest represents the structure of the request body
type CommandRequest struct {
	Command     string    `json:"command"`
	ScheduledAt time.Time `json:"scheduled_at"`
}

var schedulerDB SchedulerConn

func connSchedulerDB(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	schedulerDB.SetupDB()
	if err := schedulerDB.ConnectDB(ctx); err != nil {
		return err
	}
	return nil
}

// This method starts a scheduler server
func (s *SchedulerConn) SetupAndStartServer(addr, port string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	s.port = port
	s.addr = addr
	err := connSchedulerDB(ctx)

	if err != nil {
		return err
	}

	http.HandleFunc("/schedule", s.handleTaskInsert)
	var main_err error

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if addr == "" {
			addr = "localhost"
		}
		fmt.Printf("Scheduler server is running at: %s\n", addr+":"+port)
		err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.addr, s.port), nil)
		if err != nil {
			fmt.Println(err)
			main_err = err
		}

	}()

	wg.Wait()

	return main_err
}

func (*SchedulerConn) insertTaskIntoDB(t *CommandRequest) (*mongo.InsertOneResult, error) {
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
	return res, nil
}

// This method handles the POST request for inserting the task into DB.
func (s *SchedulerConn) handleTaskInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var commandReq CommandRequest
	if err := json.NewDecoder(r.Body).Decode(&commandReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if commandReq.Command == "" {
		http.Error(w, "please enter a valid command name.", http.StatusBadRequest)
	}

	log.Printf("Received schedule request: %+v", commandReq)

	res, err := s.insertTaskIntoDB(&CommandRequest{Command: commandReq.Command, ScheduledAt: commandReq.ScheduledAt})

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to submit task. Error: %s", err.Error()),
			http.StatusInternalServerError)
		return
	}

	insertedID := res.InsertedID
	var idString string
	if oid, ok := insertedID.(primitive.ObjectID); ok {
		// Convert the ObjectID to a string
		idString = oid.Hex()
	} else {
		fmt.Println("Inserted ID is not of type primitive.ObjectID")
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
	}

	// Respond with the parsed data (for demonstration purposes)
	response := struct {
		Command     string    `json:"command"`
		ScheduledAt time.Time `json:"scheduled_at"`
		TaskID      string    `json:"task_id"`
	}{
		Command:     commandReq.Command,
		ScheduledAt: commandReq.ScheduledAt,
		TaskID:      idString,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
