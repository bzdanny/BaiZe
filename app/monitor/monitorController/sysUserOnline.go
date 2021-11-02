package monitorController

import (
	"baize/app/common/commonModels"
	"baize/app/monitor/monitorService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserOnlineList(c *gin.Context) {
	list, total := monitorService.SelectUserOnlineList(c.Query("ipaddr"), c.Query("userName"))
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, total))
}
func ForceLogout(c *gin.Context) {
	monitorService.ForceLogout(c.Param("tokenId"))
	c.JSON(http.StatusOK, commonModels.Success())
}
