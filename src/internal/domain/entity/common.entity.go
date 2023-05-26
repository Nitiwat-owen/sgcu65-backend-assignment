package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sgcu65-backend-assignment/src/internal/utils"
	"time"
)

type Base struct {
	ID        *uuid.UUID `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:timestamp;autoUpdateTime:nano"`
	DeletedAt time.Time  `json:"deleted_at" gorm:"index;type:timestamp"`
}

func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if b.ID == nil {
		b.ID = utils.UUIDAdr(uuid.New())
	}

	return nil
}
