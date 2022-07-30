package baizeContext

import (
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func (bzc *BaiZeContext) ParamInt64(key string) int64 {
	return gconv.Int64(bzc.Param(key))
}
func (bzc *BaiZeContext) ParamInt64Array(key string) []int64 {
	split := strings.Split(bzc.Param(key), ",")
	list := make([]int64, 0, len(split))
	for _, s := range split {
		list = append(list, gconv.Int64(s))
	}
	return list
}
func (bzc *BaiZeContext) ParamStringArray(key string) []string {
	return strings.Split(bzc.Param(key), ",")
}

func (bzc *BaiZeContext) QueryInt64(key string) int64 {
	return gconv.Int64(bzc.Query(key))
}
func (bzc *BaiZeContext) QueryInt64Array(key string) []int64 {
	return gconv.Int64s(strings.Split(bzc.Query(key), ","))

}
