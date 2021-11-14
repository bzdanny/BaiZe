package monitor

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/monitorController"
	"github.com/gin-gonic/gin"
)

func InitSysOperLogRouter(router *gin.RouterGroup) {
	operlog := router.Group("/monitor/operlog")
	operlog.GET("/list", middlewares.HasPermission("system:operlog:list"), monitorController.OperLogList)
	operlog.GET("/export", middlewares.HasPermission("system:operlog:list"), monitorController.OperLogExport)
	operlog.DELETE("/:infoIds", middlewares.HasPermission("system:operlog:remove"), monitorController.OperLogRemove)
	operlog.DELETE("/clean", middlewares.HasPermission("system:operlog:remove"), monitorController.OperLogClean)
}
