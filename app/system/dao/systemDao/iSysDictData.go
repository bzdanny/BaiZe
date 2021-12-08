package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IDictDataDao interface {
	SelectDictDataByType(dictType string) (SysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total *int64)
	SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(dictData *systemModels.SysDictDataDML)
	UpdateDictData(dictData *systemModels.SysDictDataDML)
	DeleteDictDataByIds(dictCodes []int64)
	CountDictDataByTypes(dictType []string) int
}
