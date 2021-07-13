package models

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorID  string    `json:"authorId,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Likes     uint64    `json:"likes,omitempty" gorm:"-"`
	User      *User     `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}
	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title is required")
	}
	if post.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}

func (post *Post) Sanitize() {
	post.User.Password = ""
}
