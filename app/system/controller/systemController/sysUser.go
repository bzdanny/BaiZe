package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/admin"
	"baize/app/utils/slicesUtils"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func ChangeStatus(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	sysUser.SetUpdateBy(loginUser.User.UserName)
	iUser.UpdateuserStatus(sysUser)
	c.JSON(http.StatusOK, commonModels.Success())
}
func ResetPwd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	sysUser.SetUpdateBy(loginUser.User.UserName)
	iUser.ResetPwd(sysUser)
	c.JSON(http.StatusOK, commonModels.Success())
}
func UserEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
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

	sysUser.SetUpdateBy(loginUser.User.UserName)
	iUser.UpdateUser(sysUser)
	c.JSON(http.StatusOK, commonModels.Success())
}

func UserAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysUser := new(systemModels.SysUserDML)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	if iUser.CheckUserNameUnique(sysUser.UserName) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，登录账号已存在"))
		return
	}
	if iUser.CheckPhoneUnique(sysUser) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，手机号码已存在"))
		return
	}

	if iUser.CheckEmailUnique(sysUser) {
		c.JSON(http.StatusOK, commonModels.Waring("新增用户'"+sysUser.UserName+"'失败，邮箱账号已存在"))
		return
	}

	sysUser.SetCreateBy(loginUser.User.UserName)
	iUser.InsertUser(sysUser)

	c.JSON(http.StatusOK, commonModels.Success())
}
func UserList(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	user := new(systemModels.SysUserDQL)
	c.ShouldBind(user)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(page)
	user.SetLimit(page)
	user.SetDataScope(loginUser, "d", "u")
	list, count := iUser.SelectUserList(user)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}
func UserGetInfo(c *gin.Context) {
	m := make(map[string]interface{})
	m["posts"] = iPost.SelectPostAll()
	loginUser := commonController.GetCurrentLoginUser(c)
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetDataScope(loginUser, "d", "")
	roleList := iRole.SelectRoleAll(role)
	if !admin.IsAdmin(commonController.GetCurrentLoginUser(c).User.UserId) {
		for i, role := range roleList {
			if role.RoleId == 1 {
				roleList = append(roleList[:i], roleList[i+1:]...)
				break
			}
		}
	}
	m["roles"] = roleList
	c.JSON(http.StatusOK, commonModels.SuccessData(m))

}
func UserGetInfoById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
	}
	m := make(map[string]interface{})
	postList := iPost.SelectPostAll()
	m["posts"] = postList
	loginUser := commonController.GetCurrentLoginUser(c)
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetDataScope(loginUser, "d", "")
	roleList := iRole.SelectRoleAll(role)

	if !admin.IsAdmin(commonController.GetCurrentLoginUser(c).User.UserId) {
		for i, role := range roleList {
			if role.RoleId == 1 {
				roleList = append(roleList[:i], roleList[i+1:]...)
				break
			}
		}
	}
	m["roles"] = roleList
	m["postIds"] = slicesUtils.IntSlicesToString(iPost.SelectPostListByUserId(userId))
	m["roleIds"] = slicesUtils.IntSlicesToString(iRole.SelectRoleListByUserId(userId))
	m["sysUser"] = iUser.SelectUserById(userId)
	c.JSON(http.StatusOK, commonModels.SuccessData(m))

}

func UserRemove(c *gin.Context) {
	var s slicesUtils.Slices = strings.Split(c.Param("userIds"), ",")
	toInt := s.StrSlicesToInt()

	iUser.DeleteUserByIds(toInt)
	c.JSON(http.StatusOK, commonModels.Success())
}
func UserImportData(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, commonModels.Error())
	}
	file, _ := fileHeader.Open()
	defer file.Close()
	excelFile, _ := excelize.OpenReader(file)
	rows := excelFile.GetRows("Sheet1")
	loginUser := commonController.GetCurrentLoginUser(c)
	data, num := iUser.UserImportData(rows, loginUser.User.UserName, loginUser.User.DeptId)
	if num > 0 {
		c.JSON(http.StatusOK, commonModels.ErrorMsg(data))
		return
	}
	c.JSON(http.StatusOK, commonModels.SuccessMsg(data))
}

func UserExport(c *gin.Context) {

	loginUser := commonController.GetCurrentLoginUser(c)
	user := new(systemModels.SysUserDQL)
	c.ShouldBind(user)
	user.SetDataScope(loginUser, "d", "u")
	data := iUser.UserExport(user)
	c.Header("Content-Type", "application/vnd.ms-excel")
	c.Header("Content-Disposition", "attachment; filename=\"用户管理导出.xls\"")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	c.Data(http.StatusOK, "application/vnd.ms-excel", data)
	return
}
