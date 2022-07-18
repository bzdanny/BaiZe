package systemDao

import (
	systemModels2 "github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IRoleDao interface {
	SelectRoleList(db dataUtil.DB, role *systemModels2.SysRoleDQL) (roleList []*systemModels2.SysRoleVo, total *int64)
	SelectRoleById(db dataUtil.DB, roleId int64) (role *systemModels2.SysRoleVo)
	SelectBasicRolesByUserId(db dataUtil.DB, userId int64) (roles []*systemModels2.SysRole)
	SelectRolePermissionByUserId(db dataUtil.DB, userId int64) (roles []string)
	SelectRoleIdAndDataScopeByUserId(db dataUtil.DB, userId int64) (roles []*systemModels2.Role)
	SelectRoleListByUserId(db dataUtil.DB, userId int64) (list []int64)
	InsertRole(db dataUtil.DB, sysRole *systemModels2.SysRoleAdd)
	UpdateRole(db dataUtil.DB, sysRole *systemModels2.SysRoleEdit)
	DeleteRoleByIds(db dataUtil.DB, ids []int64)
	CheckRoleNameUnique(db dataUtil.DB, roleName string) int64
	CheckRoleKeyUnique(db dataUtil.DB, roleKey string) int64
	SelectAllocatedList(db dataUtil.DB, user *systemModels2.SysRoleAndUserDQL) (list []*systemModels2.SysUserVo, total *int64)
	SelectUnallocatedList(db dataUtil.DB, user *systemModels2.SysRoleAndUserDQL) (list []*systemModels2.SysUserVo, total *int64)
}
