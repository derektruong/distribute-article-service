package model

import (
	"time"
	// "gorm.io/gorm"
)

type Post struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	UserID     int       `json:"user_id"`
	Title      string    `gorm:"not null" json:"title"`
	Slug       string    `gorm:"not null" json:"slug"`
	Views      int       `gorm:"not null" json:"views"`
	Image      string    `gorm:"not null" json:"image"`
	Body       string    `gorm:"not null" json:"body"`
	Publicshed string    `gorm:"not null" json:"publicshed"`
	CreateAt   time.Time `gorm:"default: current_timestamp; not null" json:"create_at"`
	UpdateAt   time.Time `gorm:"not null" json:"update_at"`
}

func (Post) TableName() string {
	return "post"
}
