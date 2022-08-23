package monitorRoutes

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(router *gin.RouterGroup, sc *monitorController.InfoServerController) {
	server := router.Group("/monitor/server")
	server.GET("", middlewares.HasPermission("monitor:server:list"), sc.GetInfoServer)

}
