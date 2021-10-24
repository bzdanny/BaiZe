package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/constant/constants"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
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
	m["roleGroup"] = systemService.SelectUserRoleGroupByUserId(loginUser.User.UserId)
	m["postGroup"] = systemService.SelectUserPostGroupByUserId(loginUser.User.UserId)
	c.JSON(http.StatusOK, commonModels.SuccessData(m))
}

func ProfileUpdateProfile(c *gin.Context) {
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	if systemService.CheckPhoneUnique(sysUser) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，手机号码已存在"))
		return
	}

	if systemService.CheckEmailUnique(sysUser) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，邮箱账号已存在"))
		return
	}
	loginUser := commonController.GetCurrentLoginUser(c)
	user := loginUser.User
	sysUser.UserId = user.UserId
	sysUser.SetUpdateBy(user.UserName)
	systemService.UpdateUserProfile(sysUser)
	user.NickName = sysUser.NickName
	user.Phonenumber = &sysUser.Phonenumber
	user.Email = &sysUser.Email
	user.Sex = sysUser.Sex
	go token.RefreshToken(loginUser)
	c.JSON(http.StatusOK, commonModels.Success())
}

func ProfileUpdatePwd(c *gin.Context) {
	oldPassword := c.Query("oldPassword")

	password := c.Query("newPassword")
	if oldPassword == password {
		c.JSON(http.StatusOK, commonModels.Waring("新密码不能与旧密码相同"))
		return
	}
	userId := commonController.GetCurrentLoginUser(c).User.UserId
	if !systemService.MatchesPassword(oldPassword, userId) {
		c.JSON(http.StatusOK, commonModels.Waring("修改密码失败，旧密码错误"))
		return
	}
	systemService.ResetUserPwd(userId, password)
	c.JSON(http.StatusOK, commonModels.Success())

}

func ProfileAvatar(c *gin.Context) {
	file, err := c.FormFile("avatarfile")
	if err != nil {
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	filename := fileUploadUtils.Upload(constants.AvatarPath, file)
	loginUser := commonController.GetCurrentLoginUser(c)
	avatar := constants.ResourcePrefix + filename
	systemService.UpdateUserAvatar(loginUser.User.UserId, avatar)
	loginUser.User.Avatar = &avatar
	go token.RefreshToken(loginUser)
	c.JSON(http.StatusOK, commonModels.SuccessData(avatar))
}
