package monitor

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/conroller/logininforConroller"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("system:logininfor:list"), logininforConroller.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("system:logininfor:list"), logininforConroller.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.HasPermission("system:logininfor:remove"), logininforConroller.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.HasPermission("system:logininfor:remove"), logininforConroller.LogininforClean)
}
