package systemServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/system/systemDao"
	"github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/exceLize"
	"github.com/bzdanny/BaiZe/baize/utils/snowflake"
)

type DictDataService struct {
	data        *datasource.Data
	dictDataDao systemDao.IDictDataDao
}

func NewDictDataService(data *datasource.Data, ddd *systemDaoImpl.SysDictDataDao) *DictDataService {
	return &DictDataService{
		data:        data,
		dictDataDao: ddd,
	}
}

func (dictDataService *DictDataService) SelectDictDataByType(dictType string) (sysDictDataList []*systemModels.SysDictDataVo) {

	//sysDictDataList = dictUtils.GetDictCache(dictType)
	//if sysDictDataList != nil {
	//	return
	//}
	sysDictDataList = dictDataService.dictDataDao.SelectDictDataByType(dictDataService.data.GetSlaveDb(), dictType)
	//if len(sysDictDataList) != 0 {
	//	go dictUtils.SetDictCache(dictType, sysDictDataList)
	//}
	return
}
func (dictDataService *DictDataService) SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, count *int64) {
	return dictDataService.dictDataDao.SelectDictDataList(dictDataService.data.GetSlaveDb(), dictData)

}
func (dictDataService *DictDataService) ExportDictData(dictData *systemModels.SysDictDataDQL) (data []byte) {
	list, _ := dictDataService.dictDataDao.SelectDictDataList(dictDataService.data.GetSlaveDb(), dictData)
	rows := systemModels.SysDictDataListToRows(list)
	return exceLize.SetRows(rows)

}
func (dictDataService *DictDataService) SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo) {
	return dictDataService.dictDataDao.SelectDictDataById(dictDataService.data.GetSlaveDb(), dictCode)

}

func (dictDataService *DictDataService) InsertDictData(dictData *systemModels.SysDictDataAdd) {
	dictData.DictCode = snowflake.GenID()
	dictDataService.dictDataDao.InsertDictData(dictDataService.data.GetMasterDb(), dictData)
	//redis.Delete(constants.SysDictKey + "*")

}

func (dictDataService *DictDataService) UpdateDictData(dictData *systemModels.SysDictDataEdit) {
	dictDataService.dictDataDao.UpdateDictData(dictDataService.data.GetMasterDb(), dictData)
	//redis.Delete(constants.SysDictKey + "*")
}
func (dictDataService *DictDataService) DeleteDictDataByIds(dictCodes []int64) {
	dictDataService.dictDataDao.DeleteDictDataByIds(dictDataService.data.GetMasterDb(), dictCodes)
	//redis.Delete(constants.SysDictKey + "*")
}
func (dictDataService *DictDataService) CheckDictDataByTypes(dictType []string) bool {
	return dictDataService.dictDataDao.CountDictDataByTypes(dictDataService.data.GetSlaveDb(), dictType) > 0

}
