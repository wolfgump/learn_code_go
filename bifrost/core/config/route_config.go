package config

type RouteConfigs struct {
	Route []RouteConfig
}
type RouteConfig struct {
	Id      string
	Filters []FilterConfig
	//后端服务的url,支持http（http://10.11.13.13:8080）、负载均衡（lb://serviceId）
	Url string
	//匹配路径的正则表达式
	MatchPath string
}
type FilterConfig struct {
	Name  string
	Param map[string]interface{}
	Url   URL
}
type URL struct {
	Host string
}

func FindFilterConfig(id string, filterConfigs []FilterConfig) FilterConfig {
	for _, routeConfig := range filterConfigs {
		if routeConfig.Name == id {
			return routeConfig
		}
	}
	return FilterConfig{}
}
