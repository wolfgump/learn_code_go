package filter

import (
	"fmt"
	"testing"
)

func Test_Filter_Register(t *testing.T) {
	RegisterStaticFilter()
	filters := AllFilters()
	for _, filter := range filters {
		fmt.Println(filter.Name())
	}

}
