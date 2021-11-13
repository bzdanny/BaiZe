package genTableService

import (
	"baize/app/genTable/genTableModels"
)

type IGenTableColumnService interface {
	SelectGenTableColumnListByTableId(tableId int64) (list []*genTableModels.GenTableColumnVo)
}
