package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

var NacosNamingClient naming_client.INamingClient

func Init() {
	if NacosNamingClient != nil {
		return
	}
	clientConfig := constant.ClientConfig{
		//NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //we can create multiple clients with different namespaceId to support multiple namespace
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "nacos.zmaxis.com",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
	}

	// Create naming client for service discovery
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic("init nacos error,message:" + err.Error())
	}
	NacosNamingClient = namingClient
}
