package util



import (
	"seedapi/model"
	"seedapi/service"
	"seedapi/controller"
)

// SearchRoutes : list of all routes
func SearchRoutes(config *model.AppSetting) Routes {
	
	//Configure service 
	searchService := service.SearchServiceInstance(config)
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
	}
	return r
}


