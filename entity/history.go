package entity

import (
	"time"
)

type History struct {
	Id uint
	UserId uint
	Type string `gorm:"size:256"`
	Url string `gorm:"size:256"`
	Impression string `gorm:"size:1024"`
	CreatedAt time.Time
}