package routes

import (
	"net/http"

	"devbook-api/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Handler:     controllers.CreatePost,
		RequestAuth: false,
	},
	{
		URI:         "/posts/{postId}",
		Method:      http.MethodGet,
		Handler:     controllers.FindPostById,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      http.MethodPut,
		Handler:     controllers.UpdatePost,
		RequestAuth: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      http.MethodDelete,
		Handler:     controllers.DeletePost,
		RequestAuth: true,
	},
}
