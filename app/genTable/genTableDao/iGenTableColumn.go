package genTableDao

import (
	"baize/app/genTable/genTableModels"
)

type IGenTableColumn interface {
	SelectDbTableColumnsByName(tableName string) (list []*genTableModels.InformationSchemaColumn)
	SelectGenTableColumnListByTableId(tableId int64) (list []*genTableModels.GenTableColumnVo)
	BatchInsertGenTableColumn(genTables []*genTableModels.GenTableColumnDML)
	UpdateGenTableColumn(column *genTableModels.GenTableColumnDML)
	DeleteGenTableColumnByIds(ids []int64)
}
