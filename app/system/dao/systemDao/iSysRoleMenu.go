package systemDao

import (
	"baize/app/common/datasource"
	"baize/app/system/models/systemModels"
)

type IRoleMenuDao interface {
	BatchRoleMenu(list []*systemModels.SysRoleMenu, tx ...datasource.Transaction)
	DeleteRoleMenu(ids []int64, tx ...datasource.Transaction)
	DeleteRoleMenuByRoleId(roleId int64, tx ...datasource.Transaction)
	CheckMenuExistRole(menuId int64) int
}
