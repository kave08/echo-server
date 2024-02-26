package repository

import (
	"context"
	"echo-server/database"
	"echo-server/model"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	GetUserList() ([]model.User, error)
}

type repository struct {
	db database.Db
}

func NewUserService() Repository {
	return repository{}
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
