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

type DB struct {
	URI    string
	Client *mongo.Client
}

func (db *DB) setupDB() error {
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

func (db *DB) connectDB() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.URI))

	if err != nil {
		return nil, err
	}
	db.Client = client

	return client, nil
}
