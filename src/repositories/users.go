package repositories

import (
	"devbook-api/security"
	"devbook-api/src/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Users {
	return &Users{db}
}

//Create insert user in db and return user
func (repository Users) Create(user models.User) (models.User, error) {
	userId := uuid.NewV4()
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}
	saveUser := models.User{
		ID:       userId,
		Name:     user.Name,
		Nick:     user.Nick,
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	result := repository.db.Create(&saveUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	repository.db.Omit("password").First(&saveUser)
	return saveUser, nil
}

func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	var users []models.User
	like := fmt.Sprintf("%%%s%%", nameOrNick)
	result := repository.db.Where("name like ?", like).Or("nick like ?", like).Find(&users)
	if result.Error != nil {
		return []models.User{}, result.Error
	}
	return users, nil
}

func (repository Users) GetUserById(id string) (models.User, error) {
	var user models.User
	result := repository.db.First(&user, "id= ?", id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (repository Users) Update(id string, user models.User) (models.User, error) {
	var findUser models.User
	result := repository.db.First(&findUser, "id= ?", id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	findUser.Name = user.Name
	findUser.Email = user.Email
	findUser.Nick = user.Nick
	findUser.Prepare()
	repository.db.Save(&findUser)
	return findUser, nil
}

func (repository Users) DeleteUser(id string) error {
	var user models.User
	result := repository.db.First(&user, "id= ?", id)
	repository.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
