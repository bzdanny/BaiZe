package systemRouter

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysPostRouter(router *gin.RouterGroup, postController *systemController.PostController) {
	systemPost := router.Group("/system/post")
	systemPost.GET("/list", middlewares.HasPermission("system:post:list"), postController.PostList)
	systemPost.GET("/export", middlewares.HasPermission("system:post:export"), postController.PostExport)
	systemPost.GET("/:postId", middlewares.HasPermission("system:post:query"), postController.PostGetInfo)
	systemPost.POST("", middlewares.HasPermission("system:post:add"), postController.PostAdd)
	systemPost.PUT("", middlewares.HasPermission("system:post:edit"), postController.PostEdit)
	systemPost.DELETE("/:postIds", middlewares.HasPermission("system:post:remove"), postController.PostRemove)

}
