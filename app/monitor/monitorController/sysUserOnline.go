package monitorController

import (
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"github.com/gin-gonic/gin"
	"net/http"
)

var iUserOnline monitorService.ItUserOnlineService = monitorServiceImpl.GetUserOnlineService()

func UserOnlineList(c *gin.Context) {
	list, total := iUserOnline.SelectUserOnlineList(c.Query("ipaddr"), c.Query("userName"))
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, total))
}
func ForceLogout(c *gin.Context) {
	commonLog.SetLog(c, "在线用户", "FORCE")
	iUserOnline.ForceLogout(c.Param("tokenId"))
	c.JSON(http.StatusOK, commonModels.Success())
}
