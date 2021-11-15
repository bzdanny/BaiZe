package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/constant/constants"
	"baize/app/system/models/systemModels"
	"baize/app/utils/fileUploadUtils"
	"baize/app/utils/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	m := make(map[string]interface{})
	m["user"] = loginUser.User
	fmt.Println(loginUser.User.CreateTime)
	m["roleGroup"] = iRole.SelectUserRoleGroupByUserId(loginUser.User.UserId)
	m["postGroup"] = iPost.SelectUserPostGroupByUserId(loginUser.User.UserId)
	c.JSON(http.StatusOK, commonModels.SuccessData(m))
}

func ProfileUpdateProfile(c *gin.Context) {
	commonLog.SetLog(c, "个人信息", "UPDATE")
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	if iUser.CheckPhoneUnique(sysUser) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，手机号码已存在"))
		return
	}

	if iUser.CheckEmailUnique(sysUser) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，邮箱账号已存在"))
		return
	}
	loginUser := commonController.GetCurrentLoginUser(c)
	user := loginUser.User
	sysUser.UserId = user.UserId
	sysUser.SetUpdateBy(user.UserName)
	iUser.UpdateUserProfile(sysUser)
	user.NickName = sysUser.NickName
	user.Phonenumber = &sysUser.Phonenumber
	user.Email = &sysUser.Email
	user.Sex = sysUser.Sex
	go token.RefreshToken(loginUser)
	c.JSON(http.StatusOK, commonModels.Success())
}

func ProfileUpdatePwd(c *gin.Context) {
	commonLog.SetLog(c, "个人信息", "UPDATE")
	oldPassword := c.Query("oldPassword")

	password := c.Query("newPassword")
	if oldPassword == password {
		c.JSON(http.StatusOK, commonModels.Waring("新密码不能与旧密码相同"))
		return
	}
	userId := commonController.GetCurrentLoginUser(c).User.UserId
	if !iUser.MatchesPassword(oldPassword, userId) {
		c.JSON(http.StatusOK, commonModels.Waring("修改密码失败，旧密码错误"))
		return
	}
	iUser.ResetUserPwd(userId, password)
	c.JSON(http.StatusOK, commonModels.Success())

}

func ProfileAvatar(c *gin.Context) {
	commonLog.SetLog(c, "个人信息", "UPDATE")
	file, err := c.FormFile("avatarfile")
	if err != nil {
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	filename := fileUploadUtils.Upload(constants.AvatarPath, file)
	loginUser := commonController.GetCurrentLoginUser(c)
	avatar := constants.ResourcePrefix + filename
	iUser.UpdateUserAvatar(loginUser.User.UserId, avatar)
	loginUser.User.Avatar = &avatar
	go token.RefreshToken(loginUser)
	c.JSON(http.StatusOK, commonModels.SuccessData(avatar))
}
