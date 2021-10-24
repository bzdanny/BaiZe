package systemService

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
	"baize/app/utils/snowflake"
)

func SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, count *int64) {
	return systemDao.SelectDictTypeList(dictType)

}

func SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo) {
	return systemDao.SelectDictTypeById(dictId)

}
func SelectDictTypeByIds(dictId []int64) (dictTypes []string) {
	return systemDao.SelectDictTypeByIds(dictId)
}

func InsertDictType(dictType *systemModels.SysDictTypeDML) {
	dictType.DictId = snowflake.GenID()
	systemDao.InsertDictType(dictType)
}

func UpdateDictType(dictType *systemModels.SysDictTypeDML) {
	systemDao.UpdateDictType(dictType)
}
func DeleteDictTypeByIds(dictIds []int64) {
	systemDao.DeleteDictTypeByIds(dictIds)
}

func CheckDictTypeUnique(dictType *systemModels.SysDictTypeDML) bool {
	dictId := systemDao.CheckDictTypeUnique(dictType.DictType)
	if dictId == dictType.DictId || dictId == 0 {
		return false
	}
	return true
}
func DictTypeClearCache() {
	redis.Delete(constants.SysDictKey + "*")
}
func SelectDictTypeAll() (list []*systemModels.SysDictTypeVo) {
	return systemDao.SelectDictTypeAll()
}
