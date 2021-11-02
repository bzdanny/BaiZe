package systemController

import (
	commonController "baize/app/common/commonController"
	commonModels "baize/app/common/commonModels"
	"baize/app/monitor/monitorService"
	"baize/app/system/models/loginModels"
	"baize/app/system/service/loginService"
	"baize/app/system/service/systemService"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	var login loginModels.LoginBody
	if err := c.ShouldBindJSON(&login); err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	data := loginService.Login(&login, c)

	c.JSON(http.StatusOK, data)

}
func GetInfo(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)

	data := make(map[string]interface{})
	data["user"] = loginUser.User
	data["roles"] = loginUser.RolePerms
	data["permissions"] = loginUser.Permissions
	c.JSON(http.StatusOK, commonModels.SuccessData(data))

}
func GetRouters(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	menus := systemService.SelectMenuTreeByUserId(loginUser.User.UserId)
	buildMenus := systemService.BuildMenus(menus)
	c.JSON(http.StatusOK, commonModels.SuccessData(buildMenus))

}
func Logout(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	monitorService.ForceLogout(loginUser.Token)
	c.JSON(http.StatusOK, commonModels.Success())

}
