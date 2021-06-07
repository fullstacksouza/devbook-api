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
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Handler:     controllers.FindUserById,
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Handler:     controllers.UpdateUser,
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Handler:     controllers.DeleteUser,
		RequestAuth: false,
	},
}
