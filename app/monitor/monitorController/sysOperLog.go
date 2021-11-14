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

var iOperLog monitorService.ISysOperLogService = monitorServiceImpl.GetOperLogServiceService()

func OperLogList(c *gin.Context) {
	operLog := new(monitorModels.SysOpenLogDQL)
	c.ShouldBind(operLog)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(page)
	operLog.SetLimit(page)
	list, count := iOperLog.SelectOperLogList(operLog)
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func OperLogExport(c *gin.Context) {

}

func OperLogRemove(c *gin.Context) {
	var s slicesUtils.Slices = strings.Split(c.Param("operIds"), ",")
	iOperLog.DeleteOperLogByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
}

func OperLogClean(c *gin.Context) {
	iOperLog.CleanOperLog()
	c.JSON(http.StatusOK, commonModels.Success())

}
