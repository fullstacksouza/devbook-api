package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var commentsRoutes = []Route{
	{
		URI:         "/comments",
		Method:      http.MethodPost,
		RequestAuth: true,
		Handler:     controllers.CreateComment,
	},
	{
		URI:         "/comments/{commentId}",
		Method:      http.MethodPut,
		RequestAuth: true,
		Handler:     controllers.UpdateComment,
	},
	{
		URI:         "/comments/{commentId}",
		Method:      http.MethodDelete,
		RequestAuth: true,
		Handler:     controllers.DeleteComment,
	},
	{
		URI:         "/comments/{postId}/{commentId}",
		Method:      http.MethodPost,
		RequestAuth: true,
		Handler:     controllers.GetCommentReplies,
	},
}
