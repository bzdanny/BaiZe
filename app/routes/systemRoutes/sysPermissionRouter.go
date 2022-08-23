package systemRoutes

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysPermissionRouter(router *gin.RouterGroup, PermissionController *systemController.PermissionController) {
	Permission := router.Group("/system/permission")
	Permission.GET("/list", middlewares.HasPermission("system:permission:list"), PermissionController.PermissionList)
	Permission.GET("/:PermissionId", middlewares.HasPermission("system:permission:query"), PermissionController.PermissionGetInfo)
	Permission.GET("/treeselect", PermissionController.PermissionTreeSelect)
	Permission.POST("", middlewares.HasPermission("system:permission:add"), PermissionController.PermissionAdd)
	Permission.PUT("", middlewares.HasPermission("system:permission:edit"), PermissionController.PermissionEdit)
	Permission.DELETE("/:PermissionId", middlewares.HasPermission("system:permission:remove"), PermissionController.PermissionRemove)
	Permission.GET("/rolePermissionTreeselect/:roleId", PermissionController.RolePermissionTreeselect)
}
