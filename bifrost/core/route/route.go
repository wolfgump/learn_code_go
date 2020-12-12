package route

import (
	"learn_code_go/bifrost/core/config"
	"learn_code_go/bifrost/core/filter"
)

type Route struct {
	Id        string
	Url       string
	Filters   []filter.IFilter
	MatchPath string
}

func FromConfig(config config.RouteConfigs) []*Route {
	if len(config.Route) <= 0 {
		return nil
	}
	routes := make([]*Route, len(config.Route))
	for cnt, routeConfig := range config.Route {
		routes[cnt] = &Route{
			Id:        routeConfig.Id,
			Url:       routeConfig.Url,
			Filters:   filtersFormConfig(routeConfig.Filters),
			MatchPath: routeConfig.MatchPath,
		}
	}
	return routes
}
func filtersFormConfig(filterConfigs []config.FilterConfig) []filter.IFilter {
	filters := make([]filter.IFilter, len(filterConfigs))
	for cnt, filterConfig := range filterConfigs {
		filterInstance := filter.GetFilterByName(filterConfig.Name)
		if filterInstance == nil {
			panic(filterConfig.Name + " not exist")
		}
		filters[cnt] = filterInstance
	}
	return filters
}
