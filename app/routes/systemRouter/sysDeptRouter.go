package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/DeptController"
	"github.com/gin-gonic/gin"
)

func InitSysDeptRouter(router *gin.RouterGroup) {
	systemDept := router.Group("/system/dept")
	systemDept.GET("/list", middlewares.HasPermission("system:dept:list"), DeptController.DeptList)
	systemDept.GET("/:deptId", middlewares.HasPermission("system:dept:query"), DeptController.DeptGetInfo)
	systemDept.GET("/roleDeptTreeselect/:roleId", middlewares.HasPermission("system:dept:query"), DeptController.RoleDeptTreeselect)
	systemDept.POST("", middlewares.HasPermission("system:dept:add"), DeptController.DeptAdd)
	systemDept.PUT("", middlewares.HasPermission("system:dept:edit"), DeptController.DeptEdit)
	systemDept.DELETE("/:deptId", middlewares.HasPermission("system:dept:remove"), DeptController.DeptRemove)

}
