package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func DictDataList(c *gin.Context) {
	dictData := new(systemModels.SysDictDataDQL)
	c.ShouldBind(dictData)
	dictData.SetLimit(c)
	list, count := iDictData.SelectDictDataList(dictData)
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}
func DictDataExport(c *gin.Context) {
	commonLog.SetLog(c, "字典数据", "EXPORT")
}
func DictDataGetInfo(c *gin.Context) {
	dictCode, err := strconv.ParseInt(c.Param("dictCode"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	dictData := iDictData.SelectDictDataById(dictCode)
	c.JSON(http.StatusOK, commonModels.SuccessData(dictData))
}
func DictDataType(c *gin.Context) {
	sysDictDataList := iDictData.SelectDictDataByType(c.Param("dictType"))
	c.JSON(http.StatusOK, commonModels.SuccessData(sysDictDataList))

}

func DictDataAdd(c *gin.Context) {
	commonLog.SetLog(c, "字典数据", "INSERT")
	loginUser := commonController.GetCurrentLoginUser(c)
	dictData := new(systemModels.SysDictDataDML)
	c.ShouldBind(dictData)
	dictData.SetCreateBy(loginUser.User.UserName)
	iDictData.InsertDictData(dictData)

	c.JSON(http.StatusOK, commonModels.Success())
}
func DictDataEdit(c *gin.Context) {
	commonLog.SetLog(c, "字典数据", "UPDATE")
	loginUser := commonController.GetCurrentLoginUser(c)
	dictData := new(systemModels.SysDictDataDML)
	c.ShouldBind(dictData)
	dictData.SetCreateBy(loginUser.User.UserName)
	iDictData.UpdateDictData(dictData)
	c.JSON(http.StatusOK, commonModels.Success())
}
func DictDataRemove(c *gin.Context) {
	commonLog.SetLog(c, "字典数据", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("dictCodes"), ",")
	iDictData.DeleteDictDataByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
}
