package systemService

import (
	"baize/app/constant/userConstants"
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
	"baize/app/utils/admin"
	"baize/app/utils/snowflake"
	"baize/app/utils/stringUtils"
)

func SelectMenuList(menu *systemModels.SysMenuDQL, userId int64) (list []*systemModels.SysMenuVo) {
	if admin.IsAdmin(userId) {
		list = systemDao.SelectMenuList(menu)
	} else {
		menu.UserId = userId
		list = systemDao.SelectMenuListByUserId(menu)
	}
	return
}

func SelectMenuById(menuId int64) (menu *systemModels.SysMenuVo) {
	return systemDao.SelectMenuById(menuId)

}
func InsertMenu(menu *systemModels.SysMenuDML) {
	menu.MenuId = snowflake.GenID()
	systemDao.InsertMenu(menu)
}
func UpdateMenu(menu *systemModels.SysMenuDML) {
	systemDao.UpdateMenu(menu)
}
func DeleteMenuById(menuId int64) {
	systemDao.DeleteMenuById(menuId)
}

func SelectMenuTreeByUserId(userId int64) (sysMenu []*systemModels.SysMenuVo) {

	if admin.IsAdmin(userId) {
		sysMenu = systemDao.SelectMenuTreeAll()
	} else {
		sysMenu = systemDao.SelectMenuTreeByUserId(userId)
	}
	sysMenu = getChildPerms(sysMenu, 0)
	return
}

func BuildMenus(sysMenus []*systemModels.SysMenuVo) []*systemModels.RouterVo {
	routerVo := make([]*systemModels.RouterVo, 0, 2)
	for _, m := range sysMenus {
		r := new(systemModels.RouterVo)
		r.Hidden = m.Visible == "1"
		r.Name = m.GetRouteName()
		r.Path = m.GetRouterPath()
		r.Component = m.GetComponent()
		r.Meta.Title = m.MenuName
		r.Meta.Icon = m.Icon
		r.Meta.NoCache = m.IsCache == "1"
		cMenus := m.Children
		if cMenus != nil && len(cMenus) > 0 && m.MenuType == userConstants.TypeDir {
			r.AlwaysShow = true
			r.Redirect = "noRedirect"
			r.Children = BuildMenus(cMenus)
		} else if m.IsMenuFrame() {
			childrenList := make([]*systemModels.RouterVo, 0, 2)
			children := new(systemModels.RouterVo)
			children.Path = m.Path
			children.Component = *m.Component
			children.Name = stringUtils.Capitalize(m.Path)
			children.Meta.Title = m.MenuType
			r.Meta.Icon = m.Icon
			r.Meta.NoCache = m.IsCache == "1"
			childrenList = append(childrenList, children)
			r.Children = childrenList
		}
		routerVo = append(routerVo, r)
	}
	return routerVo
}

func getChildPerms(menu []*systemModels.SysMenuVo, parentId int64) []*systemModels.SysMenuVo {
	sysMenus := make([]*systemModels.SysMenuVo, 0, 2)
	for _, sysMenu := range menu {
		if sysMenu.ParentId == parentId {
			recursionFn(menu, sysMenu)
			sysMenus = append(sysMenus, sysMenu)
		}
	}
	return sysMenus
}

/**
 * 递归列表
 *
 * @param menu
 * @param s
 */
func recursionFn(menu []*systemModels.SysMenuVo, s *systemModels.SysMenuVo) {
	childList := getChildList(menu, s)
	s.Children = childList
	for _, sysMenu := range childList {
		if hasChild(menu, sysMenu) {
			recursionFn(menu, sysMenu)
		}
	}
}

/**
 * 判断是否有子节点
 */
func hasChild(list []*systemModels.SysMenuVo, m *systemModels.SysMenuVo) bool {
	return len(getChildList(list, m)) > 0
}

/**
 * 得到子节点列表
 */
func getChildList(menu []*systemModels.SysMenuVo, s *systemModels.SysMenuVo) []*systemModels.SysMenuVo {
	tlist := make([]*systemModels.SysMenuVo, 0, 2)
	for _, sysMenu := range menu {
		if sysMenu.ParentId == s.MenuId {
			tlist = append(tlist, sysMenu)
		}

	}
	return tlist
}

func CheckMenuNameUnique(menu *systemModels.SysMenuDML) bool {
	RoleId := systemDao.CheckMenuNameUnique(menu.MenuName, menu.ParentId)
	if RoleId == menu.MenuId || RoleId == 0 {
		return false
	}
	return true
}

func HasChildByMenuId(menuId int64) bool {
	return systemDao.HasChildByMenuId(menuId) > 0
}

func CheckMenuExistRole(menuId int64) bool {
	return systemDao.CheckMenuExistRole(menuId) > 0
}
