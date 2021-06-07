package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Handler     func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func Setup(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoute)
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
