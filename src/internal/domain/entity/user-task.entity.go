package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserTask struct {
	UserID    *uuid.UUID     `json:"user_id" gorm:"primary_key"`
	TaskID    *uuid.UUID     `json:"task_id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp;autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;type:timestamp"`
}
