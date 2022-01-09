package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
)

type IRoleMenuDao interface {
	BatchRoleMenu(list []*systemModels.SysRoleMenu, tx ...mysql.Transaction)
	DeleteRoleMenu(ids []int64, tx ...mysql.Transaction)
	DeleteRoleMenuByRoleId(roleId int64, tx ...mysql.Transaction)
	CheckMenuExistRole(menuId int64) int
}
