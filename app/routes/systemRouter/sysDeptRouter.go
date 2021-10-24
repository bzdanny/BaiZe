package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysDeptRouter(router *gin.RouterGroup) {
	systemDept := router.Group("/system/dept")
	systemDept.GET("/list", middlewares.HasPermission("system:dept:list"), systemController.DeptList)
	systemDept.GET("/:deptId", middlewares.HasPermission("system:dept:query"), systemController.DeptGetInfo)
	systemDept.POST("", middlewares.HasPermission("system:dept:add"), systemController.DeptAdd)
	systemDept.PUT("", middlewares.HasPermission("system:dept:edit"), systemController.DeptEdit)
	systemDept.DELETE("/:deptId", middlewares.HasPermission("system:dept:remove"), systemController.DeptRemove)

}
