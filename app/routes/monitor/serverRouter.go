package monitor

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/monitorController"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(router *gin.RouterGroup) {
	server := router.Group("/monitor/server")
	server.GET("", middlewares.HasPermission("monitor:server:list"), monitorController.GetInfoServer)

}
