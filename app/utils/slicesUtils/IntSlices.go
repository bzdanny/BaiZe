package slicesUtils

import "strconv"

func IntSlicesToString(slices []int64) []string {
	strings := make([]string, 0, len(slices))
	for _, sl := range slices {
		strings = append(strings, strconv.FormatInt(sl, 10))
	}
	return strings
}
