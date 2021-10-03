package entity

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	UID       string
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
