package filter

import (
	"sort"
)

func RunChain(filters []IFilter, ctx *FilterContext) {
	sortFilter(filters)
	for _, filter := range filters {
		filter.DoFilter(ctx)
	}
}
func sortFilter(filters []IFilter) {
	sort.Slice(filters, func(i, j int) bool {
		return filters[i].Order() < filters[j].Order()
	})
}
