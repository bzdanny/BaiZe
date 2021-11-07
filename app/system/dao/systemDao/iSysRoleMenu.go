package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IRoleMenuDao interface {
	BatchRoleMenu(list []*systemModels.SysRoleMenu)
	DeleteRoleMenu(ids []int64)
	DeleteRoleMenuByRoleId(roleId int64)
	CheckMenuExistRole(menuId int64) int
}
