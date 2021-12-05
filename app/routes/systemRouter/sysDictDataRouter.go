package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/DictDataController"
	"github.com/gin-gonic/gin"
)

func InitSysDictDataRouter(router *gin.RouterGroup) {
	systemDictData := router.Group("/system/dict/data")
	systemDictData.GET("/list", middlewares.HasPermission("system:dict:list"), DictDataController.DictDataList)
	systemDictData.GET("/export", middlewares.HasPermission("system:dict:export"), DictDataController.DictDataExport)
	systemDictData.GET("/:dictCode", middlewares.HasPermission("system:dict:query"), DictDataController.DictDataGetInfo)
	systemDictData.GET("/type/:dictType", DictDataController.DictDataType)
	systemDictData.POST("", middlewares.HasPermission("system:dict:add"), DictDataController.DictDataAdd)
	systemDictData.PUT("", middlewares.HasPermission("system:dict:edit"), DictDataController.DictDataEdit)
	systemDictData.DELETE("/:dictCodes", middlewares.HasPermission("system:dict:remove"), DictDataController.DictDataRemove)

}
