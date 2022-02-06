package DictTypeController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iDictType systemService.IDictTypeService = systemServiceImpl.GetDictTypeService()
var iDictData systemService.IDictDataService = systemServiceImpl.GetDictDataService()

func DictTypeList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictType := new(systemModels.SysDictTypeDQL)
	c.ShouldBind(dictType)
	dictType.SetLimit(c)
	list, count := iDictType.SelectDictTypeList(dictType)
	bzc.SuccessListData(list, count)

}

func DictTypeExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictType := new(systemModels.SysDictTypeDQL)
	c.ShouldBind(dictType)
	bzc.DataPackageExcel(iDictType.ExportDictType(dictType))
}

func DictTypeGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictId := bzc.ParamInt64("dictId")
	if dictId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	dictData := iDictType.SelectDictTypeById(dictId)

	bzc.SuccessData(dictData)
}

func DictTypeAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典类型", "INSERT")
	loginUser := bzc.GetCurrentLoginUser()
	dictType := new(systemModels.SysDictTypeDML)
	c.ShouldBind(dictType)
	if iDictType.CheckDictTypeUnique(dictType) {
		bzc.Waring("新增字典'" + dictType.DictName + "'失败，字典类型已存在")
		return
	}
	dictType.SetCreateBy(loginUser.User.UserName)
	iDictType.InsertDictType(dictType)
	bzc.Success()
}

func DictTypeEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典类型", "UPDATE")
	loginUser := bzc.GetCurrentLoginUser()
	dictType := new(systemModels.SysDictTypeDML)
	if iDictType.CheckDictTypeUnique(dictType) {
		bzc.Waring("修改字典'" + dictType.DictName + "'失败，字典类型已存在")
		return
	}
	c.ShouldBind(dictType)
	dictType.SetCreateBy(loginUser.User.UserName)
	iDictType.UpdateDictType(dictType)
	bzc.Success()
}

func DictTypeRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典类型", "DELETE")
	dictIds := bzc.ParamInt64Array("dictIds")
	dictTypes := iDictType.SelectDictTypeByIds(dictIds)
	if iDictData.CheckDictDataByTypes(dictTypes) {
		bzc.Waring("有已分配的字典,不能删除")
		return
	}
	iDictType.DeleteDictTypeByIds(dictIds)
	bzc.Success()
}

func DictTypeClearCache(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典类型", "CLEAN")
	iDictType.DictTypeClearCache()
	bzc.Success()
}

func DictTypeOptionselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SuccessData(iDictType.SelectDictTypeAll())
}
