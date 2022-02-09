package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesStates = []Route{
	{
		Uri:           "/states",
		Method:        http.MethodPost,
		Func:          controllers.CreateState,
		Authenticated: false,
	},
	{
		Uri:           "/states",
		Method:        http.MethodGet,
		Func:          controllers.ShowState,
		Authenticated: false,
	},
	{
		Uri:           "/states/{stateUF}",
		Method:        http.MethodGet,
		Func:          controllers.ViewState,
		Authenticated: false,
	},
	{
		Uri:           "/states/{stateUF}",
		Method:        http.MethodPut,
		Func:          controllers.UpdateState,
		Authenticated: false,
	},
	{
		Uri:           "/states/{stateUF}",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteState,
		Authenticated: false,
	},
}
