package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IDictTypeDao interface {
	SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total *int64)
	SelectDictTypeAll() (list []*systemModels.SysDictTypeVo)
	SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(dictId []int64) (dictTypes []string)
	InsertDictType(dictType *systemModels.SysDictTypeDML)
	UpdateDictType(dictType *systemModels.SysDictTypeDML)
	DeleteDictTypeByIds(dictIds []int64) (err error)
	CheckDictTypeUnique(dictType string) int64
}
