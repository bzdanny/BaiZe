package genTableService

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
)

type IGenTableService interface {
	SelectGenTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64)
	SelectDbTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64)
	SelectGenTableAll() (list []*genTableModels.GenTableVo)
	SelectGenTableById(id int64) (genTable *genTableModels.GenTableVo)
	ImportTableSave(table []string, userId int64)
	UpdateGenTable(genTable *genTableModels.GenTableEdit) (err error)
	DeleteGenTableByIds(ids []int64) (err error)
	PreviewCode(tableId int64) (genTable *genTableModels.GenTableVo, err error)
}
