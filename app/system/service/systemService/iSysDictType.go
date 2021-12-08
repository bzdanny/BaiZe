package systemService

import (
	"baize/app/system/models/systemModels"
)

type IDictTypeService interface {
	SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, count *int64)
	SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(dictId []int64) (dictTypes []string)
	InsertDictType(dictType *systemModels.SysDictTypeDML)
	UpdateDictType(dictType *systemModels.SysDictTypeDML)
	DeleteDictTypeByIds(dictIds []int64)
	CheckDictTypeUnique(dictType *systemModels.SysDictTypeDML) bool
	DictTypeClearCache()
	SelectDictTypeAll() (list []*systemModels.SysDictTypeVo)
}
