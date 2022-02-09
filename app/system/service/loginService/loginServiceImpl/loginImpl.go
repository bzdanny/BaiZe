package loginServiceImpl

import (
	"baize/app/monitor/monitorModels"
	"baize/app/monitor/monitorService"
	"baize/app/monitor/monitorService/monitorServiceImpl"
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/loginModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"baize/app/utils/admin"
	"baize/app/utils/jwt"
	"baize/app/utils/token"
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

func (loginService *loginService) Login(user *loginModels.User, l *monitorModels.Logininfor) *string {
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
	return &tokenStr
}

func (loginService *loginService) RecordLoginInfo(loginUser *monitorModels.Logininfor) {
	go loginService.iLoginforService.InserLogininfor(loginUser)
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
