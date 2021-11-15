package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysPostRouter(router *gin.RouterGroup) {
	systemPost := router.Group("/system/post")
	systemPost.GET("/list", middlewares.HasPermission("system:post:list"), systemController.PostList)
	systemPost.GET("/export", middlewares.HasPermission("system:post:export"), systemController.PostExport)
	systemPost.GET("/:postId", middlewares.HasPermission("system:post:query"), systemController.PostGetInfo)
	systemPost.POST("", middlewares.HasPermission("system:post:add"), systemController.PostAdd)
	systemPost.PUT("", middlewares.HasPermission("system:post:edit"), systemController.PostEdit)
	systemPost.DELETE("/:postIds", middlewares.HasPermission("system:post:remove"), systemController.PostRemove)

}
