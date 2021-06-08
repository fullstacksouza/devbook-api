package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var authRoute = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Handler:     controllers.Auth,
	RequestAuth: false,
}
