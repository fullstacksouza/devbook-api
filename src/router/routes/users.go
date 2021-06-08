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
}
