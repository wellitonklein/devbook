package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodGet,
		Function:               controllers.FindAllPublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.FindPublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{id}/publications",
		Method:                 http.MethodGet,
		Function:               controllers.FindByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}/like",
		Method:                 http.MethodPost,
		Function:               controllers.Like,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{id}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.Unlike,
		RequiresAuthentication: true,
	},
}
