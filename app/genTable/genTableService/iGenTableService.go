package genTableService

import (
	"baize/app/genTable/genTableModels"
)

type IGenTableService interface {
	SelectGenTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64)
	SelectDbTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64)
	SelectGenTableAll() (list []*genTableModels.GenTableVo)
	SelectGenTableById(id int64) (genTable *genTableModels.GenTableVo)
	ImportTableSave(table []string, userName string)
	UpdateGenTable(genTable *genTableModels.GenTableDML) (err error)
	DeleteGenTableByIds(ids []int64) (err error)
	PreviewCode(tableId int64) (dataMap map[string]string)
	DownloadCode(tableNames []string) []byte
	GeneratorCode(tableNames []string)
}
