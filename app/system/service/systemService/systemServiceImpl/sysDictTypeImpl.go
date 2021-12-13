package systemServiceImpl

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/utils/exceLize"

	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/systemModels"
	"baize/app/utils/snowflake"
)

var dictTypeServiceImpl *dictTypeService = &dictTypeService{dictTypeDao: systemDaoImpl.GetSysDictTypeDao()}

type dictTypeService struct {
	dictTypeDao systemDao.IDictTypeDao
}

func GetDictTypeService() *dictTypeService {
	return dictTypeServiceImpl
}

func (dictTypeService *dictTypeService) SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, count *int64) {
	return dictTypeService.dictTypeDao.SelectDictTypeList(dictType)

}
func (dictTypeService *dictTypeService) ExportDictType(dictType *systemModels.SysDictTypeDQL) (data []byte) {
	list, _ := dictTypeService.dictTypeDao.SelectDictTypeList(dictType)
	rows := systemModels.SysDictTypeListToRows(list)
	return exceLize.SetRows(rows)

}

func (dictTypeService *dictTypeService) SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeById(dictId)

}
func (dictTypeService *dictTypeService) SelectDictTypeByIds(dictId []int64) (dictTypes []string) {
	return dictTypeService.dictTypeDao.SelectDictTypeByIds(dictId)
}

func (dictTypeService *dictTypeService) InsertDictType(dictType *systemModels.SysDictTypeDML) {
	dictType.DictId = snowflake.GenID()
	dictTypeService.dictTypeDao.InsertDictType(dictType)
}

func (dictTypeService *dictTypeService) UpdateDictType(dictType *systemModels.SysDictTypeDML) {
	dictTypeService.dictTypeDao.UpdateDictType(dictType)
}
func (dictTypeService *dictTypeService) DeleteDictTypeByIds(dictIds []int64) {
	dictTypeService.dictTypeDao.DeleteDictTypeByIds(dictIds)
}

func (dictTypeService *dictTypeService) CheckDictTypeUnique(dictType *systemModels.SysDictTypeDML) bool {
	dictId := dictTypeService.dictTypeDao.CheckDictTypeUnique(dictType.DictType)
	if dictId == dictType.DictId || dictId == 0 {
		return false
	}
	return true
}
func (dictTypeService *dictTypeService) DictTypeClearCache() {
	redis.Delete(constants.SysDictKey + "*")
}
func (dictTypeService *dictTypeService) SelectDictTypeAll() (list []*systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeAll()
}
