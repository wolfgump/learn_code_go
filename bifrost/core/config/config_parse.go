package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

func ParseGlobalConfig(path string) GlobalConfig {
	var globalConfig GlobalConfig
	if _, err := toml.DecodeFile(path, &globalConfig); err != nil {
		log.Fatalln(err)
		return globalConfig
	}
	return globalConfig
}
func ParseRouteConfig(path string) RouteConfigs {
	var routeConfig RouteConfigs
	if _, err := toml.DecodeFile(path, &routeConfig); err != nil {
		log.Fatalln(err)
		return routeConfig
	}
	return routeConfig
}
