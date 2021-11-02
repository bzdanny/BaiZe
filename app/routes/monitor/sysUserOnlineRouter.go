package monitor

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/monitorController"
	"github.com/gin-gonic/gin"
)

func InitSysUserOnlineRouter(router *gin.RouterGroup) {
	online := router.Group("/monitor/online")
	online.GET("/list", middlewares.HasPermission("monitor:online:list"), monitorController.UserOnlineList)
	online.DELETE("/:tokenId", middlewares.HasPermission("monitor:online:forceLogout"), monitorController.ForceLogout)

}
