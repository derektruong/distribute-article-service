package model

import (
// "gorm.io/gorm"
)

type RoleAccount struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	RoleName string `gorm:"not null" json:"rolename"`
}

func (RoleAccount) TableName() string {
	return "role_account"
}
