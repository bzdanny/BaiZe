package systemController

import (
	commonController "baize/app/common/commonController"
	"baize/app/common/commonLog"
	commonModels "baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func DictTypeList(c *gin.Context) {
	dictType := new(systemModels.SysDictTypeDQL)
	c.ShouldBind(dictType)
	dictType.SetLimit(c)
	list, count := iDictType.SelectDictTypeList(dictType)
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func DictTypeExport(c *gin.Context) {
	commonLog.SetLog(c, "字典类型", "EXPORT")
}

func DictTypeGetInfo(c *gin.Context) {
	dictId, err := strconv.ParseInt(c.Param("dictId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	dictData := iDictType.SelectDictTypeById(dictId)

	c.JSON(http.StatusOK, commonModels.SuccessData(dictData))
}

func DictTypeAdd(c *gin.Context) {
	commonLog.SetLog(c, "字典类型", "INSERT")
	loginUser := commonController.GetCurrentLoginUser(c)
	dictType := new(systemModels.SysDictTypeDML)
	c.ShouldBind(dictType)
	if iDictType.CheckDictTypeUnique(dictType) {
		c.JSON(http.StatusOK, commonModels.Waring("新增字典'"+dictType.DictName+"'失败，字典类型已存在"))
		return
	}
	dictType.SetCreateBy(loginUser.User.UserName)
	iDictType.InsertDictType(dictType)

	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeEdit(c *gin.Context) {
	commonLog.SetLog(c, "字典类型", "UPDATE")
	loginUser := commonController.GetCurrentLoginUser(c)
	dictType := new(systemModels.SysDictTypeDML)
	if iDictType.CheckDictTypeUnique(dictType) {
		c.JSON(http.StatusOK, commonModels.Waring("修改字典'"+dictType.DictName+"'失败，字典类型已存在"))
		return
	}
	c.ShouldBind(dictType)
	dictType.SetCreateBy(loginUser.User.UserName)
	iDictType.UpdateDictType(dictType)

	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeRemove(c *gin.Context) {
	commonLog.SetLog(c, "字典类型", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("dictIds"), ",")
	dictIds := s.StrSlicesToInt()
	dictTypes := iDictType.SelectDictTypeByIds(dictIds)
	if iDictData.CheckDictDataByTypes(dictTypes) {
		c.JSON(http.StatusOK, commonModels.Waring("有已分配的字典,不能删除"))
		return
	}
	iDictType.DeleteDictTypeByIds(dictIds)
	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeClearCache(c *gin.Context) {
	commonLog.SetLog(c, "字典类型", "CLEAN")
	iDictType.DictTypeClearCache()
	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeOptionselect(c *gin.Context) {

	c.JSON(http.StatusOK, commonModels.SuccessData(iDictType.SelectDictTypeAll()))
}
