package monitorRoutes

import (
	"baize/app/common/middlewares"
	"baize/app/monitor/conroller/serverController"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(router *gin.RouterGroup) {
	server := router.Group("/monitor/server")
	server.GET("", middlewares.HasPermission("monitor:server:list"), serverController.GetInfoServer)

}
