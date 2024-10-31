package models

import (
	"github.com/google/uuid"
	"time"
)

type ArticleDbo struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}
