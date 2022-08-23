package systemServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/system/systemDao"
	"github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/exceLize"
	"github.com/bzdanny/BaiZe/baize/utils/snowflake"
)

type DictTypeService struct {
	data        *datasource.Data
	dictTypeDao systemDao.IDictTypeDao
}

func NewDictTypeService(data *datasource.Data, dtd *systemDaoImpl.SysDictTypeDao) *DictTypeService {
	return &DictTypeService{
		data:        data,
		dictTypeDao: dtd,
	}
}

func (dictTypeService *DictTypeService) SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, count *int64) {
	return dictTypeService.dictTypeDao.SelectDictTypeList(dictTypeService.data.GetSlaveDb(), dictType)

}
func (dictTypeService *DictTypeService) ExportDictType(dictType *systemModels.SysDictTypeDQL) (data []byte) {
	list, _ := dictTypeService.dictTypeDao.SelectDictTypeList(dictTypeService.data.GetSlaveDb(), dictType)
	rows := systemModels.SysDictTypeListToRows(list)
	return exceLize.SetRows(rows)

}

func (dictTypeService *DictTypeService) SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeById(dictTypeService.data.GetSlaveDb(), dictId)

}
func (dictTypeService *DictTypeService) SelectDictTypeByIds(dictId []int64) (dictTypes []string) {
	return dictTypeService.dictTypeDao.SelectDictTypeByIds(dictTypeService.data.GetSlaveDb(), dictId)
}

func (dictTypeService *DictTypeService) InsertDictType(dictType *systemModels.SysDictTypeAdd) {
	dictType.DictId = snowflake.GenID()
	dictTypeService.dictTypeDao.InsertDictType(dictTypeService.data.GetMasterDb(), dictType)
}

func (dictTypeService *DictTypeService) UpdateDictType(dictType *systemModels.SysDictTypeEdit) {
	dictTypeService.dictTypeDao.UpdateDictType(dictTypeService.data.GetMasterDb(), dictType)
}
func (dictTypeService *DictTypeService) DeleteDictTypeByIds(dictIds []int64) {
	dictTypeService.dictTypeDao.DeleteDictTypeByIds(dictTypeService.data.GetMasterDb(), dictIds)
}

func (dictTypeService *DictTypeService) CheckDictTypeUnique(id int64, dictType string) bool {
	dictId := dictTypeService.dictTypeDao.CheckDictTypeUnique(dictTypeService.data.GetSlaveDb(), dictType)
	if dictId == id || dictId == 0 {
		return false
	}
	return true
}
func (dictTypeService *DictTypeService) DictTypeClearCache() {
	//redis.Delete(constants.SysDictKey + "*")
}
func (dictTypeService *DictTypeService) SelectDictTypeAll() (list []*systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeAll(dictTypeService.data.GetSlaveDb())
}
