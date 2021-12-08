package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IMenuDao interface {
	SelectMenuById(menuId int64) (menu *systemModels.SysMenuVo)
	SelectMenuList(menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	SelectMenuListByUserId(menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	InsertMenu(menu *systemModels.SysMenuDML)
	UpdateMenu(menu *systemModels.SysMenuDML)
	DeleteMenuById(menuId int64)
	SelectMenuPermsByUserId(userId int64) (perms []string)
	SelectMenuTreeAll() (sysMenus []*systemModels.SysMenuVo)
	SelectMenuTreeByUserId(userId int64) (sysMenus []*systemModels.SysMenuVo)
	CheckMenuNameUnique(menuName string, parentId int64) int64
	HasChildByMenuId(menuId int64) int
	SelectMenuListByRoleId(roleId int64, menuCheckStrictly bool) (roleIds []string)
}
