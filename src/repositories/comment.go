package repositories

import (
	"devbook-api/src/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Comments struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *Comments {
	return &Comments{db}
}

func (repository Comments) CreateComment(newComment models.Comment) (*models.Comment, error) {
	commentId := uuid.NewV4()

	var saveComment = models.Comment{
		ID:      commentId,
		PostID:  newComment.PostID,
		UserID:  newComment.UserID,
		Comment: newComment.Comment,
	}

	result := repository.db.Create(&saveComment)

	if result.Error != nil {
		return nil, result.Error
	}
	return &saveComment, nil

}

func (repository Comments) UpdateComment(newComment models.Comment) (*models.Comment, error) {

	var comment models.Comment

	result := repository.db.Find(&comment, "id = ?", newComment.ID)

	comment.Comment = newComment.Comment

	repository.db.Save(&comment)

	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil

}

func (repository Comments) DeleteComment(commentId string) error {

	var comment models.Comment

	result := repository.db.Find(&comment, "id = ?", commentId)

	if result.Error != nil {
		return result.Error
	}

	deleteError := repository.db.Delete(&comment).Error

	if deleteError != nil {
		return deleteError
	}

	return nil

}
