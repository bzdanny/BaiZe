package systemService

import (
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
)

type IRoleService interface {
	SelectRoleList(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, count *int64)
	RoleExport(role *systemModels.SysRoleDQL) (data []byte)
	SelectRoleById(roseId int64) (role *systemModels.SysRoleVo)
	InsertRole(sysRole *systemModels.SysRoleDML)
	UpdateRole(sysRole *systemModels.SysRoleDML)
	UpdateRoleStatus(sysRole *systemModels.SysRoleDML)
	AuthDataScope(sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(ids []int64) (err error)
	CountUserRoleByRoleId(ids []int64) bool
	SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole)
	SelectRoleAll(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo)
	RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []*loginModels.Role)
	SelectRoleListByUserId(userId int64) (list []int64)
	InsertRoleMenu(sysRole *systemModels.SysRoleDML)
	CheckRoleNameUnique(sysRole *systemModels.SysRoleDML) bool
	CheckRoleKeyUnique(sysRole *systemModels.SysRoleDML) bool
	SelectUserRoleGroupByUserId(userId int64) string
	SelectAllocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
	SelectUnallocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
	InsertAuthUsers(roleId int64,userIds []int64)
	DeleteAuthUsers(roleId int64,userIds []int64)
	DeleteAuthUserRole(user *systemModels.SysUserRole)
}
