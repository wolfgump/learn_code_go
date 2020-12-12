package filter

import (
	"fmt"
	"testing"
)

func Test_Chain_Sort(t *testing.T) {
	filters := mockFilters()
	fmt.Println("before sort:")
	for _, filter := range filters {
		fmt.Println(filter.Name())
	}
	sortFilter(filters)

	fmt.Println("after sort:")
	for _, filter := range filters {
		fmt.Println(filter.Name())
	}
}
func mockFilters() []IFilter {
	filter1 := &Filter{}
	filter2 := &Filter{}
	filter3 := &PrefixStripFilter{Filter{}}
	filter4 := &Filter{}
	filter5 := &PrefixStripFilter{Filter{}}
	filters := []IFilter{filter1, filter2, filter3, filter4, filter5}
	return filters
}

func Test_Run(t *testing.T) {
	context := &FilterContext{ResponseWriter: nil}
	RunChain(mockFilters(), context)
}
