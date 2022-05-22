package models

type Account struct {
	Id       int64  `json:"id" bun:",pk"`
	EntityID int64  `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
