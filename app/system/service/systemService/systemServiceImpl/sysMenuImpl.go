package systemServiceImpl

import (
	"baize/app/constant/userConstants"
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/systemModels"
	"baize/app/utils/admin"
	"baize/app/utils/snowflake"
	"baize/app/utils/stringUtils"
)

var menuServiceImpl *menuService = &menuService{menuDao: systemDaoImpl.GetSysMenuDao(), roleMenuDao: systemDaoImpl.GetSysRoleMenuDao(), roleDao: systemDaoImpl.GetSysRoleDao()}

type menuService struct {
	menuDao     systemDao.IMenuDao
	roleMenuDao systemDao.IRoleMenuDao
	roleDao     systemDao.IRoleDao
}

func GetMenuService() *menuService {
	return menuServiceImpl
}

func (menuService *menuService) SelectMenuList(menu *systemModels.SysMenuDQL, userId int64) (list []*systemModels.SysMenuVo) {
	if admin.IsAdmin(userId) {
		list = menuService.menuDao.SelectMenuList(menu)
	} else {
		menu.UserId = userId
		list = menuService.menuDao.SelectMenuListByUserId(menu)
	}
	return
}

func (menuService *menuService) SelectMenuById(menuId int64) (menu *systemModels.SysMenuVo) {
	return menuService.menuDao.SelectMenuById(menuId)

}
func (menuService *menuService) InsertMenu(menu *systemModels.SysMenuDML) {
	menu.MenuId = snowflake.GenID()
	menuService.menuDao.InsertMenu(menu)
}
func (menuService *menuService) UpdateMenu(menu *systemModels.SysMenuDML) {
	menuService.menuDao.UpdateMenu(menu)
}
func (menuService *menuService) DeleteMenuById(menuId int64) {
	menuService.menuDao.DeleteMenuById(menuId)
}

func (menuService *menuService) SelectMenuTreeByUserId(userId int64) (sysMenu []*systemModels.SysMenuVo) {

	if admin.IsAdmin(userId) {
		sysMenu = menuService.menuDao.SelectMenuTreeAll()
	} else {
		sysMenu = menuService.menuDao.SelectMenuTreeByUserId(userId)
	}
	sysMenu = getChildPerms(sysMenu, 0)
	return
}

func (menuService *menuService) BuildMenus(sysMenus []*systemModels.SysMenuVo) []*systemModels.RouterVo {
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
			r.Children = menuService.BuildMenus(cMenus)
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

func (menuService *menuService) CheckMenuNameUnique(menu *systemModels.SysMenuDML) bool {
	RoleId := menuService.menuDao.CheckMenuNameUnique(menu.MenuName, menu.ParentId)
	if RoleId == menu.MenuId || RoleId == 0 {
		return false
	}
	return true
}

func (menuService *menuService) HasChildByMenuId(menuId int64) bool {
	return menuService.menuDao.HasChildByMenuId(menuId) > 0
}

func (menuService *menuService) CheckMenuExistRole(menuId int64) bool {
	return menuService.roleMenuDao.CheckMenuExistRole(menuId) > 0
}
func (menuService *menuService) SelectMenuListByRoleId(roleId int64) []string {
	role := menuService.roleDao.SelectRoleById(roleId)
	return menuService.menuDao.SelectMenuListByRoleId(roleId, role.MenuCheckStrictly)
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
