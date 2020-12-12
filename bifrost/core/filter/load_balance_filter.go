package filter

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"learn_code_go/bifrost/core/registry/nacos"
	"strconv"
)

type LoadBalanceFilter struct {
	Filter
}

func (l LoadBalanceFilter) Name() string {
	return "LoadBalance"
}

func (l LoadBalanceFilter) Order() int {
	return 200
}

func (l LoadBalanceFilter) DoFilter(ctx *FilterContext) {
	serviceName := ctx.RouteConfig.Url
	ip, port := GetServiceIpWrr(serviceName)
	finalUrl := "http://" + ip + ":" + strconv.FormatUint(port, 10) + ctx.finalPath
	fmt.Printf("service name:%q;finalPath:%q \n", serviceName, finalUrl)
	ctx.finalPath = finalUrl
}
func GetServiceIpWrr(serviceName string) (string, uint64) {
	instance, err := nacos.NacosNamingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   "DEFAULT_GROUP",     // default value is DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"}, // default value is DEFAULT
	})
	if err != nil {
		fmt.Println("Get service ip from nacos error,message:" + err.Error())
		return "", 0
	}
	return instance.Ip, instance.Port
}

var _ IFilter = (*LoadBalanceFilter)(nil)
