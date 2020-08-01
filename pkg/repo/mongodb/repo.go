package repo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB defines the Client
type MongoDB struct {
	DB *mongo.Database
}

// DbConnect is use to connect the DB
func DbConnect(url string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err == nil {
		log.Println("Connected!")
	} else {
		return nil, err
	}

	db := client.Database("social-naka-db")

	return &MongoDB{
		DB: db,
	}, nil
}
