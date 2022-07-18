package systemServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/constant/userConstants"
	systemDao2 "github.com/bzdanny/BaiZe/app/system/systemDao"
	systemDaoImpl2 "github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	systemModels2 "github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/utils"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/stringUtils"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
)

type MenuService struct {
	data        *datasource.Data
	menuDao     systemDao2.IMenuDao
	roleMenuDao systemDao2.IRoleMenuDao
	roleDao     systemDao2.IRoleDao
}

func NewMenuService(data *datasource.Data, md *systemDaoImpl2.SysMenuDao, rmd *systemDaoImpl2.SysRoleMenuDao, rd *systemDaoImpl2.SysRoleDao) *MenuService {
	return &MenuService{
		data:        data,
		menuDao:     md,
		roleMenuDao: rmd,
		roleDao:     rd,
	}
}

func (menuService *MenuService) SelectMenuList(menu *systemModels2.SysMenuDQL, userId int64) (list []*systemModels2.SysMenuVo) {
	if utils.IsAdmin(userId) {
		list = menuService.menuDao.SelectMenuList(menuService.data.GetSlaveDb(), menu)
	} else {
		menu.UserId = userId
		list = menuService.menuDao.SelectMenuListByUserId(menuService.data.GetSlaveDb(), menu)
	}
	return
}

func (menuService *MenuService) SelectMenuById(menuId int64) (menu *systemModels2.SysMenuVo) {
	return menuService.menuDao.SelectMenuById(menuService.data.GetSlaveDb(), menuId)

}
func (menuService *MenuService) InsertMenu(menu *systemModels2.SysMenuDML) {
	menu.MenuId = snowflake.GenID()
	menuService.menuDao.InsertMenu(menuService.data.GetMasterDb(), menu)
}
func (menuService *MenuService) UpdateMenu(menu *systemModels2.SysMenuDML) {
	menuService.menuDao.UpdateMenu(menuService.data.GetMasterDb(), menu)
}
func (menuService *MenuService) DeleteMenuById(menuId int64) {
	menuService.menuDao.DeleteMenuById(menuService.data.GetMasterDb(), menuId)
}

func (menuService *MenuService) SelectMenuTreeByUserId(userId int64) (sysMenu []*systemModels2.SysMenuVo) {

	if utils.IsAdmin(userId) {
		sysMenu = menuService.menuDao.SelectMenuTreeAll(menuService.data.GetSlaveDb())
	} else {
		sysMenu = menuService.menuDao.SelectMenuTreeByUserId(menuService.data.GetSlaveDb(), userId)
	}
	sysMenu = getChildPerms(sysMenu, 0)
	return
}

func (menuService *MenuService) BuildMenus(sysMenus []*systemModels2.SysMenuVo) []*systemModels2.RouterVo {
	routerVo := make([]*systemModels2.RouterVo, 0, 2)
	for _, m := range sysMenus {
		r := new(systemModels2.RouterVo)
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
			childrenList := make([]*systemModels2.RouterVo, 0, 2)
			children := new(systemModels2.RouterVo)
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

func (menuService *MenuService) CheckMenuNameUnique(menu *systemModels2.SysMenuDML) bool {
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

func getChildPerms(menu []*systemModels2.SysMenuVo, parentId int64) []*systemModels2.SysMenuVo {
	sysMenus := make([]*systemModels2.SysMenuVo, 0, 2)
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
func recursionFn(menu []*systemModels2.SysMenuVo, s *systemModels2.SysMenuVo) {
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
func hasChild(list []*systemModels2.SysMenuVo, m *systemModels2.SysMenuVo) bool {
	return len(getChildList(list, m)) > 0
}

/**
 * 得到子节点列表
 */
func getChildList(menu []*systemModels2.SysMenuVo, s *systemModels2.SysMenuVo) []*systemModels2.SysMenuVo {
	tlist := make([]*systemModels2.SysMenuVo, 0, 2)
	for _, sysMenu := range menu {
		if sysMenu.ParentId == s.MenuId {
			tlist = append(tlist, sysMenu)
		}

	}
	return tlist
}
