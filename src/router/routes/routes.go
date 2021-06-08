package routes

import (
	"devbook-api/src/middlewares"
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
		if route.RequestAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Authenticate(route.Handler)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)

		}
	}
	return r
}
