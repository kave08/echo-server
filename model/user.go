package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id     primitive.ObjectID
	Name   string `query:"name" json:"first_name"`
	Family string `query:"family" json:"last_name"`
	Age    int    `query:"age" json:"age"`
	Phone  string `query:"phone" json:"phone"`
}
