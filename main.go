package main

import (
	"fmt"
	"log"

	db "github.com/RushinShah22/task-scheduler/utils"
)

func main() {
	fmt.Println("Rushin Shah")

	var client db.DB

	if err := client.SetupDB(); err != nil {
		log.Fatal(err)
	}
	if err := client.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB.")

}
