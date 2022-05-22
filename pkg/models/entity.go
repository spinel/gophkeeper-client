package models

const (
	EntityTypeAccount    = "account"
	EntityTypeCreditCard = "creditcard"
)

type Entity struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Identifier string    `json:"identifier" bun:",unique,notnull"`
	TypeID     int       `json:"type_id"`
	EntityPod  EntityPod `json:"entity"`
}

type EntityPod struct {
	Account Account `json:"account"`
	//CreditCard CreditCard `json:"credit_card"`
}
