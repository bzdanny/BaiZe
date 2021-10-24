package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysDictDataRouter(router *gin.RouterGroup) {
	systemDictData := router.Group("/system/dict/data")
	systemDictData.GET("/list", middlewares.HasPermission("system:dict:list"), systemController.DictDataList)
	systemDictData.GET("/export", middlewares.HasPermission("system:dict:export"), systemController.DictDataExport)
	systemDictData.GET("/:dictCode", middlewares.HasPermission("system:dict:query"), systemController.DictDataGetInfo)
	systemDictData.GET("/type/:dictType", systemController.DictDataType)
	systemDictData.POST("", middlewares.HasPermission("system:dict:add"), systemController.DictDataAdd)
	systemDictData.PUT("", middlewares.HasPermission("system:dict:edit"), systemController.DictDataEdit)
	systemDictData.DELETE("/:dictCodes", middlewares.HasPermission("system:dict:remove"), systemController.DictDataRemove)

}
