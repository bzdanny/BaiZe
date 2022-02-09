package userOnlinController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"github.com/gin-gonic/gin"
)

var iUserOnline monitorService.ItUserOnlineService = monitorServiceImpl.GetUserOnlineService()

func UserOnlineList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	list, total := iUserOnline.SelectUserOnlineList(c.Query("ipaddr"), c.Query("userName"))
	bzc.SuccessListData(list, total)
}
func ForceLogout(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("在线用户", "FORCE")
	iUserOnline.ForceLogout(c.Param("tokenId"))
	bzc.Success()
}
