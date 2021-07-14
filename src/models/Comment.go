package models

import (
	"errors"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Comment struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	PostID          uuid.UUID `json:"post_id"`
	UserID          uuid.UUID `json:"user_id"`
	Comment         string    `json:"comment"`
	InReplyToUserID string    `json:"in_reply_to_user_id,omitempty"`
}

func (comment *Comment) validate() error {
	if comment.Comment == "" {
		return errors.New("name is required")
	}
	return nil
}

func (comment *Comment) format() {
	comment.Comment = strings.TrimSpace(comment.Comment)

}

// Call methods to validate and format comment
func (comment *Comment) Prepare() error {
	if err := comment.validate(); err != nil {
		return err
	}
	comment.format()
	return nil
}
