package controllers

import (
	"devbook-api/security"
	"devbook-api/src/authentication"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	createdUser, err := userRepository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, createdUser)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusFound, users)
}
func FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	findUser, err := repository.GetUserById(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusFound, findUser)

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	userId := params["userId"]
	tokenUserId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userId != tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}
	var userData models.User
	if err = json.Unmarshal(requestBody, &userData); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	updatedUser, err := repository.Update(userId, userData)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	tokenUserId, err := authentication.ExtractUserId(r)
	if userId != tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	err = repository.DeleteUser(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	tokenUserId, err := authentication.ExtractUserId(r)
	if userId == tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	err = repository.FollowUser(userId, tokenUserId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	tokenUserId, err := authentication.ExtractUserId(r)
	if userId == tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	err = repository.UnfollowUser(userId, tokenUserId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	followers, err := repository.GetFollowers(userId)
	fmt.Println(followers)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)

}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	followers, err := repository.GetFollowing(userId)
	fmt.Println(followers)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["userId"]

	tokenUserId, err := authentication.ExtractUserId(r)
	if userId != tokenUserId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var passowrd models.Password
	if err = json.Unmarshal(requestBody, &passowrd); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewUserRepository(db)
	currentPassword, err := repository.GetCurrentPassword(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPassword(currentPassword, passowrd.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	hashNewPassword, err := security.Hash(passowrd.NewPassword)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = repository.UpdatePassword(userId, string(hashNewPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
