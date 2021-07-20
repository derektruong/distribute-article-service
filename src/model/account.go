package model

import (
// "gorm.io/gorm"
)

type Account struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func (Account) TableName() string {
	return "account"
}
