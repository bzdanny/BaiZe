package systemController

import (
	"baize/app/common/commonModels"
	"baize/app/system/service/systemService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserOnlineList(c *gin.Context) {
	list, total := systemService.SelectUserOnlineList(c.Query("ipaddr"), c.Query("userName"))
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, total))
}
func ForceLogout(c *gin.Context) {
	systemService.ForceLogout(c.Param("tokenId"))
	c.JSON(http.StatusOK, commonModels.Success())
}
