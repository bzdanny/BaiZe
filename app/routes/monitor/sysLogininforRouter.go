package monitor

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/monitorController"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("system:logininfor:list"), monitorController.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("system:logininfor:list"), monitorController.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.HasPermission("system:logininfor:remove"), monitorController.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.HasPermission("system:logininfor:remove"), monitorController.LogininforClean)
}
