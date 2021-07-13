package controllers

import (
	"devbook-api/src/authentication"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
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

}
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
