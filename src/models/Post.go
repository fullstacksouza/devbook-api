package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uuid.UUID `json:"id,omitempty" gorm:"primaryKey"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorID  string    `json:"authorId,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Likes     uint64    `json:"likes,omitempty"`
}
