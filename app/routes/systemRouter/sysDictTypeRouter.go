package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/DictTypeController"
	"github.com/gin-gonic/gin"
)

func InitSysDictTypeRouter(router *gin.RouterGroup) {
	systemDictType := router.Group("/system/dict/type")
	systemDictType.GET("/list", middlewares.HasPermission("system:dict:list"), DictTypeController.DictTypeList)
	systemDictType.POST("/export", middlewares.HasPermission("system:dict:export"), DictTypeController.DictTypeExport)
	systemDictType.GET("/:dictId", middlewares.HasPermission("system:dict:query"), DictTypeController.DictTypeGetInfo)
	systemDictType.POST("", middlewares.HasPermission("system:dict:add"), DictTypeController.DictTypeAdd)
	systemDictType.PUT("", middlewares.HasPermission("system:dict:edit"), DictTypeController.DictTypeEdit)
	systemDictType.DELETE("/:dictIds", middlewares.HasPermission("system:dict:remove"), DictTypeController.DictTypeRemove)
	systemDictType.DELETE("/clearCache", middlewares.HasPermission("system:dict:remove"), DictTypeController.DictTypeClearCache)
	systemDictType.GET("/optionselect", DictTypeController.DictTypeOptionselect)

}
