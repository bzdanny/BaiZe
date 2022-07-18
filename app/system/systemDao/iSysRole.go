package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IRoleDao interface {
	SelectRoleList(db dataUtil.DB, role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total *int64)
	SelectRoleById(db dataUtil.DB, roleId int64) (role *systemModels.SysRoleVo)
	SelectBasicRolesByUserId(db dataUtil.DB, userId int64) (roles []*systemModels.SysRole)
	SelectRolePermissionByUserId(db dataUtil.DB, userId int64) (roles []string)
	SelectRoleIdAndDataScopeByUserId(db dataUtil.DB, userId int64) (roles []*systemModels.Role)
	SelectRoleListByUserId(db dataUtil.DB, userId int64) (list []int64)
	InsertRole(db dataUtil.DB, sysRole *systemModels.SysRoleAdd)
	UpdateRole(db dataUtil.DB, sysRole *systemModels.SysRoleEdit)
	DeleteRoleByIds(db dataUtil.DB, ids []int64)
	CheckRoleNameUnique(db dataUtil.DB, roleName string) int64
	CheckRoleKeyUnique(db dataUtil.DB, roleKey string) int64
	SelectAllocatedList(db dataUtil.DB, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
	SelectUnallocatedList(db dataUtil.DB, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64)
}
