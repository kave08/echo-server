package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitMongo() {
	ctx := context.TODO()

	client, err := mongo.Connect(
		ctx,
	)
	if err != nil {
		return
	}
	fmt.Println(client)
}
