package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysRoleRouter(router *gin.RouterGroup) {
	role := router.Group("/system/role")
	role.GET("/list", middlewares.HasPermission("system:role:list"), systemController.RoleList)
	role.GET("/export", middlewares.HasPermission("system:role:export"), systemController.RoleExport)
	role.GET("/:roleId", middlewares.HasPermission("system:role:query"), systemController.RoleGetInfo)
	role.POST("", middlewares.HasPermission("system:role:add"), systemController.RoleAdd)
	role.PUT("", middlewares.HasPermission("system:role:edit"), systemController.RoleEdit)
	role.PUT("/dataScope", middlewares.HasPermission("system:role:edit"), systemController.RoleDataScope)
	role.PUT("/changeStatus", middlewares.HasPermission("system:role:edit"), systemController.RoleChangeStatus)
	role.DELETE("/:userIds", middlewares.HasPermission("system:role:remove"), systemController.RoleRemove)
}
