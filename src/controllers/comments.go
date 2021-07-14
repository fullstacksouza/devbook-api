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
	uuid "github.com/satori/go.uuid"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	var comment models.Comment

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(requestBody, &comment); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err := comment.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	comment.UserID = uuid.FromStringOrNil(userId)
	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	postRepository := repositories.NewPostRepository(db)

	_, err = postRepository.FindPostById(comment.PostID.String())

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	commentRepository := repositories.NewCommentRepository(db)

	newComment, err := commentRepository.CreateComment(comment)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, newComment)

}
func GetComments(w http.ResponseWriter, r *http.Request) {

}
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentId := params["commentId"]
	_, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	var comment models.Comment

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(requestBody, &comment); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err := comment.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	commentRepository := repositories.NewCommentRepository(db)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	comment.ID = uuid.FromStringOrNil(commentId)
	updatedComment, err := commentRepository.UpdateComment(comment)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, updatedComment)

}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentId := params["commentId"]
	_, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()

	commentRepository := repositories.NewCommentRepository(db)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	deleteError := commentRepository.DeleteComment(commentId)

	if deleteError != nil {
		responses.Error(w, http.StatusInternalServerError, deleteError)
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
func GetCommentReplies(w http.ResponseWriter, r *http.Request) {

}
