package genTableDao

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IGenTable interface {
	SelectGenTableList(db dataUtil.DB, GenTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64)
	SelectDbTableList(db dataUtil.DB, GenTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64)
	SelectDbTableListByNames(db dataUtil.DB, tableNames []string) (list []*genTableModels.DBTableVo)
	SelectGenTableById(db dataUtil.DB, id int64) (genTable *genTableModels.GenTableVo)
	SelectGenTableAll(db dataUtil.DB) (list []*genTableModels.GenTableVo)
	BatchInsertGenTable(db dataUtil.DB, genTables []*genTableModels.GenTableAdd)
	InsertGenTable(db dataUtil.DB, genTable *genTableModels.GenTableAdd)
	UpdateGenTable(db dataUtil.DB, genTable *genTableModels.GenTableEdit)
	DeleteGenTableByIds(db dataUtil.DB, ids []int64)
}
