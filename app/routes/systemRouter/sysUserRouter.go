package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/UserController"
	"github.com/gin-gonic/gin"
)

func InitSysUserRouter(router *gin.RouterGroup) {
	systemUser := router.Group("/system/user")
	systemUser.GET("/list", middlewares.HasPermission("system:user:list"), UserController.UserList)
	systemUser.GET("/", middlewares.HasPermission("system:user:query"), UserController.UserGetInfo)
	systemUser.GET("/authRole/:userId", middlewares.HasPermission("system:user:edit"), UserController.UserAuthRole)
	systemUser.GET("/:userId", middlewares.HasPermission("system:user:query"), UserController.UserGetInfoById)
	systemUser.POST("", middlewares.HasPermission("system:user:add"), UserController.UserAdd)
	systemUser.PUT("", middlewares.HasPermission("system:user:edit"), UserController.UserEdit)
	systemUser.PUT("/resetPwd", middlewares.HasPermission("system:user:edit"), UserController.ResetPwd)
	systemUser.PUT("/changeStatus", middlewares.HasPermission("system:user:edit"), UserController.ChangeStatus)
	systemUser.DELETE("/:userIds", middlewares.HasPermission("system:user:remove"), UserController.UserRemove)
	systemUser.POST("/importData", middlewares.HasPermission("system:user:import"), UserController.UserImportData)
	systemUser.GET("/export", middlewares.HasPermission("system:user:export"), UserController.UserExport)
	systemUser.PUT("/authRole", middlewares.HasPermission("system:user:edit"), UserController.InsertAuthRole)
}
