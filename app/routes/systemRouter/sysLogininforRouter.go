package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("system:logininfor:list"), systemController.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("system:logininfor:list"), systemController.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.HasPermission("system:logininfor:remove"), systemController.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.HasPermission("system:logininfor:remove"), systemController.LogininforClean)
}
