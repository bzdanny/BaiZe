package slicesUtils

import "strconv"

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
		i64, _ := strconv.ParseInt(s, 10, 64)

		list = append(list, i64)
	}
	return list
}
