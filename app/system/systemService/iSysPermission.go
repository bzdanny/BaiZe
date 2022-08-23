package systemService

import "github.com/bzdanny/BaiZe/app/system/systemModels"

type IPermissionService interface {
	SelectPermissionList(Permission *systemModels.SysPermissionDQL, userId int64) (list []*systemModels.SysPermissionVo)
	SelectPermissionById(PermissionId int64) (Permission *systemModels.SysPermissionVo)
	InsertPermission(Permission *systemModels.SysPermissionDML)
	UpdatePermission(Permission *systemModels.SysPermissionDML)
	DeletePermissionById(PermissionId int64)
	CheckPermissionNameUnique(Permission *systemModels.SysPermissionDML) bool
	HasChildByPermissionId(PermissionId int64) bool
	CheckPermissionExistRole(PermissionId int64) bool
	SelectPermissionListByRoleId(roleId int64) []string
}
