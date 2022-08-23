package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IPermissionDao interface {
	SelectPermissionById(db dataUtil.DB, permissionId int64) (Permission *systemModels.SysPermissionVo)
	SelectPermissionList(db dataUtil.DB, permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo)
	SelectPermissionListByUserId(db dataUtil.DB, Permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo)
	InsertPermission(db dataUtil.DB, permission *systemModels.SysPermissionDML)
	UpdatePermission(db dataUtil.DB, permission *systemModels.SysPermissionDML)
	DeletePermissionById(db dataUtil.DB, permissionId int64)
	SelectPermissionPermsByUserId(db dataUtil.DB, userId int64) (perms []string)
	CheckPermissionNameUnique(db dataUtil.DB, permissionName string, parentId int64) int64
	HasChildByPermissionId(db dataUtil.DB, permissionId int64) int
	SelectPermissionListByRoleId(db dataUtil.DB, roleId int64, PermissionCheckStrictly bool) (roleIds []string)
}
