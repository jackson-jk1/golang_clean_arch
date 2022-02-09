package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		Uri:           "/users",
		Method:        http.MethodPost,
		Func:          controllers.CreateUser,
		Authenticated: false,
	},
	{
		Uri:           "/users/{userId}",
		Method:        http.MethodGet,
		Func:          controllers.ViewUser,
		Authenticated: false,
	},
	{
		Uri:           "/users/{userId}",
		Method:        http.MethodPut,
		Func:          controllers.UpdateUser,
		Authenticated: false,
	},
	{
		Uri:           "/users/{userId}",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteUser,
		Authenticated: false,
	},
}
