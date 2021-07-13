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
		URI:         "/posts",
		Method:      http.MethodGet,
		Handler:     controllers.GetAllPosts,
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
	{
		URI:         "/posts/{postId}/like",
		Method:      http.MethodPost,
		Handler:     controllers.LikePost,
		RequestAuth: false,
	},
	{
		URI:         "/posts/{postId}/unlike",
		Method:      http.MethodPost,
		Handler:     controllers.UnlikePost,
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}/posts",
		Method:      http.MethodGet,
		Handler:     controllers.GetPostsByUserId,
		RequestAuth: true,
	},
}
