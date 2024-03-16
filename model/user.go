package model

import "time"

type User struct {
	Id         string    `bson:"_id,omitempty"`
	FirstName  string    `bson:"first_name,omitempty" query:"first_name" json:"first_name"`
	LastName   string    `bson:"last_name,omitempty" query:"last_name" json:"last_name"`
	Age        int       `bson:"age,omitempty" query:"age" json:"age"`
	Email      string    `bson:"email,omitempty" query:"email" json:"email"`
	Phone      string    `bson:"phone,omitempty" query:"phone" json:"phone"`
	UserName   string    `bson:"user_name,omitempty" query:"user_name" json:"user_name"`
	Password   string    `bson:"password,omitempty" query:"password" json:"password"`
	Created_at time.Time `bson:"created_at,omitempty" query:"created_at" json:"created_at"`
}
