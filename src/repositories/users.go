package repositories

import (
	"devbook-api/security"
	"devbook-api/src/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repository Users) GetUserByEmail(email string) (models.User, error) {
	var findUser models.User
	result := repository.db.First(&findUser, "email=?", email)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return findUser, nil
}

func (repository Users) FollowUser(userToFollowId, followerId string) error {
	parsedUserId, err := uuid.FromString(userToFollowId)

	if err != nil {
		return err
	}
	parsedFollowerId, err := uuid.FromString(followerId)
	if err != nil {
		return err
	}
	follower := models.Follower{UserID: parsedUserId, FollowerId: parsedFollowerId}
	result := repository.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&follower)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository Users) UnfollowUser(userToUnfollowId, followerId string) error {
	parsedUserId, err := uuid.FromString(userToUnfollowId)

	if err != nil {
		return err
	}
	parsedFollowerId, err := uuid.FromString(followerId)
	if err != nil {
		return err
	}
	follower := models.Follower{UserID: parsedUserId, FollowerId: parsedFollowerId}
	result := repository.db.First(&follower).Delete(&follower)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository Users) GetFollowers(userId string) ([]models.User, error) {
	parsedUserId, err := uuid.FromString(userId)

	if err != nil {
		return nil, err
	}

	var followers []models.User
	res := repository.db.Model(&models.User{}).Select("users.name, users.nick", "users.id").Joins("inner join followers on users.id = followers.follower_id").Where("followers.user_id = ?", parsedUserId).Scan(&followers)
	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println(followers)
	return followers, nil
}

func (repository Users) GetFollowing(userId string) ([]models.User, error) {
	parsedUserId, err := uuid.FromString(userId)

	if err != nil {
		return nil, err
	}

	var followers []models.User
	res := repository.db.Model(&models.User{}).Select("users.name, users.nick", "users.id").Joins("inner join followers on users.id = followers.user_id").Where("followers.follower_id = ?", parsedUserId).Scan(&followers)
	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println(followers)
	return followers, nil
}

func (repository Users) GetCurrentPassword(userId string) (string, error) {
	parsedUserId, err := uuid.FromString(userId)

	if err != nil {
		return "", err
	}

	var user models.User
	res := repository.db.Model(&user).Select("password").Where("id = ?", parsedUserId).Scan(&user)
	if res.Error != nil {
		return "", res.Error
	}
	return user.Password, nil
}

func (repository Users) UpdatePassword(userId, newPassword string) error {
	parsedUserId, err := uuid.FromString(userId)

	if err != nil {
		return err
	}

	var user models.User
	res := repository.db.Find(&user, "id = ?", parsedUserId)
	if res.Error != nil {
		return res.Error
	}
	user.Password = newPassword
	repository.db.Save(&user)
	return nil
}
