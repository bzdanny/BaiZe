package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup) {
	menu := router.Group("/system/menu")
	menu.GET("/list", middlewares.HasPermission("system:menu:list"), systemController.MenuList)
	menu.GET("/:menuId", middlewares.HasPermission("system:menu:query"), systemController.MenuGetInfo)
	menu.GET("/treeselect", systemController.MenuTreeSelect)
	menu.POST("", middlewares.HasPermission("system:menu:add"), systemController.MenuAdd)
	menu.PUT("", middlewares.HasPermission("system:menu:edit"), systemController.MenuEdit)
	menu.DELETE("/:menuId", middlewares.HasPermission("system:menu:remove"), systemController.MenuRemove)
	menu.GET("/roleMenuTreeselect/:roleId", systemController.RoleMenuTreeselect)
}
