package db_models

import (
	"time"

	"gorm.io/gorm"
)

type Like struct {
	gorm.Model

	ID        uint64 `json:"id"`
	VideoID   uint64 `json:"video_id"`
	CommentID uint64 `json:"comment_id"`
	UserID    uint64 `json:"user_id"`

	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
