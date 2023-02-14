package model

type User struct {
	Name   string `query:"name" json:"first_name"`
	Family string `query:"family" json:"last_name"`
	Age    int    `query:"age" json:"age"`
	Phone  string `query:"phone" json:"phone"`
}
