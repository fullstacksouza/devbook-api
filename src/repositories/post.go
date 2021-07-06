package repositories

import (
	"devbook-api/src/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Posts struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *Posts {
	return &Posts{db}
}

func (postRepository Posts) Create(post models.Post, userId string) (models.Post, error) {
	postId := uuid.NewV4()
	newPost := models.Post{
		ID:       postId,
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: userId,
	}
	result := postRepository.db.Omit("likes").Create(&newPost)
	if result.Error != nil {
		return models.Post{}, result.Error
	}
	fmt.Println(newPost)
	return newPost, nil
}
