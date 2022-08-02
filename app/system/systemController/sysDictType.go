package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictTypeController struct {
	dts systemService.IDictTypeService
}

func NewDictTypeController(dts *systemServiceImpl.DictTypeService) *DictTypeController {
	return &DictTypeController{dts: dts}
}

func (dtc *DictTypeController) DictTypeList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictType := new(systemModels.SysDictTypeDQL)
	_ = c.ShouldBind(dictType)
	list, count := dtc.dts.SelectDictTypeList(dictType)
	bzc.SuccessListData(list, count)

}

func (dtc *DictTypeController) DictTypeExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictType := new(systemModels.SysDictTypeDQL)
	c.ShouldBind(dictType)
	bzc.DataPackageExcel(dtc.dts.ExportDictType(dictType))
}

func (dtc *DictTypeController) DictTypeGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictId := bzc.ParamInt64("dictId")
	if dictId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	dictData := dtc.dts.SelectDictTypeById(dictId)

	bzc.SuccessData(dictData)
}

func (dtc *DictTypeController) DictTypeAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictType := new(systemModels.SysDictTypeAdd)
	if err := c.ShouldBindJSON(dictType); err != nil {
		bzc.ParameterError()
		return
	}
	if dtc.dts.CheckDictTypeUnique(dictType.DictId, dictType.DictType) {
		bzc.Waring("新增字典'" + dictType.DictName + "'失败，字典类型已存在")
		return
	}
	dictType.SetCreateBy(bzc.GetUserId())
	dtc.dts.InsertDictType(dictType)
	bzc.Success()
}

func (dtc *DictTypeController) DictTypeEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictType := new(systemModels.SysDictTypeEdit)
	if err := c.ShouldBindJSON(dictType); err != nil {
		bzc.ParameterError()
		return
	}
	if dtc.dts.CheckDictTypeUnique(dictType.DictId, dictType.DictType) {
		bzc.Waring("修改字典'" + dictType.DictName + "'失败，字典类型已存在")
		return
	}

	dictType.SetUpdateBy(bzc.GetUserId())
	dtc.dts.UpdateDictType(dictType)
	bzc.Success()
}

func (dtc *DictTypeController) DictTypeRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dictIds := bzc.ParamInt64Array("dictIds")
	//dictTypes := dtc.dts.SelectDictTypeByIds(dictIds)
	//if dtc.dts.CheckDictDataByTypes(dictTypes) {
	//	bzc.Waring("有已分配的字典,不能删除")
	//	return
	//}
	dtc.dts.DeleteDictTypeByIds(dictIds)
	bzc.Success()
}

func (dtc *DictTypeController) DictTypeClearCache(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dtc.dts.DictTypeClearCache()
	bzc.Success()
}

func (dtc *DictTypeController) DictTypeOptionselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SuccessData(dtc.dts.SelectDictTypeAll())
}
