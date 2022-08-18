package genTableServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableDao"
	"github.com/bzdanny/BaiZe/app/genTable/genTableDao/genTableDaoImpl"
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
)

type GenTabletColumnService struct {
	data               *datasource.Data
	genTabletColumnDao genTableDao.IGenTableColumn
}

func GetGenTabletColumnService(data *datasource.Data, gtc *genTableDaoImpl.GenTableColumnDao) *GenTabletColumnService {
	return &GenTabletColumnService{data: data, genTabletColumnDao: gtc}
}

func (g *GenTabletColumnService) SelectGenTableColumnListByTableId(tableId int64) (list []*genTableModels.GenTableColumnVo) {

	return g.genTabletColumnDao.SelectGenTableColumnListByTableId(g.data.GetSlaveDb(), tableId)
}
