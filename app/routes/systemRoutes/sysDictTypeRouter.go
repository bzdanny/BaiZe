package systemRoutes

import (
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysDictTypeRouter(router *gin.RouterGroup, dictTypeController *systemController.DictTypeController) {
	systemDictType := router.Group("/system/dict/type")
	systemDictType.GET("/list", middlewares.HasPermission("system:dict:list"), dictTypeController.DictTypeList)
	systemDictType.POST("/export", middlewares.HasPermission("system:dict:export"), dictTypeController.DictTypeExport)
	systemDictType.GET("/:dictId", middlewares.HasPermission("system:dict:query"), dictTypeController.DictTypeGetInfo)
	systemDictType.POST("", middlewares.HasPermission("system:dict:add"), dictTypeController.DictTypeAdd)
	systemDictType.PUT("", middlewares.HasPermission("system:dict:edit"), dictTypeController.DictTypeEdit)
	systemDictType.DELETE("/:dictIds", middlewares.HasPermission("system:dict:remove"), dictTypeController.DictTypeRemove)
	systemDictType.DELETE("/clearCache", middlewares.HasPermission("system:dict:remove"), dictTypeController.DictTypeClearCache)
	systemDictType.GET("/optionselect", dictTypeController.DictTypeOptionselect)

}
