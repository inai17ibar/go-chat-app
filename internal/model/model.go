package model

import "time"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"uniqueIndex"`
	HashedPassword string
	DeletedAt      *time.Time `gorm:"index"`
}
