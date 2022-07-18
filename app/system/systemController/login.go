package systemController

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iUserOnline monitorService.ItUserOnlineService = monitorServiceImpl.GetUserOnlineService()
var iLogin loginService.ILoginService = loginServiceImpl.GetLoginService()
var iMenu systemService.IMenuService = systemServiceImpl.GetMenuService()
var iLogininfor monitorService.ILogininforService = monitorServiceImpl.GetLogininforService()
var iUser systemService.IUserService = systemServiceImpl.GetUserService()

func Login(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	var login loginModels.LoginBody
	if err := c.ShouldBindJSON(&login); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		bzc.ParameterError()
		return
	}
	logininfor := new(monitorModels.Logininfor)
	logininfor.UserName = login.Username
	defer iLogin.RecordLoginInfo(logininfor)
	bzc.SetUserAgent(logininfor)
	captcha := loginServiceImpl.VerityCaptcha(login.Uuid, login.Code)
	if !captcha {
		logininfor.Status = 1
		logininfor.Msg = "验证码错误"
		bzc.Waring("验证码错误")
		return
	}
	user := iUser.SelectUserByUserName(login.Username)
	if user == nil {
		logininfor.Status = 1
		logininfor.Msg = login.Username + " 用户不存在"
		bzc.Waring("用户不存在/密码错误")
		return
	} else if userStatus.Deleted == user.DelFlag {
		logininfor.Status = 1
		logininfor.Msg = login.Username + " 已被删除"
		bzc.Waring("对不起，您的账号：" + login.Username + " 已被删除")
		return
	} else if userStatus.Disable == user.Status {
		logininfor.Status = 1
		logininfor.Msg = login.Username + " 已停用"
		bzc.Waring("对不起，您的账号：" + login.Username + " 已停用")
		return
	} else if !bCryptPasswordEncoder.CheckPasswordHash(login.Password, user.Password) {
		logininfor.Status = 1
		logininfor.Msg = login.Username + "密码错误"
		bzc.Waring("用户不存在/密码错误")
		return
	}
	bzc.SuccessData(iLogin.Login(user, logininfor))
}
func GetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	data := make(map[string]interface{})
	data["user"] = loginUser.User
	data["roles"] = loginUser.RolePerms
	data["permissions"] = loginUser.Permissions
	bzc.SuccessData(data)

}
func GetRouters(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	menus := iMenu.SelectMenuTreeByUserId(bzc.GetCurrentUserId())
	buildMenus := iMenu.BuildMenus(menus)
	bzc.SuccessData(buildMenus)

}
func Logout(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	if loginUser != nil {
		iUserOnline.ForceLogout(loginUser.Token)
	}
	bzc.Success()
}
