package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/MenuController"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup) {
	menu := router.Group("/system/menu")
	menu.GET("/list", middlewares.HasPermission("system:menu:list"), MenuController.MenuList)
	menu.GET("/:menuId", middlewares.HasPermission("system:menu:query"), MenuController.MenuGetInfo)
	menu.GET("/treeselect", MenuController.MenuTreeSelect)
	menu.POST("", middlewares.HasPermission("system:menu:add"), MenuController.MenuAdd)
	menu.PUT("", middlewares.HasPermission("system:menu:edit"), MenuController.MenuEdit)
	menu.DELETE("/:menuId", middlewares.HasPermission("system:menu:remove"), MenuController.MenuRemove)
	menu.GET("/roleMenuTreeselect/:roleId", MenuController.RoleMenuTreeselect)
}
