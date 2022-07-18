package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IRoleMenuDao interface {
	BatchRoleMenu(db dataUtil.DB, list []*systemModels.SysRoleMenu)
	DeleteRoleMenu(db dataUtil.DB, ids []int64)
	DeleteRoleMenuByRoleId(db dataUtil.DB, roleId int64)
	CheckMenuExistRole(db dataUtil.DB, menuId int64) int
}
