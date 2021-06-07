package models

import (
	"errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Nick == "" {
		return errors.New("nick is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

}

// Call methods to validate and format user
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}
