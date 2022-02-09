package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri           string
	Method        string
	Func          func(http.ResponseWriter, *http.Request)
	Authenticated bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := routesStates

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}

	return r
}
