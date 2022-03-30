package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.FindAllUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
