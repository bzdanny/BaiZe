package monitorRoutes

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysOperLogRouter(router *gin.RouterGroup, oper *monitorController.OperLogController) {
	operlog := router.Group("/monitor/operlog")
	operlog.GET("/list", middlewares.HasPermission("system:operlog:list"), oper.OperLogList)
	operlog.GET("/export", middlewares.HasPermission("system:operlog:list"), oper.OperLogExport)
	operlog.DELETE("/:operIds", middlewares.HasPermission("system:operlog:remove"), oper.OperLogRemove)
	operlog.DELETE("/clean", middlewares.HasPermission("system:operlog:remove"), oper.OperLogClean)
}
