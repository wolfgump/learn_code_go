package server

import (

	"html"
	"learn_code_go/bifrost/core/config"
	"learn_code_go/bifrost/core/filter"
	"learn_code_go/bifrost/core/route"
	"net/http"
	"regexp"
	log "github.com/go-kratos/kratos/pkg/log"
)

type GateWayRequestHandler struct {
	RouteRepo    *route.RouteRepo
	RoutesConfig config.RouteConfigs
}

func (g GateWayRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Info( "Hello, %q \n", html.EscapeString(request.URL.Path))

	matchedRoute,matchedRouteConfig := findMatchRoute(g.RouteRepo, g.RoutesConfig, request.URL.Path)

	if matchedRoute == nil {
		log.Info("match none route,please check your path or connect admin configure route for you  \n")
	} else {
		log.Info("Match route:%q \n", matchedRoute.Id)
		ctx := &filter.FilterContext{ResponseWriter: writer, Request: request, RouteConfig: matchedRouteConfig}
		filter.RunChain(matchedRoute.Filters, ctx)
	}
}

func findMatchRoute(routeRepo *route.RouteRepo, routesConfig config.RouteConfigs, path string) (*route.Route, config.RouteConfig) {
	var matchRoute *route.Route
	var matchRouteConfig config.RouteConfig
	for _, route := range routeRepo.Routes {
		matchPath := route.MatchPath
		if result, err := regexp.Match(matchPath, []byte(path)); result && err == nil {
			matchRoute = route
		}
	}
	if matchRoute != nil {
		for _, routeConfig := range routesConfig.Route {
			if routeConfig.Id == matchRoute.Id {
				matchRouteConfig = routeConfig
			}
		}
	}
	return matchRoute, matchRouteConfig
}

var _ http.Handler = (*GateWayRequestHandler)(nil)
