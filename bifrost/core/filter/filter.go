package filter

import (
	"fmt"
	"learn_code_go/bifrost/core/config"
	"net/http"
)

//base filter
type Filter struct {
}
type FilterContext struct {
	ResponseWriter      http.ResponseWriter
	Request     *http.Request
	RouteConfig config.RouteConfig
	finalPath   string
	Context     map[string]interface{}
}
type IFilter interface {
	Name() string
	Order() int
	DoFilter(ctx *FilterContext)
}

var _ IFilter = (*Filter)(nil)

func (b *Filter) Name() string {
	return "Base_Filter"
}
func (b *Filter) Order() int {
	return 0
}

func (b *Filter) DoFilter(ctx *FilterContext) {
	//do nothing
	fmt.Println("do nothing in base filter")
}
