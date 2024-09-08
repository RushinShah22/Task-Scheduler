package db

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is use to represent the DB connection
type DB struct {
	URI    string
	Client *mongo.Client
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

func (db *DB) ConnectDB() error {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.URI))

	if err != nil {
		return err
	}
	db.Client = client

	return nil
}
