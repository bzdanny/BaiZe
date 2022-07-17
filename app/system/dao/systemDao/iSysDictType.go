package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IDictTypeDao interface {
	SelectDictTypeList(db dataUtil.DB, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total *int64)
	SelectDictTypeAll(db dataUtil.DB) (list []*systemModels.SysDictTypeVo)
	SelectDictTypeById(db dataUtil.DB, dictId int64) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(db dataUtil.DB, dictId []int64) (dictTypes []string)
	InsertDictType(db dataUtil.DB, dictType *systemModels.SysDictTypeAdd)
	UpdateDictType(db dataUtil.DB, dictType *systemModels.SysDictTypeEdit)
	DeleteDictTypeByIds(db dataUtil.DB, dictIds []int64) (err error)
	CheckDictTypeUnique(db dataUtil.DB, dictType string) int64
}
