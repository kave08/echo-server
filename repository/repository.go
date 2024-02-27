package repository

import (
	"context"
	"echo-server/database"
	"echo-server/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetUserList() ([]model.User, error)
	GetUserId(id string) (*model.User, error)
}

type repository struct {
	db database.Db
}

func NewUserService() Repository {
	db, err := database.InitMongo()
	if err != nil {
		log.Fatalln(err)
	}
	return repository{
		db: *db,
	}
}

func (r repository) GetUserList() ([]model.User, error) {
	var users []model.User

	c := r.db.GetUserCollection()

	cur, err := c.Find(context.TODO(), bson.D{})
	if err != nil {

		return nil, err
	}
	err = cur.All(context.TODO(), &users)
	if err != nil {

		return nil, err
	}

	return users, nil
}

func (r repository) GetUserId(id string) (*model.User, error) {
	var user model.User

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {

		return nil, err
	}

	c := r.db.GetUserCollection()

	err = c.FindOne(context.TODO(), bson.D{
		{
			Key: "_id", Value: objectId,
		},
	}).Decode(&user)
	if err != nil {

		return nil, err
	}

	return &user, nil
}
