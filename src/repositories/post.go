package repositories

import (
	"devbook-api/src/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Posts struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *Posts {
	var post Posts
	db.Model(&post).Association("Users")

	return &Posts{db}
}

func (postRepository Posts) Create(post models.Post) (models.Post, error) {
	postId := uuid.NewV4()
	newPost := models.Post{
		ID:       postId,
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}
	result := postRepository.db.Omit("likes").Create(&newPost)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	createdPost := postRepository.db.Preload("User").Find(&newPost, "id = ?", postId)
	if createdPost.Error != nil {
		return models.Post{}, result.Error
	}
	newPost.Sanitize()
	return newPost, nil
}

func (postRepository Posts) FindPostById(postId string) (models.Post, error) {
	var post models.Post
	result := postRepository.db.Preload("User").First(&post, "id= ?", postId)
	if result.Error != nil {
		return models.Post{}, result.Error
	}
	post.Sanitize()
	return post, nil
}

func (postRepository *Posts) GetPosts(userId string) ([]models.Post, error) {
	var posts []models.Post

	result := postRepository.db.Preload("User").Select("distinct(posts.id),posts.title,posts.content,posts.author_id,count(likes) as likes").Joins("left join likes on posts.id = likes.post_id").Group("posts.id").Find(&posts)
	if result.Error != nil {
		return []models.Post{}, result.Error
	}
	for _, post := range posts {
		post.Sanitize()
	}
	return posts, nil
}

func (postRepository Posts) UpdatePost(postId string, post models.Post) (models.Post, error) {
	var findPost models.Post

	result := postRepository.db.Preload("User").First(&findPost, "id = ?", postId)
	if result.Error != nil {
		return models.Post{}, result.Error
	}
	findPost.Content = post.Content
	findPost.Title = post.Title
	findPost.Prepare()
	postRepository.db.Save(&findPost)
	findPost.Sanitize()
	return findPost, nil
}

func (postRepository Posts) DeletePost(postId string) error {
	var findPost models.Post

	result := postRepository.db.First(&findPost, "id = ?", postId)
	if result.Error != nil {
		return result.Error
	}
	postRepository.db.Delete(&findPost)
	return nil
}

func (postRepository Posts) GetPostsByUserId(userId string) ([]models.Post, error) {
	var posts []models.Post
	result := postRepository.db.Preload("User").Find(&posts, "author_id = ?", userId)

	if result.Error != nil {
		return posts, result.Error
	}
	for _, post := range posts {
		post.Sanitize()
	}
	return posts, nil
}

func (postRepository Posts) LikePost(postId, userId string) error {

	parsedPostId, err := uuid.FromString(postId)
	if err != nil {
		return err
	}
	parsedUserId, err := uuid.FromString(userId)
	if err != nil {
		return err
	}
	like := models.Like{UserID: parsedUserId, PostID: parsedPostId}

	result := postRepository.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&like)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (postRepository Posts) UnlikePost(postId, userId string) error {

	parsedPostId, err := uuid.FromString(postId)
	if err != nil {
		return err
	}
	parsedUserId, err := uuid.FromString(userId)
	if err != nil {
		return err
	}
	like := models.Like{UserID: parsedUserId, PostID: parsedPostId}

	result := postRepository.db.First(&like).Delete(&like)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
