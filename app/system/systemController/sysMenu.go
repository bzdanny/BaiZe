package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuController struct {
	ms systemService.IMenuService
}

func NewMenuController(ms *systemServiceImpl.MenuService) *MenuController {
	return &MenuController{ms: ms}
}

func (mc *MenuController) MenuList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	menu := new(systemModels.SysMenuDQL)
	_ = c.ShouldBind(menu)
	list := mc.ms.SelectMenuList(menu, bzc.GetUserId())
	bzc.SuccessData(list)
}
func (mc *MenuController) MenuGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	menuId := bzc.ParamInt64("menuId")
	if menuId == 0 {
		bzc.ParameterError()
		return
	}
	menu := mc.ms.SelectMenuById(menuId)
	bzc.SuccessData(menu)
}
func (mc *MenuController) MenuTreeSelect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.GetUserId()
	bzc.SuccessData(mc.ms.SelectMenuList(new(systemModels.SysMenuDQL), userId))
}
func (mc *MenuController) MenuAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("菜单管理", "INSERT")
	sysMenu := new(systemModels.SysMenuDML)
	c.ShouldBind(sysMenu)
	if mc.ms.CheckMenuNameUnique(sysMenu) {
		bzc.Waring("新增菜单'" + sysMenu.MenuName + "'失败，菜单名称已存在")
		return
	}
	sysMenu.SetCreateBy(bzc.GetUserId())
	mc.ms.InsertMenu(sysMenu)
	bzc.Success()
}
func (mc *MenuController) MenuEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("菜单管理", "UPDATE")
	sysMenu := new(systemModels.SysMenuDML)
	if mc.ms.CheckMenuNameUnique(sysMenu) {
		bzc.Waring("修改菜单'" + sysMenu.MenuName + "'失败，菜单名称已存在")
		return
	}
	c.ShouldBind(sysMenu)
	sysMenu.SetCreateBy(bzc.GetUserId())
	mc.ms.UpdateMenu(sysMenu)
	bzc.Success()
}
func (mc *MenuController) MenuRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("菜单管理", "DELETE")
	menuId := bzc.ParamInt64("menuId")
	if menuId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	if mc.ms.HasChildByMenuId(menuId) {
		bzc.Waring("存在子菜单,不允许删除")
		return
	}
	if mc.ms.CheckMenuExistRole(menuId) {
		bzc.Waring("菜单已分配,不允许删除")
		return
	}
	mc.ms.DeleteMenuById(menuId)
	bzc.Success()
}
func (mc *MenuController) RoleMenuTreeselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	userId := bzc.GetUserId()
	m := make(map[string]interface{})
	m["checkedKeys"] = mc.ms.SelectMenuListByRoleId(roleId)
	m["menus"] = mc.ms.SelectMenuList(new(systemModels.SysMenuDQL), userId)
	bzc.SuccessData(m)
}
