package genTableController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/genTable/genTableModels"
	"baize/app/genTable/genTableService"
	"baize/app/genTable/genTableService/genTableServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var iGenTableColumnService genTableService.IGenTableColumnService = genTableServiceImpl.GetGenTabletColumnService()
var iGenTableService genTableService.IGenTableService = genTableServiceImpl.GetGenTabletService()

func GenTableList(c *gin.Context) {
	getTable := new(genTableModels.GenTableDQL)
	c.ShouldBind(getTable)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(page)
	getTable.SetLimit(page)
	list, count := iGenTableService.SelectGenTableList(getTable)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func GenTableGetInfo(c *gin.Context) {
	talbleId, _ := strconv.ParseInt(c.Param("talbleId"), 10, 64)
	genTable := iGenTableService.SelectGenTableById(talbleId)
	tables := iGenTableService.SelectGenTableAll()
	list := iGenTableColumnService.SelectGenTableColumnListByTableId(talbleId)
	data := make(map[string]interface{})
	data["info"] = genTable
	data["rows"] = list
	data["tables"] = tables
	c.JSON(http.StatusOK, commonModels.SuccessData(data))
}

func DataList(c *gin.Context) {
	getTable := new(genTableModels.GenTableDQL)
	c.ShouldBind(getTable)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(page)
	getTable.SetLimit(page)
	list, count := iGenTableService.SelectDbTableList(getTable)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}
func ColumnList(c *gin.Context) {
	talbleId, _ := strconv.ParseInt(c.Param("talbleId"), 10, 64)
	list := iGenTableColumnService.SelectGenTableColumnListByTableId(talbleId)
	total := int64(len(list))
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, &total))
}
func ImportTable(c *gin.Context) {
	iGenTableService.ImportTableSave(strings.Split(c.Query("tables"), ","), commonController.GetCurrentLoginUser(c).User.UserName)
	c.JSON(http.StatusOK, commonModels.Success())
}
func EditSave(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	genTable := new(genTableModels.GenTableDML)
	c.ShouldBindJSON(genTable)
	genTable.SetUpdateBy(loginUser.User.UserName)
	iGenTableService.UpdateGenTable(genTable)

	c.JSON(http.StatusOK, commonModels.Success())

}
func GenTableRemove(c *gin.Context) {

	var s slicesUtils.Slices = strings.Split(c.Param("tableIds"), ",")
	iGenTableService.DeleteGenTableByIds(s.StrSlicesToInt())

	c.JSON(http.StatusOK, commonModels.Success())
}
func Preview(c *gin.Context) {

	tableId, _ := strconv.ParseInt(c.Param("tableId"), 10, 64)
	iGenTableService.PreviewCode(tableId)

	c.JSON(http.StatusOK, commonModels.Success())
}
func Download(c *gin.Context) {

}

func GenCode(c *gin.Context) {

}

func SynchDb(c *gin.Context) {

}

func BatchGenCode(c *gin.Context) {

}
