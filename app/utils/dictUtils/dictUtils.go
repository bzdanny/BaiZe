package dictUtils

import (
	redis "baize/app/common/redis"
	"baize/app/constant/constants"
	systemModels "baize/app/system/models/systemModels"
	"encoding/json"
)

func GetDictCache(dictType string) (dictDataList []*systemModels.SysDictDataVo) {
	s := constants.SysDictKey + dictType
	getString := redis.GetString(s)
	if getString != "" {
		dictDataList = make([]*systemModels.SysDictDataVo, 0, 0)
		json.Unmarshal([]byte(getString), &dictDataList)
	}
	return
}

func SetDictCache(dictType string, dictDataList []*systemModels.SysDictDataVo) {
	redis.SetStruct(constants.SysDictKey+dictType, dictDataList, 0)
}
