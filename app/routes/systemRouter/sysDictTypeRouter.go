package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysDictTypeRouter(router *gin.RouterGroup) {
	systemDictType := router.Group("/system/dict/type")
	systemDictType.GET("/list", middlewares.HasPermission("system:dict:list"), systemController.DictTypeList)
	systemDictType.GET("/export", middlewares.HasPermission("system:dict:export"), systemController.DictTypeExport)
	systemDictType.GET("/:dictId", middlewares.HasPermission("system:dict:query"), systemController.DictTypeGetInfo)
	systemDictType.POST("", middlewares.HasPermission("system:dict:add"), systemController.DictTypeAdd)
	systemDictType.PUT("", middlewares.HasPermission("system:dict:edit"), systemController.DictTypeEdit)
	systemDictType.DELETE("/:dictIds", middlewares.HasPermission("system:dict:remove"), systemController.DictTypeRemove)
	systemDictType.DELETE("/clearCache", middlewares.HasPermission("system:dict:remove"), systemController.DictTypeClearCache)
	systemDictType.GET("/optionselect", systemController.DictTypeOptionselect)

}
