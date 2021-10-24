package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup) {
	systemUser := router.Group("/system/menu")
	systemUser.GET("/list", middlewares.HasPermission("system:menu:list"), systemController.MenuList)
	systemUser.GET("/:menuId", middlewares.HasPermission("system:menu:query"), systemController.MenuGetInfo)
	systemUser.GET("/treeselect", systemController.MenuTreeSelect)
	systemUser.POST("", middlewares.HasPermission("system:menu:add"), systemController.MenuAdd)
	systemUser.PUT("", middlewares.HasPermission("system:menu:edit"), systemController.MenuEdit)
	systemUser.DELETE("/:menuId", middlewares.HasPermission("system:menu:remove"), systemController.MenuRemove)
}
