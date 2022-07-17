package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IDictDataDao interface {
	SelectDictDataByType(db dataUtil.DB, dictType string) (SysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(db dataUtil.DB, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total *int64)
	SelectDictDataById(db dataUtil.DB, dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(db dataUtil.DB, dictData *systemModels.SysDictDataDML)
	UpdateDictData(db dataUtil.DB, dictData *systemModels.SysDictDataDML)
	DeleteDictDataByIds(db dataUtil.DB, dictCodes []int64)
	CountDictDataByTypes(db dataUtil.DB, dictType []string) int
}
