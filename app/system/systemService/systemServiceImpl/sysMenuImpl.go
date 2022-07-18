package systemServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/constant/userConstants"
	systemDao "github.com/bzdanny/BaiZe/app/system/systemDao"
	systemDaoImpl "github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	systemModels "github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/utils"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/stringUtils"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
)

type MenuService struct {
	data        *datasource.Data
	menuDao     systemDao.IMenuDao
	roleMenuDao systemDao.IRoleMenuDao
	roleDao     systemDao.IRoleDao
}

func NewMenuService(data *datasource.Data, md *systemDaoImpl.SysMenuDao, rmd *systemDaoImpl.SysRoleMenuDao, rd *systemDaoImpl.SysRoleDao) *MenuService {
	return &MenuService{
		data:        data,
		menuDao:     md,
		roleMenuDao: rmd,
		roleDao:     rd,
	}
}

func (menuService *MenuService) SelectMenuList(menu *systemModels.SysMenuDQL, userId int64) (list []*systemModels.SysMenuVo) {
	if utils.IsAdmin(userId) {
		list = menuService.menuDao.SelectMenuList(menuService.data.GetSlaveDb(), menu)
	} else {
		menu.UserId = userId
		list = menuService.menuDao.SelectMenuListByUserId(menuService.data.GetSlaveDb(), menu)
	}
	return
}

func (menuService *MenuService) SelectMenuById(menuId int64) (menu *systemModels.SysMenuVo) {
	return menuService.menuDao.SelectMenuById(menuService.data.GetSlaveDb(), menuId)

}
func (menuService *MenuService) InsertMenu(menu *systemModels.SysMenuDML) {
	menu.MenuId = snowflake.GenID()
	menuService.menuDao.InsertMenu(menuService.data.GetMasterDb(), menu)
}
func (menuService *MenuService) UpdateMenu(menu *systemModels.SysMenuDML) {
	menuService.menuDao.UpdateMenu(menuService.data.GetMasterDb(), menu)
}
func (menuService *MenuService) DeleteMenuById(menuId int64) {
	menuService.menuDao.DeleteMenuById(menuService.data.GetMasterDb(), menuId)
}

func (menuService *MenuService) SelectMenuTreeByUserId(userId int64) (sysMenu []*systemModels.SysMenuVo) {

	if utils.IsAdmin(userId) {
		sysMenu = menuService.menuDao.SelectMenuTreeAll(menuService.data.GetSlaveDb())
	} else {
		sysMenu = menuService.menuDao.SelectMenuTreeByUserId(menuService.data.GetSlaveDb(), userId)
	}
	sysMenu = getChildPerms(sysMenu, 0)
	return
}

func (menuService *MenuService) BuildMenus(sysMenus []*systemModels.SysMenuVo) []*systemModels.RouterVo {
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

func (menuService *MenuService) CheckMenuNameUnique(menu *systemModels.SysMenuDML) bool {
	RoleId := menuService.menuDao.CheckMenuNameUnique(menuService.data.GetSlaveDb(), menu.MenuName, menu.ParentId)
	if RoleId == menu.MenuId || RoleId == 0 {
		return false
	}
	return true
}

func (menuService *MenuService) HasChildByMenuId(menuId int64) bool {
	return menuService.menuDao.HasChildByMenuId(menuService.data.GetSlaveDb(), menuId) > 0
}

func (menuService *MenuService) CheckMenuExistRole(menuId int64) bool {
	return menuService.roleMenuDao.CheckMenuExistRole(menuService.data.GetSlaveDb(), menuId) > 0
}
func (menuService *MenuService) SelectMenuListByRoleId(roleId int64) []string {
	role := menuService.roleDao.SelectRoleById(menuService.data.GetSlaveDb(), roleId)
	return menuService.menuDao.SelectMenuListByRoleId(menuService.data.GetSlaveDb(), roleId, role.MenuCheckStrictly)
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
