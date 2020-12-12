package config

import (
	"fmt"
	"testing"
)

func Test_Parse_Route_Config(t *testing.T) {
	routeConfig := ParseRouteConfig("route_test.toml")
	for _, route := range routeConfig.Route {
		fmt.Println(route.Id)
	}
}