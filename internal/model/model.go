package model

import (
	"time"
)

type Users struct {
	Acct       string    `gorm:"size:50;primary_key;column=acct"`
	Pwd        string    `gorm:"type:text;not null;column=pwd"`
	Fullname   string    `gorm:"size:50;not null;column=fullname;"`
	Created_at time.Time `gorm:"type:timestamp;not null;column=created_at;default:CURRENT_TIMESTAMP"`
	Updated_at time.Time `gorm:"type:timestamp;not null;column=updated_at;default:CURRENT_TIMESTAMP"`
}

func (Users) TableName() string {
	return "users"
}

type ApiUsers struct {
	Acct       string    `json:"account"`
	Fullname   string    `json:"fullname"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreateUserData struct {
	Account  string `json:"account" validate:"min=4,max=50,regexp=^[a-zA-Z0-9]*"`
	Fullname string `json:"fullname" validate:"min=4,max=50,regexp=^[a-zA-Z0-9]*"`
	Password string `json:"password" validate:"min=8,max=50"`
}
