package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesStates = []Route{
	{
		Uri:           "/states",
		Method:        http.MethodPost,
		Func:          controllers.Create,
		Authenticated: false,
	},
	{
		Uri:           "/states",
		Method:        http.MethodGet,
		Func:          controllers.Show,
		Authenticated: false,
	},
	{
		Uri:           "/states/{stateUF}",
		Method:        http.MethodGet,
		Func:          controllers.View,
		Authenticated: false,
	},
	{
		Uri:           "/states/{stateUF}",
		Method:        http.MethodPut,
		Func:          controllers.Update,
		Authenticated: false,
	},
	{
		Uri:           "/states/{stateUF}",
		Method:        http.MethodDelete,
		Func:          controllers.Delete,
		Authenticated: false,
	},
}
