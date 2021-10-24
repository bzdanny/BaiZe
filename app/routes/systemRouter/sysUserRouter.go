package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysUserRouter(router *gin.RouterGroup) {
	systemUser := router.Group("/system/user")
	systemUser.GET("/list", middlewares.HasPermission("system:user:list"), systemController.UserList)
	systemUser.GET("/", middlewares.HasPermission("system:user:query"), systemController.UserGetInfo)
	systemUser.GET("/:userId", middlewares.HasPermission("system:user:query"), systemController.UserGetInfoById)
	systemUser.POST("", middlewares.HasPermission("system:user:add"), systemController.UserAdd)
	systemUser.PUT("", middlewares.HasPermission("system:user:edit"), systemController.UserEdit)
	systemUser.PUT("/resetPwd", middlewares.HasPermission("system:user:edit"), systemController.ResetPwd)
	systemUser.PUT("/changeStatus", middlewares.HasPermission("system:user:edit"), systemController.ChangeStatus)
	systemUser.DELETE("/:userIds", middlewares.HasPermission("system:user:remove"), systemController.UserRemove)
	systemUser.POST("/importData", middlewares.HasPermission("system:user:import"), systemController.UserImportData)
	systemUser.GET("/export", middlewares.HasPermission("system:user:export"), systemController.UserExport)
}
