package genTableController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/genTable/genTableModels"
	"baize/app/genTable/genTableService"
	"baize/app/genTable/genTableService/genTableServiceImpl"
	"github.com/gin-gonic/gin"
	"strings"
)

var iGenTableColumnService genTableService.IGenTableColumnService = genTableServiceImpl.GetGenTabletColumnService()
var iGenTableService genTableService.IGenTableService = genTableServiceImpl.GetGenTabletService()

func GenTableList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	getTable := new(genTableModels.GenTableDQL)
	c.ShouldBind(getTable)
	getTable.SetLimit(c)
	list, count := iGenTableService.SelectGenTableList(getTable)

	bzc.SuccessListData(list, count)

}

func GenTableGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	tableId := bzc.ParamInt64("tableId")
	genTable := iGenTableService.SelectGenTableById(tableId)
	tables := iGenTableService.SelectGenTableAll()
	list := iGenTableColumnService.SelectGenTableColumnListByTableId(tableId)
	data := make(map[string]interface{})
	data["info"] = genTable
	data["rows"] = list
	data["tables"] = tables
	bzc.SuccessData(data)
}

func DataList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	getTable := new(genTableModels.GenTableDQL)
	c.ShouldBind(getTable)
	getTable.SetLimit(c)
	list, count := iGenTableService.SelectDbTableList(getTable)

	bzc.SuccessListData(list, count)

}
func ColumnList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	tableId := bzc.ParamInt64("tableId")
	list := iGenTableColumnService.SelectGenTableColumnListByTableId(tableId)
	total := int64(len(list))
	bzc.SuccessListData(list, &total)
}
func ImportTable(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	iGenTableService.ImportTableSave(strings.Split(c.Query("tables"), ","), bzc.GetCurrentUserName())
	bzc.SuccessMsg("导入成功")
}
func EditSave(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	genTable := new(genTableModels.GenTableDML)
	c.ShouldBindJSON(genTable)
	genTable.SetUpdateBy(bzc.GetCurrentUserName())
	iGenTableService.UpdateGenTable(genTable)
	bzc.Success()
}
func GenTableRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	iGenTableService.DeleteGenTableByIds(bzc.ParamInt64Array("tableIds"))
	bzc.Success()
}
func Preview(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	iGenTableService.PreviewCode(bzc.ParamInt64("tableId"))
	bzc.Success()
}
func Download(c *gin.Context) {

}

func GenCode(c *gin.Context) {

}

func SynchDb(c *gin.Context) {

}

func BatchGenCode(c *gin.Context) {

}
