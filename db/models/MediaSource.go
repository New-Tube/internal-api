package db_models

import (
	"time"

	"gorm.io/gorm"
)

type FileState = uint16

const (
	Queued FileState = iota
	Rendering
	Done
)

type MediaSource struct {
	gorm.Model

	ID                  uint64 `json:"id"`
	OriginalFileID      uint64 `json:"original_file_id"`
	OriginalFileDeleted bool   `json:"original_file_deleted"`

	LowQualityFileID    uint64 `json:"low_quality_file_id"`
	MediumQualityFileID uint64 `json:"medium_quality_file_id"`
	HighQualityFileID   uint64 `json:"high_quality_file_id"`

	LowQualityFileState    FileState `json:"low_quality_file_state"`
	MediumQualityFileState FileState `json:"medium_quality_file_state"`
	HighQualityFileState   FileState `json:"high_quality_file_state"`

	CreatedAt time.Time      `json:"created_at"`
	EditedAt  time.Time      `json:"edited_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
