package UserController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"baize/app/utils/admin"
	"baize/app/utils/slicesUtils"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

var iUser systemService.IUserService = systemServiceImpl.GetUserService()
var iPost systemService.IPostService = systemServiceImpl.GetPostService()
var iRole systemService.IRoleService = systemServiceImpl.GetRoleService()

func ChangeStatus(c *gin.Context) {
	commonLog.SetLog(c, "用户管理", "UPDATE")
	loginUser := commonController.GetCurrentLoginUser(c)
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	sysUser.SetUpdateBy(loginUser.User.UserName)
	iUser.UpdateuserStatus(sysUser)
	c.JSON(http.StatusOK, commonModels.Success())
}
func ResetPwd(c *gin.Context) {
	commonLog.SetLog(c, "用户管理", "UPDATE")
	loginUser := commonController.GetCurrentLoginUser(c)
	sysUser := new(systemModels.SysUserDML)
	c.ShouldBindJSON(sysUser)
	sysUser.SetUpdateBy(loginUser.User.UserName)
	iUser.ResetPwd(sysUser)
	c.JSON(http.StatusOK, commonModels.Success())

}
func UserEdit(c *gin.Context) {
	commonLog.SetLog(c, "用户管理", "UPDATE")
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
	commonLog.SetLog(c, "用户管理", "INSERT")
	loginUser := commonController.GetCurrentLoginUser(c)
	sysUser := new(systemModels.SysUserDML)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	if sysUser.DeptId==nil{
		sysUser.UserId=loginUser.User.UserId
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
	user.SetLimit(c)
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
func UserAuthRole(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
	}
	m := make(map[string]interface{})
	m["user"] = iUser.SelectUserById(userId)
	role := new(systemModels.SysRoleDQL)
	loginUser := commonController.GetCurrentLoginUser(c)
	role.SetDataScope(loginUser, "d", "")
	roles := iRole.SelectRoleAll(role)
	if !admin.IsAdmin(commonController.GetCurrentLoginUser(c).User.UserId) {
		for i, role := range roles {
			if role.RoleId == 1 {
				roles = append(roles[:i], roles[i+1:]...)
				break
			}
		}
	}
	m["roles"] = roles
	m["roleIds"] = slicesUtils.IntSlicesToString(iRole.SelectRoleListByUserId(userId))
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
	commonLog.SetLog(c, "用户管理", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("userIds"), ",")
	iUser.DeleteUserByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
}
func UserImportData(c *gin.Context) {
	commonLog.SetLog(c, "用户管理", "IMPORT")
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
	commonLog.SetLog(c, "用户管理", "EXPORT")
	loginUser := commonController.GetCurrentLoginUser(c)
	user := new(systemModels.SysUserDQL)
	c.ShouldBind(user)
	user.SetDataScope(loginUser, "d", "u")
	data := iUser.UserExport(user)
	c.Header("Content-Type", "application/vnd.ms-excel")
	c.Header("Content-Disposition", "attachment; filename=\"用户管理导出.xls\"")
	c.Header("Pragma", "public")
	c.Header("Cache-Control", "no-store")
	c.Header("Cache-Control", "max-age=0")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	c.Data(http.StatusOK, "application/vnd.ms-excel", data)
	return
}
func InsertAuthRole(c *gin.Context) {
	commonLog.SetLog(c, "用户管理", "GRANT")
	var s slicesUtils.Slices = strings.Split(c.Query("roleIds"), ",")
	userId := c.Query("userId")
	iUser.InsertUserAuth(gconv.Int64(userId),s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
	return
}
