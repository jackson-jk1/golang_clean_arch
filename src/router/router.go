package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func ReturnRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}
