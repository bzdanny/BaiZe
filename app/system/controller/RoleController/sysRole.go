package RoleController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iRole systemService.IRoleService = systemServiceImpl.GetRoleService()

func RoleList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetLimit(c)
	role.SetDataScope(loginUser, "d", "")
	list, count := iRole.SelectRoleList(role)
	bzc.SuccessListData(list, count)

}

func RoleExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetDataScope(loginUser, "d", "")
	data := iRole.RoleExport(role)
	bzc.DataPackageExcel(data)

}
func RoleGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	sysUser := iRole.SelectRoleById(roleId)
	bzc.SuccessData(sysUser)
}
func RoleAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "INSERT")
	loginUser := bzc.GetCurrentLoginUser()
	sysRole := new(systemModels.SysRoleDML)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		bzc.ParameterError()
		return
	}
	if iRole.CheckRoleNameUnique(sysRole) {
		bzc.Waring("新增角色'" + sysRole.RoleName + "'失败，角色名称已存在")
		return
	}
	if iRole.CheckRoleKeyUnique(sysRole) {
		bzc.Waring("新增角色'" + sysRole.RoleKey + "'失败，角色权限已存在")
		return
	}
	sysRole.SetCreateBy(loginUser.User.UserName)
	iRole.InsertRole(sysRole)
	bzc.Success()

}
func RoleEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "UPDATE")
	loginUser := bzc.GetCurrentLoginUser()
	sysRole := new(systemModels.SysRoleDML)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		bzc.ParameterError()
		return
	}
	if iRole.CheckRoleNameUnique(sysRole) {
		bzc.Waring("新增角色'" + sysRole.RoleName + "'失败，角色名称已存在")
		return
	}
	if iRole.CheckRoleKeyUnique(sysRole) {
		bzc.Waring("新增角色'" + sysRole.RoleKey + "'失败，角色权限已存在")
		return
	}
	sysRole.SetUpdateBy(loginUser.User.UserName)
	iRole.UpdateRole(sysRole)
	bzc.Success()
}
func RoleDataScope(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "UPDATE")
	loginUser := bzc.GetCurrentLoginUser()
	sysRole := new(systemModels.SysRoleDML)
	c.ShouldBindJSON(sysRole)
	sysRole.SetUpdateBy(loginUser.User.UserName)
	iRole.AuthDataScope(sysRole)
	bzc.Success()
}

func RoleChangeStatus(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "UPDATE")
	loginUser := bzc.GetCurrentLoginUser()
	sysRole := new(systemModels.SysRoleDML)
	c.ShouldBindJSON(sysRole)
	sysRole.SetUpdateBy(loginUser.User.UserName)
	iRole.UpdateRoleStatus(sysRole)
	bzc.Success()
}
func RoleRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "DELETE")
	ids := bzc.ParamInt64Array("rolesIds")
	if iRole.CountUserRoleByRoleId(ids) {
		bzc.Waring("角色已分配，不能删除")
		return
	}
	iRole.DeleteRoleByIds(ids)
	bzc.Success()
}
func AllocatedList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	user := new(systemModels.SysRoleAndUserDQL)
	c.ShouldBind(user)
	user.SetLimit(c)
	user.SetDataScope(loginUser, "d", "u")
	list, count := iRole.SelectAllocatedList(user)
	bzc.SuccessListData(list, count)

}
func UnallocatedList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	user := new(systemModels.SysRoleAndUserDQL)
	c.ShouldBind(user)
	user.SetLimit(c)
	user.SetDataScope(loginUser, "d", "u")
	list, count := iRole.SelectUnallocatedList(user)
	bzc.SuccessListData(list, count)
}
func InsertAuthUser(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "GRANT")
	iRole.InsertAuthUsers(bzc.QueryInt64("roleId"), bzc.QueryInt64Array("userIds"))
	bzc.Success()
	return
}
func CancelAuthUser(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "GRANT")
	userRole := new(systemModels.SysUserRole)
	c.ShouldBindJSON(userRole)
	iRole.DeleteAuthUserRole(userRole)
	bzc.Success()
	return
}
func CancelAuthUserAll(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "GRANT")
	iRole.DeleteAuthUsers(bzc.QueryInt64("roleId"), bzc.QueryInt64Array("userIds"))
	bzc.Success()
	return
}
