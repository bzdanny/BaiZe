package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IMenuDao interface {
	SelectMenuById(db dataUtil.DB, menuId int64) (menu *systemModels.SysMenuVo)
	SelectMenuList(db dataUtil.DB, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	SelectMenuListByUserId(db dataUtil.DB, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	InsertMenu(db dataUtil.DB, menu *systemModels.SysMenuDML)
	UpdateMenu(db dataUtil.DB, menu *systemModels.SysMenuDML)
	DeleteMenuById(db dataUtil.DB, menuId int64)
	SelectMenuPermsByUserId(db dataUtil.DB, userId int64) (perms []string)
	SelectMenuTreeAll(db dataUtil.DB) (sysMenus []*systemModels.SysMenuVo)
	SelectMenuTreeByUserId(db dataUtil.DB, userId int64) (sysMenus []*systemModels.SysMenuVo)
	CheckMenuNameUnique(db dataUtil.DB, menuName string, parentId int64) int64
	HasChildByMenuId(db dataUtil.DB, menuId int64) int
	SelectMenuListByRoleId(db dataUtil.DB, roleId int64, menuCheckStrictly bool) (roleIds []string)
}
