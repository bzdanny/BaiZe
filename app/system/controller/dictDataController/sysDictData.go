package dictDataController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iDictData systemService.IDictDataService = systemServiceImpl.GetDictDataService()

func DictDataList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictData := new(systemModels.SysDictDataDQL)
	c.ShouldBind(dictData)
	dictData.SetLimit(c)
	list, count := iDictData.SelectDictDataList(dictData)
	bzc.SuccessListData(list, count)

}
func DictDataExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictData := new(systemModels.SysDictDataDQL)
	c.ShouldBind(dictData)
	bzc.DataPackageExcel(iDictData.ExportDictData(dictData))
}
func DictDataGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictCode := bzc.ParamInt64("dictCode")
	if dictCode == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	dictData := iDictData.SelectDictDataById(dictCode)
	bzc.SuccessData(dictData)
}
func DictDataType(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysDictDataList := iDictData.SelectDictDataByType(c.Param("dictType"))
	bzc.SuccessData(sysDictDataList)
}

func DictDataAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典数据", "INSERT")
	dictData := new(systemModels.SysDictDataDML)
	c.ShouldBind(dictData)
	dictData.SetCreateBy(bzc.GetCurrentUserName())
	iDictData.InsertDictData(dictData)
	bzc.Success()
}
func DictDataEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典数据", "UPDATE")
	dictData := new(systemModels.SysDictDataDML)
	c.ShouldBind(dictData)
	dictData.SetUpdateBy(bzc.GetCurrentUserName())
	iDictData.UpdateDictData(dictData)
	bzc.Success()
}
func DictDataRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("字典数据", "DELETE")
	iDictData.DeleteDictDataByIds(bzc.ParamInt64Array("dictCodes"))
	bzc.Success()
}
