package systemService

import (
	"baize/app/system/models/systemModels"
)

type IDictDataService interface {
	SelectDictDataByType(dictType string) (sysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, count *int64)
	SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(dictData *systemModels.SysDictDataDML)
	UpdateDictData(dictData *systemModels.SysDictDataDML)
	DeleteDictDataByIds(dictCodes []int64)
	CheckDictDataByTypes(dictType []string) bool
}
