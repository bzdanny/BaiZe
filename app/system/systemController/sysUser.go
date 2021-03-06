package systemController

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	us systemService.IUserService
}

func NewUserController(us *systemServiceImpl.UserService) *UserController {
	return &UserController{us: us}
}

func (uc *UserController) ChangeStatus(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "UPDATE")

	sysUser := new(systemModels.SysUserEdit)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	sysUser.SetUpdateBy(bzc.GetUserId())
	uc.us.UpdateuserStatus(sysUser)
	bzc.Success()
}
func (uc *UserController) ResetPwd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "UPDATE")
	sysUser := new(systemModels.SysUserEdit)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	sysUser.SetUpdateBy(bzc.GetUserId())
	uc.us.ResetPwd(sysUser)
	bzc.Success()

}
func (uc *UserController) UserEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "UPDATE")
	sysUser := new(systemModels.SysUserEdit)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	if uc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
		bzc.Waring("新增用户'" + sysUser.Phonenumber + "'失败，手机号码已存在")
		return
	}
	if uc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
		bzc.Waring("新增用户'" + sysUser.Email + "'失败，邮箱账号已存在")
		return
	}
	sysUser.SetUpdateBy(bzc.GetUserId())
	uc.us.UpdateUser(sysUser)
	bzc.Success()
}

func (uc *UserController) UserAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "INSERT")
	user := bzc.GetUser()
	sysUser := new(systemModels.SysUserAdd)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	if sysUser.DeptId == nil {
		sysUser.DeptId = user.DeptId
	}
	if uc.us.CheckUserNameUnique(sysUser.UserName) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，登录账号已存在")
		return
	}
	if uc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
		bzc.Waring("新增用户'" + sysUser.Phonenumber + "'失败，手机号码已存在")
		return
	}

	if uc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
		bzc.Waring("新增用户'" + sysUser.Email + "'失败，邮箱账号已存在")
		return
	}
	sysUser.SetCreateBy(user.UserId)
	uc.us.InsertUser(sysUser)
	bzc.Success()
}
func (uc *UserController) UserList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysUserDQL)
	_ = c.ShouldBind(user)
	user.SetDataScope(bzc.GetUser(), "d", "u")
	list, count := uc.us.SelectUserList(user)
	bzc.SuccessListData(list, count)

}
func (uc *UserController) UserGetInfo(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//m := make(map[string]interface{})
	//m["posts"] = iPost.SelectPostAll()
	//user := bzc.GetCurrentUser()
	//role := new(systemModels.SysRoleDQL)
	//c.ShouldBind(role)
	//role.SetDataScope(user, "d", "")
	//roleList := iRole.SelectRoleAll(role)
	//if !admin.IsAdmin(user.UserId) {
	//	for i, role := range roleList {
	//		if role.RoleId == 1 {
	//			roleList = append(roleList[:i], roleList[i+1:]...)
	//			break
	//		}
	//	}
	//}
	//m["roles"] = roleList
	//bzc.SuccessData(m)

}
func (uc *UserController) UserAuthRole(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.ParamInt64("userId")
	if userId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	m := make(map[string]interface{})
	m["user"] = uc.us.SelectUserById(userId)
	role := new(systemModels.SysRoleDQL)
	user := bzc.GetUser()
	role.SetDataScope(user, "d", "")
	//roles := iRole.SelectRoleAll(role)
	//if !admin.IsAdmin(user.UserId) {
	//	for i, role := range roles {
	//		if role.RoleId == 1 {
	//			roles = append(roles[:i], roles[i+1:]...)
	//			break
	//		}
	//	}
	//}
	//m["roles"] = roles
	//m["roleIds"] = slicesUtils.IntSlicesToString(iRole.SelectRoleListByUserId(userId))
	bzc.SuccessData(m)
}

func (uc *UserController) UserGetInfoById(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.ParamInt64("userId")
	if userId == 0 {
		bzc.ParameterError()
	}
	m := make(map[string]interface{})
	//postList := iPost.SelectPostAll()
	//m["posts"] = postList
	user := bzc.GetUser()
	role := new(systemModels.SysRoleDQL)
	role.SetDataScope(user, "d", "")
	//roleList := iRole.SelectRoleAll(role)
	//if !admin.IsAdmin(user.UserId) {
	//	for i, role := range roleList {
	//		if role.RoleId == 1 {
	//			roleList = append(roleList[:i], roleList[i+1:]...)
	//			break
	//		}
	//	}
	//}
	//m["roles"] = roleList
	//m["postIds"] = slicesUtils.IntSlicesToString(iPost.SelectPostListByUserId(userId))
	//m["roleIds"] = slicesUtils.IntSlicesToString(iRole.SelectRoleListByUserId(userId))
	//m["sysUser"] = iUser.SelectUserById(userId)
	bzc.SuccessData(m)

}

func (uc *UserController) UserRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "DELETE")
	uc.us.DeleteUserByIds(bzc.ParamInt64Array("userIds"))
	bzc.Success()
}
func (uc *UserController) UserImportData(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "IMPORT")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	file, _ := fileHeader.Open()
	defer file.Close()
	excelFile, _ := excelize.OpenReader(file)
	rows := excelFile.GetRows("Sheet1")
	loginUser := bzc.GetUser()
	data, num := uc.us.UserImportData(rows, loginUser.UserName, loginUser.DeptId)
	if num > 0 {
		bzc.Waring(data)
		return
	}
	bzc.SuccessMsg(data)
}

func (uc *UserController) UserExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysUserDQL)
	c.ShouldBind(user)
	user.SetDataScope(bzc.GetUser(), "d", "u")
	bzc.DataPackageExcel(uc.us.UserExport(user))
	return
}

func (uc *UserController) ImportTemplate(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	data := uc.us.ImportTemplate()
	bzc.DataPackageExcel(data)
	return
}

func (uc *UserController) InsertAuthRole(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("用户管理", "GRANT")
	uc.us.InsertUserAuth(bzc.QueryInt64("userId"), bzc.QueryInt64Array("roleIds"))
	bzc.Success()
	return
}
