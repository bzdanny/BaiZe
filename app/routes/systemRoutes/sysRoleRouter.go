package systemRoutes

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysRoleRouter(router *gin.RouterGroup, roleController *systemController.RoleController) {
	role := router.Group("/system/role")
	role.GET("/list", middlewares.HasPermission("system:role:list"), roleController.RoleList)
	role.POST("/export", middlewares.HasPermission("system:role:export"), roleController.RoleExport)
	role.GET("/:roleId", middlewares.HasPermission("system:role:query"), roleController.RoleGetInfo)
	role.POST("", middlewares.HasPermission("system:role:add"), roleController.RoleAdd)
	role.PUT("", middlewares.HasPermission("system:role:edit"), roleController.RoleEdit)
	role.PUT("/dataScope", middlewares.HasPermission("system:role:edit"), roleController.RoleDataScope)
	role.PUT("/changeStatus", middlewares.HasPermission("system:role:edit"), roleController.RoleChangeStatus)
	role.DELETE("/:rolesIds", middlewares.HasPermission("system:role:remove"), roleController.RoleRemove)
	role.GET("/authUser/allocatedList", middlewares.HasPermission("system:role:list"), roleController.AllocatedList)
	role.GET("/authUser/unallocatedList", middlewares.HasPermission("system:role:list"), roleController.UnallocatedList)
	role.PUT("/authUser/selectAll", middlewares.HasPermission("system:role:edit"), roleController.InsertAuthUser)
	role.PUT("/authUser/cancelAll", middlewares.HasPermission("system:role:edit"), roleController.CancelAuthUserAll)
	role.PUT("/authUser/cancel", middlewares.HasPermission("system:role:edit"), roleController.CancelAuthUser)

}
