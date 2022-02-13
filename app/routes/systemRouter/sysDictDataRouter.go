package systemRouter

import (
	"baize/app/common/middlewares"
	"baize/app/system/controller/dictDataController"
	"github.com/gin-gonic/gin"
)

func InitSysDictDataRouter(router *gin.RouterGroup) {
	systemDictData := router.Group("/system/dict/data")
	systemDictData.GET("/list", middlewares.HasPermission("system:dict:list"), dictDataController.DictDataList)
	systemDictData.GET("/export", middlewares.HasPermission("system:dict:export"), dictDataController.DictDataExport)
	systemDictData.GET("/:dictCode", middlewares.HasPermission("system:dict:query"), dictDataController.DictDataGetInfo)
	systemDictData.GET("/type/:dictType", dictDataController.DictDataType)
	systemDictData.POST("", middlewares.HasPermission("system:dict:add"), dictDataController.DictDataAdd)
	systemDictData.PUT("", middlewares.HasPermission("system:dict:edit"), dictDataController.DictDataEdit)
	systemDictData.DELETE("/:dictCodes", middlewares.HasPermission("system:dict:remove"), dictDataController.DictDataRemove)

}
