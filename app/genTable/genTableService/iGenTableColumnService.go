package genTableService

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
)

type IGenTableColumnService interface {
	SelectGenTableColumnListByTableId(tableId int64) (list []*genTableModels.GenTableColumnVo)
}
