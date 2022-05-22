package models

// User is a canonical user model
type User struct {
	Base
	ID       int    `json:"id" bun:",pk,autoincrement"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"-" validate:"required"`
	Token    string `json:"-"`
}

type UserForm struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
}
