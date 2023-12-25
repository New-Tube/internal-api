package db_models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	ID        uint64 `json:"id"`
	Text      string `json:"text"`
	VideoID   uint64 `json:"video_id"`
	CommentID uint64 `json:"comment_id"`
	UserID    uint64 `json:"user_id"`

	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`

	CreatedAt time.Time      `json:"created_at"`
	EditedAt  time.Time      `json:"edited_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
