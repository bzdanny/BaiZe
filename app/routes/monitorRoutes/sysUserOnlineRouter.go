package monitorRoutes

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysUserOnlineRouter(router *gin.RouterGroup, uoc *monitorController.UserOnlineController) {
	online := router.Group("/monitor/online")
	online.GET("/list", middlewares.HasPermission("monitor:online:list"), uoc.UserOnlineList)
	online.DELETE("/:tokenId", middlewares.HasPermission("monitor:online:forceLogout"), uoc.ForceLogout)

}
