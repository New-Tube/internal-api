package db_models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`

	CreatedAt time.Time      `json:"created_at"`
	EditedAt  time.Time      `json:"edited_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
