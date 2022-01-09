package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
)

type IRoleDao interface {
	SelectRoleList(role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total *int64)
	SelectRoleById(roleId int64) (role *systemModels.SysRoleVo)
	SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole)
	SelectRolePermissionByUserId(userId int64) (roles []string)
	SelectRoleIdAndDataScopeByUserId(userId int64) (roles []*loginModels.Role)
	SelectRoleListByUserId(userId int64) (list []int64)
	InsertRole(sysRole *systemModels.SysRoleDML, tx ...mysql.Transaction)
	UpdateRole(sysRole *systemModels.SysRoleDML, tx ...mysql.Transaction)
	DeleteRoleByIds(ids []int64, tx ...mysql.Transaction)
	CheckRoleNameUnique(roleName string) int64
	CheckRoleKeyUnique(roleKey string) int64
	SelectAllocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
	SelectUnallocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
}
