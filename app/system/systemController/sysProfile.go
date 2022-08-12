package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/bzdanny/BaiZe/baize/utils/token"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	rs systemService.IRoleService
	ps systemService.IPostService
	us systemService.IUserService
}

func NewProfileController(rs *systemServiceImpl.RoleService, ps *systemServiceImpl.PostService, us *systemServiceImpl.UserService) *ProfileController {
	return &ProfileController{rs: rs, ps: ps, us: us}
}

func (pc *ProfileController) Profile(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	User := bzc.GetUser()
	m := make(map[string]interface{})
	m["user"] = User
	m["roleGroup"] = pc.rs.SelectUserRoleGroupByUserId(User.UserId)
	m["postGroup"] = pc.ps.SelectUserPostGroupByUserId(User.UserId)
	bzc.SuccessData(m)
}

func (pc *ProfileController) ProfileUpdateProfile(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysUser := new(systemModels.SysUserEdit)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	if pc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
		bzc.Waring("修改失败'" + sysUser.Phonenumber + "'失败，手机号码已存在")
		return
	}

	if pc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
		bzc.Waring("修改失败'" + sysUser.Email + "'失败，邮箱账号已存在")
		return
	}
	loginUser := bzc.GetCurrentUser()
	user := loginUser.User
	sysUser.UserId = user.UserId
	sysUser.SetUpdateBy(user.UserId)
	pc.us.UpdateUserProfile(sysUser)
	user.NickName = sysUser.NickName
	user.Phonenumber = &sysUser.Phonenumber
	user.Email = &sysUser.Email
	user.Sex = sysUser.Sex
	go token.RefreshToken(loginUser)
	bzc.Success()
}

func (pc *ProfileController) ProfileUpdatePwd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	oldPassword := c.Query("oldPassword")
	password := c.Query("newPassword")
	if oldPassword == password {
		bzc.Waring("新密码不能与旧密码相同")
		return
	}
	userId := bzc.GetUserId()
	if !pc.us.MatchesPassword(oldPassword, userId) {
		bzc.Waring("修改密码失败，旧密码错误")
		return
	}
	pc.us.ResetUserPwd(userId, password)
	bzc.Success()

}

func (pc *ProfileController) ProfileAvatar(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	file, err := c.FormFile("avatarfile")
	if err != nil {
		bzc.ParameterError()
		return
	}
	bzc.SuccessData(pc.us.UpdateUserAvatar(bzc.GetCurrentUser(), file))
}
