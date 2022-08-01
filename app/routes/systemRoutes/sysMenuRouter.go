package systemRoutes

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup, menuController *systemController.MenuController) {
	menu := router.Group("/system/menu")
	menu.GET("/list", middlewares.HasPermission("system:menu:list"), menuController.MenuList)
	menu.GET("/:menuId", middlewares.HasPermission("system:menu:query"), menuController.MenuGetInfo)
	menu.GET("/treeselect", menuController.MenuTreeSelect)
	menu.POST("", middlewares.HasPermission("system:menu:add"), menuController.MenuAdd)
	menu.PUT("", middlewares.HasPermission("system:menu:edit"), menuController.MenuEdit)
	menu.DELETE("/:menuId", middlewares.HasPermission("system:menu:remove"), menuController.MenuRemove)
	menu.GET("/roleMenuTreeselect/:roleId", menuController.RoleMenuTreeselect)
}
