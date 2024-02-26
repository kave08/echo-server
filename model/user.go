package model

type User struct {
	Id     string `bson:"_id"`
	Name   string `bson:"_id" query:"name" json:"first_name"`
	Family string `bson:"_id" query:"family" json:"last_name"`
	Age    int    `bson:"_id" query:"age" json:"age"`
	Phone  string `bson:"_id" query:"phone" json:"phone"`
}
