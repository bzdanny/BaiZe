package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	rs systemService.IRoleService
}

func NewRoleController(rs *systemServiceImpl.RoleService) *RoleController {
	return &RoleController{rs: rs}
}

func (rc *RoleController) RoleList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	role := new(systemModels.SysRoleDQL)
	_ = c.ShouldBind(role)
	role.SetDataScope(bzc.GetUser(), "d", "")
	list, count := rc.rs.SelectRoleList(role)
	bzc.SuccessListData(list, count)

}

func (rc *RoleController) RoleExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	role := new(systemModels.SysRoleDQL)
	_ = c.ShouldBind(role)
	role.SetDataScope(bzc.GetUser(), "d", "")
	bzc.DataPackageExcel(rc.rs.RoleExport(role))

}
func (rc *RoleController) RoleGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		bzc.ParameterError()
		return
	}
	sysUser := rc.rs.SelectRoleById(roleId)
	bzc.SuccessData(sysUser)
}
func (rc *RoleController) RoleAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "INSERT")
	sysRole := new(systemModels.SysRoleAdd)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		bzc.ParameterError()
		return
	}
	if rc.rs.CheckRoleNameUnique(0, sysRole.RoleName) {
		bzc.Waring("新增角色'" + sysRole.RoleName + "'失败，角色名称已存在")
		return
	}
	if rc.rs.CheckRoleKeyUnique(0, sysRole.RoleKey) {
		bzc.Waring("新增角色'" + sysRole.RoleKey + "'失败，角色权限已存在")
		return
	}
	sysRole.SetCreateBy(bzc.GetUserId())
	rc.rs.InsertRole(sysRole)
	bzc.Success()

}
func (rc *RoleController) RoleEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "UPDATE")

	sysRole := new(systemModels.SysRoleEdit)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		bzc.ParameterError()
		return
	}
	if rc.rs.CheckRoleNameUnique(sysRole.RoleId, sysRole.RoleName) {
		bzc.Waring("新增角色'" + sysRole.RoleName + "'失败，角色名称已存在")
		return
	}
	if rc.rs.CheckRoleKeyUnique(sysRole.RoleId, sysRole.RoleKey) {
		bzc.Waring("新增角色'" + sysRole.RoleKey + "'失败，角色权限已存在")
		return
	}
	sysRole.SetUpdateBy(bzc.GetUserId())
	rc.rs.UpdateRole(sysRole)
	bzc.Success()
}
func (rc *RoleController) RoleDataScope(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "UPDATE")
	sysRole := new(systemModels.SysRoleEdit)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		bzc.ParameterError()
		return
	}
	sysRole.SetUpdateBy(bzc.GetUserId())
	rc.rs.AuthDataScope(sysRole)
	bzc.Success()
}

func (rc *RoleController) RoleChangeStatus(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "UPDATE")
	sysRole := new(systemModels.SysRoleEdit)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		bzc.ParameterError()
		return
	}
	sysRole.SetUpdateBy(bzc.GetUserId())
	rc.rs.UpdateRoleStatus(sysRole)
	bzc.Success()
}
func (rc *RoleController) RoleRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "DELETE")
	ids := bzc.ParamInt64Array("rolesIds")
	if rc.rs.CountUserRoleByRoleId(ids) {
		bzc.Waring("角色已分配，不能删除")
		return
	}
	rc.rs.DeleteRoleByIds(ids)
	bzc.Success()
}
func (rc *RoleController) AllocatedList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysRoleAndUserDQL)
	if err := c.ShouldBind(user); err != nil {
		bzc.ParameterError()
		return
	}
	user.SetDataScope(bzc.GetUser(), "d", "u")
	list, count := rc.rs.SelectAllocatedList(user)
	bzc.SuccessListData(list, count)

}
func (rc *RoleController) UnallocatedList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysRoleAndUserDQL)
	if err := c.ShouldBind(user); err != nil {
		bzc.ParameterError()
		return
	}
	user.SetDataScope(bzc.GetUser(), "d", "u")
	list, count := rc.rs.SelectUnallocatedList(user)
	bzc.SuccessListData(list, count)
}
func (rc *RoleController) InsertAuthUser(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "GRANT")
	rc.rs.InsertAuthUsers(bzc.QueryInt64("roleId"), bzc.QueryInt64Array("userIds"))
	bzc.Success()
	return
}
func (rc *RoleController) CancelAuthUser(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "GRANT")
	userRole := new(systemModels.SysUserRole)
	if err := c.ShouldBindJSON(userRole); err != nil {
		bzc.ParameterError()
		return
	}
	rc.rs.DeleteAuthUserRole(userRole)
	bzc.Success()
	return
}
func (rc *RoleController) CancelAuthUserAll(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("角色管理", "GRANT")
	rc.rs.DeleteAuthUsers(bzc.QueryInt64("roleId"), bzc.QueryInt64Array("userIds"))
	bzc.Success()
	return
}
