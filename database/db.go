package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	client *mongo.Client
}

func InitMongo() (*Db, error) {
	ctx := context.TODO()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &Db{
		client: client,
	}, nil

}

func (db Db) GetUserCollection() *mongo.Collection {

	userCollection := db.client.Database("mydb").Collection("users")

	return userCollection
}
