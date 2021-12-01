package monitor

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/conroller/logininforController"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("system:logininfor:list"), logininforController.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("system:logininfor:list"), logininforController.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.HasPermission("system:logininfor:remove"), logininforController.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.HasPermission("system:logininfor:remove"), logininforController.LogininforClean)
}
