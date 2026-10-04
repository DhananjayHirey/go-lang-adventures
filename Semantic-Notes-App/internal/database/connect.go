package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func Connect(uri string) (*mongo.Client, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(
		context.Background(),
		readpref.Primary(),
	); err != nil {

		return nil, err
	}

	fmt.Println("✅ MongoDB Connected")

	return client, nil

}
