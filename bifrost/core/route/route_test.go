package route

import (
	"fmt"
	"bifrost/core/config"
	"bifrost/core/filter"
	"testing"
)

func Test_Parse_Route(t *testing.T) {
	filter.RegisterStaticFilter()
	routes := FromConfig(config.RouteConfigs{})
	for _, route := range routes {
		fmt.Println(route.Id)
	}
}
