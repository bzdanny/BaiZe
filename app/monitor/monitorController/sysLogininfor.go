package monitorController

import (
	"baize/app/common/commonModels"
	"baize/app/monitor/monitorModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var iLoginfor monitorService.ILogininforService = monitorServiceImpl.GetLogininforService()

func LogininforList(c *gin.Context) {
	loginfor := new(monitorModels.LogininforDQL)
	c.ShouldBind(loginfor)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(loginfor)
	loginfor.SetLimit(page)

	list, count := iLoginfor.SelectLogininforList(loginfor)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func LogininforExport(c *gin.Context) {

}

func LogininforRemove(c *gin.Context) {
	var s slicesUtils.Slices = strings.Split(c.Param("infoIds"), ",")
	iLoginfor.DeleteLogininforByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())

}

func LogininforClean(c *gin.Context) {
	iLoginfor.CleanLogininfor()
	c.JSON(http.StatusOK, commonModels.Success())

}
