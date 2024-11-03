package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const articleTableName = "articles"

type ArticleDbo struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null" json:"title"`
	Description *string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
}

func (ArticleDbo) TableName() string {
	return articleTableName
}
