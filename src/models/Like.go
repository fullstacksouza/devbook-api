package models

import uuid "github.com/satori/go.uuid"

type Like struct {
	PostID uuid.UUID `json:"post_id" gorm:"primaryKey"`
	UserID uuid.UUID `json:"user_id" gorm:"primaryKey"`
}
