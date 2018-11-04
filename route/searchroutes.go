package route

import (
	"seedapi/model"
	"seedapi/service"
	"seedapi/controller"
	"go.uber.org/zap"
)

// SearchRoutes : list of all routes
func SearchRoutes(config *model.AppSetting,  Logger *zap.Logger) Routes {
	
	//Configure service 
	searchService := service.SearchServiceInstance(config, Logger)
	//Configure controller
	searchController := controller.SearchControllerInstance(searchService)


	r :=  Routes{
		model.Route{
			"GetALL",
			"GET",
			"/search/getall",
			searchController.Index,
			true,
		},
		model.Route{
			"GetALL",
			"GET",
			"/search/panic",
			searchController.PanicRequest,
			false,
		},
		model.Route{
			"GetALL",
			"GET",
			"/search/error",
			searchController.CustomError,
			false,
		},
	}
	return r
}


