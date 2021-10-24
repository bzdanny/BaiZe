package systemController

import (
	commonController "baize/app/common/commonController"
	commonModels "baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
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
	var page = commonModels.NewPageDomain()
	c.ShouldBind(page)
	dictType.SetLimit(page)
	list, count := systemService.SelectDictTypeList(dictType)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func DictTypeExport(c *gin.Context) {

}

func DictTypeGetInfo(c *gin.Context) {
	dictId, err := strconv.ParseInt(c.Param("dictId"), 10, 64)
	if err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	dictData := systemService.SelectDictTypeById(dictId)

	c.JSON(http.StatusOK, commonModels.SuccessData(dictData))
}

func DictTypeAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	dictType := new(systemModels.SysDictTypeDML)
	c.ShouldBind(dictType)
	if systemService.CheckDictTypeUnique(dictType) {
		c.JSON(http.StatusOK, commonModels.Waring("新增字典'"+dictType.DictName+"'失败，字典类型已存在"))
		return
	}
	dictType.SetCreateBy(loginUser.User.UserName)
	systemService.InsertDictType(dictType)

	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	dictType := new(systemModels.SysDictTypeDML)
	if systemService.CheckDictTypeUnique(dictType) {
		c.JSON(http.StatusOK, commonModels.Waring("修改字典'"+dictType.DictName+"'失败，字典类型已存在"))
		return
	}
	c.ShouldBind(dictType)
	dictType.SetCreateBy(loginUser.User.UserName)
	systemService.UpdateDictType(dictType)

	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeRemove(c *gin.Context) {
	var s slicesUtils.Slices = strings.Split(c.Param("dictIds"), ",")
	dictIds := s.StrSlicesToInt()
	dictTypes := systemService.SelectDictTypeByIds(dictIds)
	if systemService.CheckDictDataByTypes(dictTypes) {
		c.JSON(http.StatusOK, commonModels.Waring("有已分配的字典,不能删除"))
		return
	}
	systemService.DeleteDictTypeByIds(dictIds)
	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeClearCache(c *gin.Context) {
	systemService.DictTypeClearCache()
	c.JSON(http.StatusOK, commonModels.Success())
}

func DictTypeOptionselect(c *gin.Context) {

	c.JSON(http.StatusOK, commonModels.SuccessData(systemService.SelectDictTypeAll()))
}
