package systemDao

import (
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
	InsertRole(sysRole *systemModels.SysRoleDML)
	UpdateRole(sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(ids []int64)
	CheckRoleNameUnique(roleName string) int64
	CheckRoleKeyUnique(roleKey string) int64
}
