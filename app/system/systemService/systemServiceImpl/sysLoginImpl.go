package systemServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/app/system/systemDao"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/utils"
	"github.com/bzdanny/BaiZe/app/utils/jwt"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/token"
	"github.com/mojocn/base64Captcha"
	uuid "github.com/satori/go.uuid"
	"image/color"
	"time"
)

type LoginService struct {
	data    *datasource.Data
	userDao systemDao.IUserDao
	menuDao systemDao.IMenuDao
	roleDao systemDao.IRoleDao
	//iLoginforService monitorService.ILogininforService
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (loginService *LoginService) Login(user *systemModels.User, l *monitorModels.Logininfor) *string {
	l.Status = 0
	l.Msg = "登录成功"
	loginUser := new(systemModels.LoginUser)
	loginUser.User = user
	//roles := loginService.roleDao.SelectBasicRolesByUserId(loginService.data.GetSlaveDb(),user.UserId)
	//byRoles, loginRoles := loginService.roleDao.RolePermissionByRoles(roles)
	//loginUser.User.Roles = loginRoles
	//loginUser.RolePerms = byRoles
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
	token.RefreshToken(loginUser)
	go loginService.userDao.UpdateLoginInformation(loginService.data.GetMasterDb(), user.UserId, l.IpAddr)
	return &tokenStr
}

func (loginService *LoginService) RecordLoginInfo(loginUser *monitorModels.Logininfor) {
	//go loginService.iLoginforService.InserLogininfor(loginService.data.GetMasterDb(), loginUser)
}

func (loginService *LoginService) getMenuPermission(userId int64) []string {
	perms := make([]string, 0, 1)
	if utils.IsAdmin(userId) {
		perms = append(perms, "*:*:*")
	} else {
		mysqlPerms := loginService.menuDao.SelectMenuPermsByUserId(loginService.data.GetMasterDb(), userId)

		for _, perm := range mysqlPerms {
			if len(perm) != 0 {
				perms = append(perms, perm)
			}
		}
	}
	return perms
}

var store = base64Captcha.DefaultMemStore

//生成driver，      高，宽，背景文字的干扰，画线条数，背景颜色的指针，字体
var driver = base64Captcha.NewDriverMath(38, 106, 0, 0, &color.RGBA{0, 0, 0, 0}, nil, []string{"wqy-microhei.ttc"})

func (loginService *LoginService) GenerateCode() (m *systemModels.CaptchaVo) {
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		panic(err)
	}
	m = new(systemModels.CaptchaVo)
	m.Id = id
	m.Img = b64s
	return m
}

func (loginService *LoginService) VerityCaptcha(id, base64 string) bool {
	return store.Verify(id, base64, true)
}
