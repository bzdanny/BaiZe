package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/deptController"
	"github.com/gin-gonic/gin"
)

func InitSysDeptRouter(router *gin.RouterGroup) {
	systemDept := router.Group("/system/dept")
	systemDept.GET("/list", middlewares.HasPermission("system:dept:list"), deptController.DeptList)
	systemDept.GET("/:deptId", middlewares.HasPermission("system:dept:query"), deptController.DeptGetInfo)
	systemDept.GET("/roleDeptTreeselect/:roleId", middlewares.HasPermission("system:dept:query"), deptController.RoleDeptTreeselect)
	systemDept.POST("", middlewares.HasPermission("system:dept:add"), deptController.DeptAdd)
	systemDept.PUT("", middlewares.HasPermission("system:dept:edit"), deptController.DeptEdit)
	systemDept.DELETE("/:deptId", middlewares.HasPermission("system:dept:remove"), deptController.DeptRemove)

}
