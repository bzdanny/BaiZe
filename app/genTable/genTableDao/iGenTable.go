package genTableDao

import (
	"baize/app/genTable/genTableModels"
)

type IGenTable interface {
	SelectGenTableList(GenTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64)
	SelectDbTableList(GenTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64)
	SelectDbTableListByNames(tableNames []string) (list []*genTableModels.DBTableVo)
	SelectGenTableById(id int64) (genTable *genTableModels.GenTableVo)
	SelectGenTableByName(name string) (genTable *genTableModels.GenTableVo)
	SelectGenTableAll() (list []*genTableModels.GenTableVo)
	BatchInsertGenTable(genTables []*genTableModels.GenTableDML)
	InsertGenTable(genTable *genTableModels.GenTableDML)
	UpdateGenTable(genTable *genTableModels.GenTableDML)
	DeleteGenTableByIds(ids []int64)
}
