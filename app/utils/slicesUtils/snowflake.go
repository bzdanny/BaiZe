package slicesUtils

import (
	"github.com/gogf/gf/util/gconv"
)

type Slices []string

func (sli Slices) Contains(str string) int {
	for i, item := range sli {
		if item == str {
			return i
		}
	}
	return -1
}

func (sli Slices) StrSlicesToInt() []int64 {
	list := make([]int64, 0, len(sli))
	for _, s := range sli {
		list = append(list, gconv.Int64(s))
	}
	return list
}
