package db

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DB is use to represent the DB connection
type DB struct {
	URI         string
	Client      *mongo.Client
	IsConnected bool
}

// SetupDB is use to get MongoDB URI from env file

func (db *DB) SetupDB() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env File Found")
		return err
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Println("No uri found in env file.")
		return errors.New("no uri found in env file")
	}
	db.URI = uri
	return nil

}

// ConnectDB is use to connect to the DB

func (db *DB) ConnectDB(ctx context.Context) error {

	retry := 5
	var err_main error
	for ; retry > 0; retry-- {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.URI))

		if err != nil {
			return err
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			err_main = err
			continue
		}
		db.IsConnected = true
		db.Client = client
		return nil
	}

	return err_main

}
