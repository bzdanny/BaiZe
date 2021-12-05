package genTableServiceImpl

import (
	"baize/app/genTable/genTableDao"
	"baize/app/genTable/genTableDao/genTableDaoImpl"
	"baize/app/genTable/genTableModels"
)

var genTabletColumnServiceImpl = &genTabletColumnService{genTabletColumnDao: genTableDaoImpl.GetGenTableColumnDao()}

type genTabletColumnService struct {
	genTabletColumnDao genTableDao.IGenTableColumn
}

func GetGenTabletColumnService() *genTabletColumnService {
	return genTabletColumnServiceImpl
}

func (genTabletColumnService *genTabletColumnService) SelectGenTableColumnListByTableId(tableId int64) (list []*genTableModels.GenTableColumnVo) {

	return genTabletColumnService.genTabletColumnDao.SelectGenTableColumnListByTableId(tableId)
}
