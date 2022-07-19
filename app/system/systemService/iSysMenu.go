package systemService

import "github.com/bzdanny/BaiZe/app/system/systemModels"

type IMenuService interface {
	SelectMenuList(menu *systemModels.SysMenuDQL, userId int64) (list []*systemModels.SysMenuVo)
	SelectMenuById(menuId int64) (menu *systemModels.SysMenuVo)
	InsertMenu(menu *systemModels.SysMenuDML)
	UpdateMenu(menu *systemModels.SysMenuDML)
	DeleteMenuById(menuId int64)
	SelectMenuTreeByUserId(userId int64) (sysMenu []*systemModels.SysMenuVo)
	BuildMenus(sysMenus []*systemModels.SysMenuVo) []*systemModels.RouterVo
	CheckMenuNameUnique(menu *systemModels.SysMenuDML) bool
	HasChildByMenuId(menuId int64) bool
	CheckMenuExistRole(menuId int64) bool
	SelectMenuListByRoleId(roleId int64) []string
}
