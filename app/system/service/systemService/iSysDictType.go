package systemService

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
)

type IDictTypeService interface {
	SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, count *int64)
	ExportDictType(dictType *systemModels.SysDictTypeDQL) (data []byte)
	SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(dictId []int64) (dictTypes []string)
	InsertDictType(dictType *systemModels.SysDictTypeAdd)
	UpdateDictType(dictType *systemModels.SysDictTypeEdit)
	DeleteDictTypeByIds(dictIds []int64)
	CheckDictTypeUnique(id int64, dictType string) bool
	DictTypeClearCache()
	SelectDictTypeAll() (list []*systemModels.SysDictTypeVo)
}
