package genTableRoutes

import (
	"baize/app/common/middlewares"
	"baize/app/genTable/genTableController"
	"github.com/gin-gonic/gin"
)

func InitGenTableRouter(router *gin.RouterGroup) {
	genTable := router.Group("/tool/gen")
	genTable.GET("/list", middlewares.HasPermission("tool:gen:list"), genTableController.GenTableList)
	genTable.GET(":talbleId", middlewares.HasPermission("tool:gen:query"), genTableController.GenTableGetInfo)
	genTable.GET("/db/list", middlewares.HasPermission("tool:gen:list"), genTableController.DataList)
	genTable.GET("/column/:talbleId", middlewares.HasPermission("tool:gen:list"), genTableController.ColumnList)
	genTable.POST("/importTable", middlewares.HasPermission("tool:gen:list"), genTableController.ImportTable)
	genTable.PUT("", middlewares.HasPermission("tool:gen:edit"), genTableController.EditSave)
	genTable.DELETE("/:tableIds", middlewares.HasPermission("tool:gen:remove"), genTableController.GenTableRemove)
	genTable.GET("/preview/:tableId", middlewares.HasPermission("tool:gen:remove"), genTableController.Preview)
	genTable.GET("/download/:tableName", middlewares.HasPermission("tool:gen:preview"), genTableController.Download)
	genTable.GET("/genCode/:tableName", middlewares.HasPermission("tool:gen:code"), genTableController.GenCode)
	genTable.GET("/synchDb/:tableName", middlewares.HasPermission("tool:gen:edit"), genTableController.SynchDb)
	genTable.GET("/batchGenCode", middlewares.HasPermission("tool:gen:code"), genTableController.BatchGenCode)

}
