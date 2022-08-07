package systemController

import (
	"github.com/bzdanny/BaiZe/app/constant/userStatus"
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/bzdanny/BaiZe/baize/utils/bCryptPasswordEncoder"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginController struct {
	ls systemService.ILoginService
	us systemService.IUserService
}

func NewLoginController(ls *systemServiceImpl.LoginService, us *systemServiceImpl.UserService) *LoginController {
	return &LoginController{ls: ls, us: us}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 管理端登录
// @Param  object body systemModels.LoginBody true "登录信息"
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData "登录成功"
// @Failure 412 {object}  commonModels.ResponseData "参数错误"
// @Failure 500 {object}  commonModels.ResponseData "服务器错误"
// @Failure 601 {object}  commonModels.ResponseData "用户名密码错误"
// @Router /login [post]
func (lc *LoginController) Login(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	var login systemModels.LoginBody
	if err := c.ShouldBindJSON(&login); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		bzc.ParameterError()
		return
	}
	logininfor := new(monitorModels.Logininfor)
	logininfor.UserName = login.Username
	defer lc.ls.RecordLoginInfo(logininfor)
	bzc.SetUserAgent(logininfor)
	captcha := lc.ls.VerityCaptcha(login.Uuid, login.Code)
	if !captcha {
		logininfor.Status = 1
		logininfor.Msg = "验证码错误"
		bzc.Waring("验证码错误")
		return
	}
	user := lc.us.SelectUserByUserName(login.Username)
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
	bzc.SuccessData(lc.ls.Login(user, logininfor))
}
func (lc *LoginController) GetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentUser()
	data := make(map[string]interface{})
	data["user"] = loginUser.User
	data["roles"] = loginUser.RolePerms
	data["permissions"] = loginUser.Permissions
	bzc.SuccessData(data)

}

func (lc *LoginController) Logout(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentUser()
	if loginUser != nil {
		//iUserOnline.ForceLogout(loginUser.Token)
	}
	bzc.Success()
}
func (lc *LoginController) GetCode(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	code := lc.ls.GenerateCode()
	bzc.SuccessData(code)

}
