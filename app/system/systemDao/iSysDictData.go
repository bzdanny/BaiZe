package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	systemModels2 "github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IDictDataDao interface {
	SelectDictDataByType(db dataUtil.DB, dictType string) (SysDictDataList []*systemModels2.SysDictDataVo)
	SelectDictDataList(db dataUtil.DB, dictData *systemModels2.SysDictDataDQL) (list []*systemModels2.SysDictDataVo, total *int64)
	SelectDictDataById(db dataUtil.DB, dictCode int64) (dictData *systemModels2.SysDictDataVo)
	InsertDictData(db dataUtil.DB, dictData *systemModels.SysDictDataDML)
	UpdateDictData(db dataUtil.DB, dictData *systemModels.SysDictDataDML)
	DeleteDictDataByIds(db dataUtil.DB, dictCodes []int64)
	CountDictDataByTypes(db dataUtil.DB, dictType []string) int
}
