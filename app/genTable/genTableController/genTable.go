package genTableController

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/app/genTable/genTableService"
	"github.com/bzdanny/BaiZe/app/genTable/genTableService/genTableServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"strings"
)

type GenTableController struct {
	gtc genTableService.IGenTableColumnService
	gts genTableService.IGenTableService
}

func NewGenTableController(gtc *genTableServiceImpl.GenTabletColumnService, gts *genTableServiceImpl.GenTabletService) *GenTableController {
	return &GenTableController{gtc: gtc, gts: gts}
}

func (g *GenTableController) GenTableList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	getTable := new(genTableModels.GenTableDQL)
	c.ShouldBind(getTable)
	list, count := g.gts.SelectGenTableList(getTable)

	bzc.SuccessListData(list, count)

}

func (g *GenTableController) GenTableGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	tableId := bzc.ParamInt64("tableId")
	genTable := g.gts.SelectGenTableById(tableId)
	tables := g.gts.SelectGenTableAll()
	list := g.gtc.SelectGenTableColumnListByTableId(tableId)
	data := make(map[string]interface{})
	data["info"] = genTable
	data["rows"] = list
	data["tables"] = tables
	bzc.SuccessData(data)
}

func (g *GenTableController) DataList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	getTable := new(genTableModels.GenTableDQL)
	_ = c.ShouldBind(getTable)
	list, count := g.gts.SelectDbTableList(getTable)

	bzc.SuccessListData(list, count)

}
func (g *GenTableController) ColumnList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	tableId := bzc.ParamInt64("tableId")
	list := g.gtc.SelectGenTableColumnListByTableId(tableId)
	total := int64(len(list))
	bzc.SuccessListData(list, &total)
}
func (g *GenTableController) ImportTable(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	g.gts.ImportTableSave(strings.Split(c.Query("tables"), ","), bzc.GetUserId())
	bzc.SuccessMsg("导入成功")
}
func (g *GenTableController) EditSave(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	genTable := new(genTableModels.GenTableEdit)
	c.ShouldBindJSON(genTable)
	genTable.SetUpdateBy(bzc.GetUserId())
	g.gts.UpdateGenTable(genTable)
	bzc.Success()
}
func (g *GenTableController) GenTableRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	g.gts.DeleteGenTableByIds(bzc.ParamInt64Array("tableIds"))
	bzc.Success()
}
func (g *GenTableController) Preview(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	g.gts.PreviewCode(bzc.ParamInt64("tableId"))
	bzc.Success()
}
func (g *GenTableController) Download(c *gin.Context) {

}

func (g *GenTableController) GenCode(c *gin.Context) {

}

func (g *GenTableController) SynchDb(c *gin.Context) {

}

func (g *GenTableController) BatchGenCode(c *gin.Context) {

}
