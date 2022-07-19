package systemService

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
)

type IRoleService interface {
	SelectRoleList(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, count *int64)
	RoleExport(role *systemModels.SysRoleDQL) (data []byte)
	SelectRoleById(roseId int64) (role *systemModels.SysRoleVo)
	InsertRole(sysRole *systemModels.SysRoleAdd)
	UpdateRole(sysRole *systemModels.SysRoleEdit)
	UpdateRoleStatus(sysRole *systemModels.SysRoleEdit)
	AuthDataScope(sysRole *systemModels.SysRoleEdit)
	DeleteRoleByIds(ids []int64)
	CountUserRoleByRoleId(ids []int64) bool
	SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole)
	SelectRoleAll(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo)
	RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []*systemModels.Role)
	SelectRoleListByUserId(userId int64) (list []int64)
	CheckRoleNameUnique(id int64, roleName string) bool
	CheckRoleKeyUnique(id int64, roleKey string) bool
	SelectUserRoleGroupByUserId(userId int64) string
	SelectAllocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
	SelectUnallocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
	InsertAuthUsers(roleId int64, userIds []int64)
	DeleteAuthUsers(roleId int64, userIds []int64)
	DeleteAuthUserRole(user *systemModels.SysUserRole)
}
