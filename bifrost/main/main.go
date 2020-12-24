package main

import (
	"learn_code_go/bifrost/core/config"
	"learn_code_go/bifrost/core/registry/nacos"
	"learn_code_go/bifrost/core/route"
	"learn_code_go/bifrost/core/server"

)

func main() {
	routeConfig := config.ParseRouteConfig("route.toml")
	routeRepo := route.GetRoutes(routeConfig)
	globalConfig := config.ParseGlobalConfig("server.toml")
	nacos.Init()
	//server初始化,必须放到最后执行
	server.Init(globalConfig.Server.Port, routeRepo, routeConfig)

}
