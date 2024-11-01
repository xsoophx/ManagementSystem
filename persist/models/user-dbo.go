package models

import (
	"github.com/google/uuid"
	"time"
)

const tableName = "users"

type UserDbo struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FirstName string     `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string     `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}

func (UserDbo) TableName() string {
	return tableName
}
