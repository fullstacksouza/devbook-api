package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Handler:     controllers.CreateUser,
		RequestAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Handler:     controllers.FindUsers,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Handler:     controllers.FindUserById,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Handler:     controllers.UpdateUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Handler:     controllers.DeleteUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Handler:     controllers.FollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}/unfollow",
		Method:      http.MethodPost,
		Handler:     controllers.UnfollowUser,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}/followers",
		Method:      http.MethodGet,
		Handler:     controllers.GetFollowers,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}/following",
		Method:      http.MethodGet,
		Handler:     controllers.GetFollowing,
		RequestAuth: true,
	},
	{
		URI:         "/users/{userId}/update-password",
		Method:      http.MethodPost,
		Handler:     controllers.UpdatePassword,
		RequestAuth: true,
	},
}
