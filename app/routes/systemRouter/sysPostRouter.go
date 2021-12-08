package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/PostController"
	"github.com/gin-gonic/gin"
)

func InitSysPostRouter(router *gin.RouterGroup) {
	systemPost := router.Group("/system/post")
	systemPost.GET("/list", middlewares.HasPermission("system:post:list"), PostController.PostList)
	systemPost.GET("/export", middlewares.HasPermission("system:post:export"), PostController.PostExport)
	systemPost.GET("/:postId", middlewares.HasPermission("system:post:query"), PostController.PostGetInfo)
	systemPost.POST("", middlewares.HasPermission("system:post:add"), PostController.PostAdd)
	systemPost.PUT("", middlewares.HasPermission("system:post:edit"), PostController.PostEdit)
	systemPost.DELETE("/:postIds", middlewares.HasPermission("system:post:remove"), PostController.PostRemove)

}
