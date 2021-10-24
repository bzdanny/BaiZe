package loginService

import (
	"baize/app/common/commonModels"
	"baize/app/constant/userStatus"
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/jwt"
	"baize/app/utils/snowflake"
	"baize/app/utils/token"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

func Login(login *loginModels.LoginBody, c *gin.Context) *commonModels.ResponseData {
	l := new(systemModels.Logininfor)
	l.InfoId = snowflake.GenID()
	defer recordLoginInfo(l)
	setUserAgent(l, c)

	captcha := VerityCaptcha(login.Uuid, login.Code)
	if !captcha {
		l.Status = 1
		l.Msg = "验证码错误"
		return commonModels.Waring("验证码错误")
	}
	user := systemDao.SelectUserByUserName(login.Username)
	if user == nil {
		return commonModels.Waring("用户不存在/密码错误")
	} else if userStatus.Deleted == user.DelFlag {
		l.Status = 1
		l.Msg = login.Username + " 已被删除"
		return commonModels.Waring("对不起，您的账号：" + login.Username + " 已被删除")
	} else if userStatus.Disable == user.Status {
		l.Status = 1
		l.Msg = login.Username + " 已停用"
		return commonModels.Waring("对不起，您的账号：" + login.Username + " 已停用")
	} else if !bCryptPasswordEncoder.CheckPasswordHash(login.Password, user.Password) {
		l.Status = 1
		l.Msg = login.Username + "密码错误"
		return commonModels.Waring("用户不存在/密码错误")
	}

	l.Status = 0
	l.Msg = "登录成功"
	l.UserName = user.UserName
	loginUser := new(loginModels.LoginUser)
	loginUser.User = user
	roles := systemService.SelectBasicRolesByUserId(user.UserId)
	byRoles, loginRoles := systemService.RolePermissionByRoles(roles)
	loginUser.User.Roles = loginRoles
	loginUser.RolePerms = byRoles
	permission := GetMenuPermission(user.UserId)
	loginUser.Permissions = permission

	loginUser.User.LoginIp = l.IpAddr
	now := time.Now()
	loginUser.Token = uuid.NewV4().String()
	loginUser.LoginLocation = l.LoginLocation
	loginUser.Os = l.Os
	loginUser.Browser = l.Browser
	loginUser.LoginTime = now.Unix()

	tokenStr := jwt.GenToken(loginUser.Token)
	go token.RefreshToken(loginUser)
	go systemService.UpdateLoginInformation(user.UserId, l.IpAddr)
	return commonModels.SuccessData(tokenStr)
}

func recordLoginInfo(loginUser *systemModels.Logininfor) {
	go systemService.InserLogininfor(loginUser)
}
