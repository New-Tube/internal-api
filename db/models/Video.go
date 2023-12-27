package db_models

import (
	"time"

	"gorm.io/gorm"
)

type Privacy = uint16

const (
	Public Privacy = iota
	Link
	Private
)

type Video struct {
	gorm.Model

	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"`
	User        *User  `json:"user"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Privacy Privacy `json:"privacy"`
	Link    string  `json:"link"`

	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`

	MediaSourceID uint64       `json:"media_source_id"`
	MediaSource   *MediaSource `json:"media_source"`

	CreatedAt time.Time      `json:"created_at"`
	EditedAt  time.Time      `json:"edited_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
