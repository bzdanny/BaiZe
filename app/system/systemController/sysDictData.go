package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type DictDataController struct {
	dds systemService.IDictDataService
}

func NewDictDataController(dds *systemServiceImpl.DictDataService) *DictDataController {
	return &DictDataController{dds: dds}
}

func (ddc *DictDataController) DictDataList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictData := new(systemModels.SysDictDataDQL)
	_ = c.ShouldBind(dictData)
	list, count := ddc.dds.SelectDictDataList(dictData)
	bzc.SuccessListData(list, count)

}
func (ddc *DictDataController) DictDataExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictData := new(systemModels.SysDictDataDQL)
	_ = c.ShouldBind(dictData)
	bzc.DataPackageExcel(ddc.dds.ExportDictData(dictData))
}
func (ddc *DictDataController) DictDataGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictCode := bzc.ParamInt64("dictCode")
	if dictCode == 0 {
		bzc.ParameterError()
		return
	}
	dictData := ddc.dds.SelectDictDataById(dictCode)
	bzc.SuccessData(dictData)
}
func (ddc *DictDataController) DictDataType(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysDictDataList := ddc.dds.SelectDictDataByType(c.Param("dictType"))
	bzc.SuccessData(sysDictDataList)
}

func (ddc *DictDataController) DictDataAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictData := new(systemModels.SysDictDataAdd)
	if err := c.ShouldBindJSON(dictData); err != nil {
		bzc.ParameterError()
		return
	}
	dictData.SetCreateBy(bzc.GetUserId())
	ddc.dds.InsertDictData(dictData)
	bzc.Success()
}
func (ddc *DictDataController) DictDataEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictData := new(systemModels.SysDictDataEdit)
	if err := c.ShouldBindJSON(dictData); err != nil {
		bzc.ParameterError()
		return
	}
	dictData.SetUpdateBy(bzc.GetUserId())
	ddc.dds.UpdateDictData(dictData)
	bzc.Success()
}
func (ddc *DictDataController) DictDataRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	ddc.dds.DeleteDictDataByIds(bzc.ParamInt64Array("dictCodes"))
	bzc.Success()
}
