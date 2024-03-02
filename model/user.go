package model

type User struct {
	Id     string `bson:"_id,omitempty"`
	Name   string `bson:"name,omitempty" query:"name" json:"first_name"`
	Family string `bson:"family,omitempty" query:"family" json:"last_name"`
	Age    int    `bson:"age,omitempty" query:"age" json:"age"`
	Phone  string `bson:"phone,omitempty" query:"phone" json:"phone"`
}
