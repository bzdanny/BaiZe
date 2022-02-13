package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/userController"
	"github.com/gin-gonic/gin"
)

func InitSysUserRouter(router *gin.RouterGroup) {
	systemUser := router.Group("/system/user")
	systemUser.GET("/list", middlewares.HasPermission("system:user:list"), userController.UserList)
	systemUser.GET("/", middlewares.HasPermission("system:user:query"), userController.UserGetInfo)
	systemUser.GET("/authRole/:userId", middlewares.HasPermission("system:user:edit"), userController.UserAuthRole)
	systemUser.GET("/:userId", middlewares.HasPermission("system:user:query"), userController.UserGetInfoById)
	systemUser.POST("", middlewares.HasPermission("system:user:add"), userController.UserAdd)
	systemUser.PUT("", middlewares.HasPermission("system:user:edit"), userController.UserEdit)
	systemUser.PUT("/resetPwd", middlewares.HasPermission("system:user:edit"), userController.ResetPwd)
	systemUser.PUT("/changeStatus", middlewares.HasPermission("system:user:edit"), userController.ChangeStatus)
	systemUser.DELETE("/:userIds", middlewares.HasPermission("system:user:remove"), userController.UserRemove)
	systemUser.POST("/importData", middlewares.HasPermission("system:user:import"), userController.UserImportData)
	systemUser.POST("/importTemplate", middlewares.HasPermission("system:user:add"), userController.ImportTemplate)
	systemUser.POST("/export", middlewares.HasPermission("system:user:export"), userController.UserExport)
	systemUser.PUT("/authRole", middlewares.HasPermission("system:user:edit"), userController.InsertAuthRole)
}
