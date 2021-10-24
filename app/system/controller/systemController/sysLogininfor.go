package systemController

import (
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LogininforList(c *gin.Context) {
	loginfor := new(systemModels.LogininforDQL)
	c.ShouldBind(loginfor)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(loginfor)
	loginfor.SetLimit(page)

	list, count := systemService.SelectLogininforList(loginfor)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func LogininforExport(c *gin.Context) {

}

func LogininforRemove(c *gin.Context) {
	var s slicesUtils.Slices = strings.Split(c.Param("infoIds"), ",")
	systemService.DeleteLogininforByIds(s.StrSlicesToInt())

}

func LogininforClean(c *gin.Context) {
	systemService.CleanLogininfor()
	c.JSON(http.StatusOK, commonModels.Success())

}
