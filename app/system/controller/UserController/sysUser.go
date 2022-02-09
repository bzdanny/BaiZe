package UserController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"baize/app/utils/admin"
	"baize/app/utils/slicesUtils"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iUser systemService.IUserService = systemServiceImpl.GetUserService()
var iPost systemService.IPostService = systemServiceImpl.GetPostService()
var iRole systemService.IRoleService = systemServiceImpl.GetRoleService()

func ChangeStatus(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "UPDATE")

	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	sysUser.SetUpdateBy(bzc.GetCurrentUserName())
	iUser.UpdateuserStatus(sysUser)
	bzc.Success()
}
func ResetPwd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "UPDATE")
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	sysUser.SetUpdateBy(bzc.GetCurrentUserName())
	iUser.ResetPwd(sysUser)
	bzc.Success()

}
func UserEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "UPDATE")
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
	sysUser.SetUpdateBy(bzc.GetCurrentUserName())
	iUser.UpdateUser(sysUser)
	bzc.Success()
}

func UserAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "INSERT")
	user := bzc.GetCurrentUser()
	sysUser := new(systemModels.SysUserDML)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		bzc.ParameterError()
		return
	}
	if sysUser.DeptId == nil {
		sysUser.UserId = user.UserId
	}
	if iUser.CheckUserNameUnique(sysUser.UserName) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，登录账号已存在")
		return
	}
	if iUser.CheckPhoneUnique(sysUser) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，手机号码已存在")
		return
	}

	if iUser.CheckEmailUnique(sysUser) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，邮箱账号已存在")
		return
	}
	sysUser.SetCreateBy(user.UserName)
	iUser.InsertUser(sysUser)
	bzc.Success()
}
func UserList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysUserDQL)
	c.ShouldBind(user)
	user.SetLimit(c)
	user.SetDataScope(bzc.GetCurrentUser(), "d", "u")
	list, count := iUser.SelectUserList(user)
	bzc.SuccessListData(list, count)

}
func UserGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	m := make(map[string]interface{})
	m["posts"] = iPost.SelectPostAll()
	user := bzc.GetCurrentUser()
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetDataScope(user, "d", "")
	roleList := iRole.SelectRoleAll(role)
	if !admin.IsAdmin(user.UserId) {
		for i, role := range roleList {
			if role.RoleId == 1 {
				roleList = append(roleList[:i], roleList[i+1:]...)
				break
			}
		}
	}
	m["roles"] = roleList
	bzc.SuccessData(m)

}
func UserAuthRole(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.ParamInt64("userId")
	if userId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	m := make(map[string]interface{})
	m["user"] = iUser.SelectUserById(userId)
	role := new(systemModels.SysRoleDQL)
	user := bzc.GetCurrentUser()
	role.SetDataScope(user, "d", "")
	roles := iRole.SelectRoleAll(role)
	if !admin.IsAdmin(user.UserId) {
		for i, role := range roles {
			if role.RoleId == 1 {
				roles = append(roles[:i], roles[i+1:]...)
				break
			}
		}
	}
	m["roles"] = roles
	m["roleIds"] = slicesUtils.IntSlicesToString(iRole.SelectRoleListByUserId(userId))
	bzc.SuccessData(m)
}

func UserGetInfoById(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.ParamInt64("userId")
	if userId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	m := make(map[string]interface{})
	postList := iPost.SelectPostAll()
	m["posts"] = postList
	user := bzc.GetCurrentUser()
	role := new(systemModels.SysRoleDQL)
	role.SetDataScope(user, "d", "")
	roleList := iRole.SelectRoleAll(role)
	if !admin.IsAdmin(user.UserId) {
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
	bzc.SuccessData(m)

}

func UserRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "DELETE")
	iUser.DeleteUserByIds(bzc.ParamInt64Array("userIds"))
	bzc.Success()
}
func UserImportData(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "IMPORT")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		bzc.BzError()
	}
	file, _ := fileHeader.Open()
	defer file.Close()
	excelFile, _ := excelize.OpenReader(file)
	rows := excelFile.GetRows("Sheet1")
	loginUser := bzc.GetCurrentUser()
	data, num := iUser.UserImportData(rows, loginUser.UserName, loginUser.DeptId)
	if num > 0 {
		bzc.ErrorMsg(data)
		return
	}
	bzc.SuccessMsg(data)
}

func UserExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysUserDQL)
	c.ShouldBind(user)
	user.SetDataScope(bzc.GetCurrentUser(), "d", "u")
	data := iUser.UserExport(user)
	bzc.DataPackageExcel(data)
	return
}

func ImportTemplate(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	data := iUser.ImportTemplate()
	bzc.DataPackageExcel(data)
	return
}

func InsertAuthRole(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "GRANT")
	iUser.InsertUserAuth(bzc.QueryInt64("userId"), bzc.QueryInt64Array("roleIds"))
	bzc.Success()
	return
}
