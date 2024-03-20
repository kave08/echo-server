package model

type Login struct {
	UserName string `validate:"required"  query:"user_name" json:"user_name"`
	Password string `validate:"required"  query:"password" json:"password"`
}
