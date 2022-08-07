package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IRolePermissionDao interface {
	BatchRolePermission(db dataUtil.DB, list []*systemModels.SysRolePermission)
	DeleteRolePermission(db dataUtil.DB, ids []int64)
	DeleteRolePermissionByRoleId(db dataUtil.DB, roleId int64)
	CheckPermissionExistRole(db dataUtil.DB, permissionId int64) int
}
