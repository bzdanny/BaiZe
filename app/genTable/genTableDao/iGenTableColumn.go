package genTableDao

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IGenTableColumn interface {
	SelectDbTableColumnsByName(db dataUtil.DB, tableName string) (list []*genTableModels.InformationSchemaColumn)
	SelectGenTableColumnListByTableId(db dataUtil.DB, tableId int64) (list []*genTableModels.GenTableColumnVo)
	BatchInsertGenTableColumn(db dataUtil.DB, genTables []*genTableModels.GenTableColumnDML)
	UpdateGenTableColumn(db dataUtil.DB, column *genTableModels.GenTableColumnDML)
	DeleteGenTableColumnByIds(db dataUtil.DB, ids []int64)
}
