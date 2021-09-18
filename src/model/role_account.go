package model

import (
// "gorm.io/gorm"
)

type RoleAccount struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	RoleName string `gorm:"not null" json:"role_name"`
}

func (RoleAccount) TableName() string {
	return "role_account"
}
