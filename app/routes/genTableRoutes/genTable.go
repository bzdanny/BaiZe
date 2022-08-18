package genTableRoutes

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitGenTableRouter(router *gin.RouterGroup, g *genTableController.GenTableController) {
	genTable := router.Group("/tool/gen")
	genTable.GET("/list", middlewares.HasPermission("tool:gen:list"), g.GenTableList)
	genTable.GET(":tableId", middlewares.HasPermission("tool:gen:query"), g.GenTableGetInfo)
	genTable.GET("/db/list", middlewares.HasPermission("tool:gen:list"), g.DataList)
	genTable.GET("/column/:talbleId", middlewares.HasPermission("tool:gen:list"), g.ColumnList)
	genTable.POST("/importTable", middlewares.HasPermission("tool:gen:list"), g.ImportTable)
	genTable.PUT("", middlewares.HasPermission("tool:gen:edit"), g.EditSave)
	genTable.DELETE("/:tableIds", middlewares.HasPermission("tool:gen:remove"), g.GenTableRemove)
	genTable.GET("/preview/:tableId", middlewares.HasPermission("tool:gen:remove"), g.Preview)
	genTable.GET("/download/:tableName", middlewares.HasPermission("tool:gen:preview"), g.Download)
	genTable.GET("/genCode/:tableName", middlewares.HasPermission("tool:gen:code"), g.GenCode)
	genTable.GET("/synchDb/:tableName", middlewares.HasPermission("tool:gen:edit"), g.SynchDb)
	genTable.GET("/batchGenCode", middlewares.HasPermission("tool:gen:code"), g.BatchGenCode)

}
