package filter

import (
	"testing"
)

func Test_Pre_Construct(t *testing.T) {
	filter := PrefixStripFilter{Filter{}}
	context := &FilterContext{ResponseWriter: nil}
	filter.DoFilter(context)
}
