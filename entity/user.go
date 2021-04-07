package entity

import (
	"time"
)

type User struct {
	Id uint
	GithubId string `gorm:"size:64"`
	Token string `gorm:"size:64"`
	CreatedAt time.Time
}