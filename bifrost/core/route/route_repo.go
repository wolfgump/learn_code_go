package route

import (
	"learn_code_go/bifrost/core/config"
	"learn_code_go/bifrost/core/filter"
)

type RouteRepo struct {
	Routes []*Route
}

var routeRepo *RouteRepo

func GetRoutes(config config.RouteConfigs) *RouteRepo {
	if routeRepo == nil {
		filter.RegisterStaticFilter()
		routeRepo = &RouteRepo{}
		routeRepo.Routes = FromConfig(config)
	}
	return routeRepo

}
