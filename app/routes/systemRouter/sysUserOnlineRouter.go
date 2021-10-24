package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysUserOnlineRouter(router *gin.RouterGroup) {
	online := router.Group("/monitor/online")
	online.GET("/list", middlewares.HasPermission("monitor:online:list"), systemController.UserOnlineList)
	online.DELETE("/:tokenId", middlewares.HasPermission("monitor:online:forceLogout"), systemController.ForceLogout)

}
