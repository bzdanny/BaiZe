package systemController

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/bzdanny/BaiZe/baize/utils/token"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	User := bzc.GetCurrentUser()
	m := make(map[string]interface{})
	m["user"] = User
	m["roleGroup"] = iRole.SelectUserRoleGroupByUserId(User.UserId)
	m["postGroup"] = iPost.SelectUserPostGroupByUserId(User.UserId)
	bzc.SuccessData(m)
}

func ProfileUpdateProfile(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("个人信息", "UPDATE")
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	if iUser.CheckPhoneUnique(sysUser) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，手机号码已存在")
		return
	}

	if iUser.CheckEmailUnique(sysUser) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，邮箱账号已存在")
		return
	}
	loginUser := bzc.GetCurrentLoginUser()
	user := loginUser.User
	sysUser.UserId = user.UserId
	sysUser.SetUpdateBy(user.UserName)
	iUser.UpdateUserProfile(sysUser)
	user.NickName = sysUser.NickName
	user.Phonenumber = &sysUser.Phonenumber
	user.Email = &sysUser.Email
	user.Sex = sysUser.Sex
	go token.RefreshToken(loginUser)
	bzc.Success()
}

func ProfileUpdatePwd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("个人信息", "UPDATE")
	oldPassword := c.Query("oldPassword")
	password := c.Query("newPassword")
	if oldPassword == password {
		bzc.Waring("新密码不能与旧密码相同")
		return
	}
	userId := bzc.GetCurrentUserId()
	if !iUser.MatchesPassword(oldPassword, userId) {
		bzc.Waring("修改密码失败，旧密码错误")
		return
	}
	iUser.ResetUserPwd(userId, password)
	bzc.Success()

}

func ProfileAvatar(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("个人信息", "UPDATE")
	file, err := c.FormFile("avatarfile")
	if err != nil {
		bzc.ParameterError()
		return
	}
	filename := fileUploadUtils.Upload(constants.AvatarPath, file)
	loginUser := bzc.GetCurrentLoginUser()
	avatar := constants.ResourcePrefix + filename
	iUser.UpdateUserAvatar(loginUser.User.UserId, avatar)
	loginUser.User.Avatar = &avatar
	go token.RefreshToken(loginUser)
	bzc.SuccessData(avatar)
}
