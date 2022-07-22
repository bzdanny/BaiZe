package systemService

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
)

type IDictDataService interface {
	SelectDictDataByType(dictType string) (sysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, count *int64)
	ExportDictData(dictData *systemModels.SysDictDataDQL) (data []byte)
	SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(dictData *systemModels.SysDictDataAdd)
	UpdateDictData(dictData *systemModels.SysDictDataEdit)
	DeleteDictDataByIds(dictCodes []int64)
	CheckDictDataByTypes(dictType []string) bool
}
