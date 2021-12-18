package monitorRoutes

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/conroller/operLogController"
	"github.com/gin-gonic/gin"
)

func InitSysOperLogRouter(router *gin.RouterGroup) {
	operlog := router.Group("/monitor/operlog")
	operlog.GET("/list", middlewares.HasPermission("system:operlog:list"), operLogController.OperLogList)
	operlog.GET("/export", middlewares.HasPermission("system:operlog:list"), operLogController.OperLogExport)
	operlog.DELETE("/:operIds", middlewares.HasPermission("system:operlog:remove"), operLogController.OperLogRemove)
	operlog.DELETE("/clean", middlewares.HasPermission("system:operlog:remove"), operLogController.OperLogClean)
}
