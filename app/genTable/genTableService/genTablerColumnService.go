package genTableService

import (
	"baize/app/genTable/genTableDao"
	"baize/app/genTable/genTableModels"
)

func SelectGenTableColumnListByTableId(tableId int64) (list []*genTableModels.GenTableColumnVo) {

	return genTableDao.SelectGenTableColumnListByTableId(tableId)
}
