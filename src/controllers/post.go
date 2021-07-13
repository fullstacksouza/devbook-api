package controllers

import (
	"devbook-api/src/authentication"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	post.AuthorID = userId
	if err = post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewPostRepository(db)
	createdPost, err := repository.Create(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, createdPost)

}
func FindPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["postId"]

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewPostRepository(db)

	findPost, err := repository.FindPostById(postId)
	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, findPost)

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewPostRepository(db)
	posts, err := repository.GetPosts(userId)
	if err != nil {
		responses.Error(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	postId := mux.Vars(r)["postId"]
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var postData models.Post
	if err = json.Unmarshal(requestBody, &postData); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db)

	findPost, err := repository.FindPostById(postId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if findPost.AuthorID != userId {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	updatedPost, err := repository.UpdatePost(postId, postData)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedPost)

}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	postId := mux.Vars(r)["postId"]

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db)

	findPost, err := repository.FindPostById(postId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if findPost.AuthorID != userId {
		responses.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	err = repository.DeletePost(postId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db)

	userPosts, err := repository.GetPostsByUserId(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, userPosts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	params := mux.Vars(r)
	postId := params["postId"]

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewPostRepository(db)

	likeError := repository.LikePost(postId, userId)

	if likeError != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	params := mux.Vars(r)
	postId := params["postId"]

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewPostRepository(db)

	likeError := repository.UnlikePost(postId, userId)

	if likeError != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func GetPostLikes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["postId"]

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewPostRepository(db)

	likes, err := repository.GetPostLikes(postId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, likes)
}
