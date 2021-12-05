package monitorRoutes

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/conroller/userOnlinController"
	"github.com/gin-gonic/gin"
)

func InitSysUserOnlineRouter(router *gin.RouterGroup) {
	online := router.Group("/monitor/online")
	online.GET("/list", middlewares.HasPermission("monitor:online:list"), userOnlinController.UserOnlineList)
	online.DELETE("/:tokenId", middlewares.HasPermission("monitor:online:forceLogout"), userOnlinController.ForceLogout)

}
