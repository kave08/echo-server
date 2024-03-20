package model

import "time"

type Login struct {
	UserName string `validate:"required"  query:"user_name" json:"user_name"`
	Password string `validate:"required"  query:"password" json:"password"`
}

type User struct {
	Id         string    `bson:"_id,omitempty"`
	FirstName  string    `bson:"first_name,omitempty" validate:"required" query:"first_name" json:"first_name"`
	LastName   string    `bson:"last_name,omitempty" validate:"required" query:"last_name" json:"last_name"`
	Age        int       `bson:"age,omitempty" query:"age" json:"age"`
	Email      string    `bson:"email,omitempty" validate:"required" query:"email" json:"email"`
	Phone      string    `bson:"phone,omitempty" validate:"required" query:"phone" json:"phone"`
	UserName   string    `bson:"user_name,omitempty" validate:"required" query:"user_name" json:"user_name"`
	Password   string    `bson:"password,omitempty" validate:"required" query:"password" json:"password"`
	Created_at time.Time `bson:"created_at,omitempty" query:"created_at" json:"created_at"`
}
