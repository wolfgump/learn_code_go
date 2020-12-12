package filter

import (
	"fmt"
	"learn_code_go/bifrost/core/config"
	"strings"
)

const defaultParts = 2

type PrefixStripFilter struct {
	Filter
}

var _ IFilter = (*PrefixStripFilter)(nil)

func (psf *PrefixStripFilter) Name() string {
	return "PrefixStrip"
}
func (psf *PrefixStripFilter) Order() int {
	return 100
}
func (psf *PrefixStripFilter) DoFilter(ctx *FilterContext) {
	fmt.Println( "prefix strip filter do filter")
	//psf.Filter.DoFilter(ctx)
	routeConfig := config.FindFilterConfig(psf.Name(), ctx.RouteConfig.Filters)
	parts := routeConfig.Param["parts"]
	if parts == nil {
		parts = defaultParts
	}
	path := ctx.Request.URL.Path
	pathArray := strings.Split(path, "/")
	if len(pathArray) <= (int(parts.(int64))) {
		return
	}
	pathArrayNew := pathArray[3:]
	newPath := strings.Join(pathArrayNew, "/")
	ctx.finalPath = "/" + newPath
	fmt.Printf( "final path:%q \n", ctx.finalPath)
}
