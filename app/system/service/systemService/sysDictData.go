package systemService

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
	"baize/app/utils/dictUtils"
	"baize/app/utils/snowflake"
)

func SelectDictDataByType(dictType string) (sysDictDataList []*systemModels.SysDictDataVo) {

	sysDictDataList = dictUtils.GetDictCache(dictType)
	if sysDictDataList != nil {
		return
	}
	sysDictDataList = systemDao.SelectDictDataByType(dictType)
	if len(sysDictDataList) != 0 {
		go dictUtils.SetDictCache(dictType, sysDictDataList)
	}
	return
}
func SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, count *int64) {
	return systemDao.SelectDictDataList(dictData)

}

func SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo) {
	return systemDao.SelectDictDataById(dictCode)

}

func InsertDictData(dictData *systemModels.SysDictDataDML) {
	dictData.DictCode = snowflake.GenID()
	systemDao.InsertDictData(dictData)
	redis.Delete(constants.SysDictKey + "*")

}

func UpdateDictData(dictData *systemModels.SysDictDataDML) {
	systemDao.UpdateDictData(dictData)
	redis.Delete(constants.SysDictKey + "*")
}
func DeleteDictDataByIds(dictCodes []int64) {
	systemDao.DeleteDictDataByIds(dictCodes)
	redis.Delete(constants.SysDictKey + "*")
}
func CheckDictDataByTypes(dictType []string) bool {
	return systemDao.CountDictDataByTypes(dictType) > 0

}
