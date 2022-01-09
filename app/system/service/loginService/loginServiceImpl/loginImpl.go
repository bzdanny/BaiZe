package loginServiceImpl

import (
	"baize/app/common/commonModels"
	"baize/app/constant/userStatus"
	"baize/app/monitor/monitorModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/loginModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"baize/app/utils/IpUtils"
	"baize/app/utils/admin"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/jwt"
	"baize/app/utils/snowflake"
	"baize/app/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	uuid "github.com/satori/go.uuid"
	"time"
)

var instance *loginService

func init() {
	instance = &loginService{
		userDao:          systemDaoImpl.GetSysUserDao(),
		menuDao:          systemDaoImpl.GetSysMenuDao(),
		roleService:      systemServiceImpl.GetRoleService(),
		iLoginforService: monitorServiceImpl.GetLogininforService(),
	}
}

type loginService struct {
	userDao          systemDao.IUserDao
	menuDao          systemDao.IMenuDao
	roleService      systemService.IRoleService
	iLoginforService monitorService.ILogininforService
}

func GetLoginService() *loginService {
	return instance
}

func (loginService *loginService) Login(login *loginModels.LoginBody, c *gin.Context) *commonModels.ResponseData {
	l := new(monitorModels.Logininfor)
	l.UserName = login.Username
	defer loginService.recordLoginInfo(l)
	setUserAgent(l, c)

	captcha := VerityCaptcha(login.Uuid, login.Code)
	if !captcha {
		l.Status = 1
		l.Msg = "验证码错误"
		return commonModels.Waring("验证码错误")
	}
	user := loginService.userDao.SelectUserByUserName(login.Username)
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
	loginUser := new(loginModels.LoginUser)
	loginUser.User = user
	roles := loginService.roleService.SelectBasicRolesByUserId(user.UserId)
	byRoles, loginRoles := loginService.roleService.RolePermissionByRoles(roles)
	loginUser.User.Roles = loginRoles
	loginUser.RolePerms = byRoles
	permission := loginService.getMenuPermission(user.UserId)
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
	go loginService.userDao.UpdateLoginInformation(user.UserId, l.IpAddr)
	return commonModels.SuccessData(tokenStr)
}

func (loginService *loginService) recordLoginInfo(loginUser *monitorModels.Logininfor) {
	go loginService.iLoginforService.InserLogininfor(loginUser)
}

func setUserAgent(login *monitorModels.Logininfor, c *gin.Context) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(c.Request.Header.Get("User-Agent"))
	ip := c.ClientIP()
	login.IpAddr = ip
	login.Os = ua.OS()
	login.LoginLocation = IpUtils.GetRealAddressByIP(ip)
	login.Browser, _ = ua.Browser()

}

func (loginService *loginService) getMenuPermission(userId int64) []string {
	perms := make([]string, 0, 1)
	if admin.IsAdmin(userId) {
		perms = append(perms, "*:*:*")
	} else {
		mysqlPerms := loginService.menuDao.SelectMenuPermsByUserId(userId)

		for _, perm := range mysqlPerms {
			if len(perm) != 0 {
				perms = append(perms, perm)
			}
		}
	}
	return perms
}
