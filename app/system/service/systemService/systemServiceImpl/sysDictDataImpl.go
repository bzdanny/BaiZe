package systemServiceImpl

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/utils/exceLize"

	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/systemModels"
	"baize/app/utils/dictUtils"
	"baize/app/utils/snowflake"
)

var dictDataServiceImpl *dictDataService = &dictDataService{dictDataDao: systemDaoImpl.GetSysDictDataDao()}

type dictDataService struct {
	dictDataDao systemDao.IDictDataDao
}

func GetDictDataService() *dictDataService {
	return dictDataServiceImpl
}

func (dictDataService *dictDataService) SelectDictDataByType(dictType string) (sysDictDataList []*systemModels.SysDictDataVo) {

	sysDictDataList = dictUtils.GetDictCache(dictType)
	if sysDictDataList != nil {
		return
	}
	sysDictDataList = dictDataService.dictDataDao.SelectDictDataByType(dictType)
	if len(sysDictDataList) != 0 {
		go dictUtils.SetDictCache(dictType, sysDictDataList)
	}
	return
}
func (dictDataService *dictDataService) SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, count *int64) {
	return dictDataService.dictDataDao.SelectDictDataList(dictData)

}
func (dictDataService *dictDataService) ExportDictData(dictData *systemModels.SysDictDataDQL) (data []byte) {
	list, _ := dictDataService.dictDataDao.SelectDictDataList(dictData)
	rows := systemModels.SysDictDataListToRows(list)
	return exceLize.SetRows(rows)

}
func (dictDataService *dictDataService) SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo) {
	return dictDataService.dictDataDao.SelectDictDataById(dictCode)

}

func (dictDataService *dictDataService) InsertDictData(dictData *systemModels.SysDictDataDML) {
	dictData.DictCode = snowflake.GenID()
	dictDataService.dictDataDao.InsertDictData(dictData)
	redis.Delete(constants.SysDictKey + "*")

}

func (dictDataService *dictDataService) UpdateDictData(dictData *systemModels.SysDictDataDML) {
	dictDataService.dictDataDao.UpdateDictData(dictData)
	redis.Delete(constants.SysDictKey + "*")
}
func (dictDataService *dictDataService) DeleteDictDataByIds(dictCodes []int64) {
	dictDataService.dictDataDao.DeleteDictDataByIds(dictCodes)
	redis.Delete(constants.SysDictKey + "*")
}
func (dictDataService *dictDataService) CheckDictDataByTypes(dictType []string) bool {
	return dictDataService.dictDataDao.CountDictDataByTypes(dictType) > 0

}
