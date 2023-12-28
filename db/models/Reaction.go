package db_models

import (
	"time"

	"gorm.io/gorm"
)

type Reaction struct {
	gorm.Model

	ID        uint64 `json:"id"`
	VideoID   uint64 `json:"video_id"`
	CommentID uint64 `json:"comment_id"`
	UserID    uint64 `json:"user_id"`

	IsLike    bool `json:"is_like"`
	IsDislike bool `json:"is_dislike"`

	CreatedAt time.Time `json:"created_at"`
}
