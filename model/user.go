package model

type User struct {
	Account  string `json:"account" gorm:"column:account;primary_key"`
	Password string `json:"password"`
}
