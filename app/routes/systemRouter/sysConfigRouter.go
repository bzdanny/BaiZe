package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysConfigRouter(router *gin.RouterGroup) {
	systemConfig := router.Group("/system/config")
	systemConfig.GET("/list", middlewares.HasPermission("system:config:list"), systemController.ConfigList)
	systemConfig.GET("/export", middlewares.HasPermission("system:config:export"), systemController.ConfigExport)
	systemConfig.GET("/:configId", middlewares.HasPermission("system:config:query"), systemController.ConfigGetInfo)
	systemConfig.GET("/configKey/:configKey", middlewares.HasPermission("system:config:query"), systemController.ConfigGetConfigKey)
	systemConfig.POST("", middlewares.HasPermission("system:config:add"), systemController.ConfigAdd)
	systemConfig.PUT("", middlewares.HasPermission("system:config:edit"), systemController.ConfigEdit)
	systemConfig.DELETE("/:configIds", middlewares.HasPermission("system:config:remove"), systemController.ConfigRemove)
	systemConfig.POST("/clearCache", middlewares.HasPermission("system:config:remove"), systemController.ConfigClearCache)

}
