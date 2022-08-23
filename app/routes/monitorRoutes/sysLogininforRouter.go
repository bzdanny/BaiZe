package monitorRoutes

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup, lc *monitorController.LogininforController) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("system:logininfor:list"), lc.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("system:logininfor:list"), lc.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.HasPermission("system:logininfor:remove"), lc.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.HasPermission("system:logininfor:remove"), lc.LogininforClean)
}
