package models

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	Id          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	UserId      uuid.UUID  `json:"userId"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"createdAt"`
}
