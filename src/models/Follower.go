package models

import (
	uuid "github.com/satori/go.uuid"
)

type Follower struct {
	UserID     uuid.UUID `json:"user_id,omitempty" gorm:"primaryKey"`
	FollowerId uuid.UUID `json:"follower_id,omitempty" gorm:"primaryKey"`
}
