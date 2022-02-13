package menuController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iMenu systemService.IMenuService = systemServiceImpl.GetMenuService()

func MenuList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	menu := new(systemModels.SysMenuDQL)
	c.ShouldBind(menu)
	list := iMenu.SelectMenuList(menu, bzc.GetCurrentUserId())
	bzc.SuccessData(list)
}
func MenuGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	menuId := bzc.ParamInt64("menuId")
	if menuId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	menu := iMenu.SelectMenuById(menuId)
	bzc.SuccessData(menu)
}
func MenuTreeSelect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.GetCurrentUserId()
	bzc.SuccessData(iMenu.SelectMenuList(new(systemModels.SysMenuDQL), userId))
}
func MenuAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("菜单管理", "INSERT")
	sysMenu := new(systemModels.SysMenuDML)
	c.ShouldBind(sysMenu)
	if iMenu.CheckMenuNameUnique(sysMenu) {
		bzc.Waring("新增菜单'" + sysMenu.MenuName + "'失败，菜单名称已存在")
		return
	}
	sysMenu.SetCreateBy(bzc.GetCurrentUserName())
	iMenu.InsertMenu(sysMenu)
	bzc.Success()
}
func MenuEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("菜单管理", "UPDATE")
	sysMenu := new(systemModels.SysMenuDML)
	if iMenu.CheckMenuNameUnique(sysMenu) {
		bzc.Waring("修改菜单'" + sysMenu.MenuName + "'失败，菜单名称已存在")
		return
	}
	c.ShouldBind(sysMenu)
	sysMenu.SetCreateBy(bzc.GetCurrentUserName())
	iMenu.UpdateMenu(sysMenu)
	bzc.Success()
}
func MenuRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("菜单管理", "DELETE")
	menuId := bzc.ParamInt64("menuId")
	if menuId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	if iMenu.HasChildByMenuId(menuId) {
		bzc.Waring("存在子菜单,不允许删除")
		return
	}
	if iMenu.CheckMenuExistRole(menuId) {
		bzc.Waring("菜单已分配,不允许删除")
		return
	}
	iMenu.DeleteMenuById(menuId)
	bzc.Success()
}
func RoleMenuTreeselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	userId := bzc.GetCurrentUserId()
	m := make(map[string]interface{})
	m["checkedKeys"] = iMenu.SelectMenuListByRoleId(roleId)
	m["menus"] = iMenu.SelectMenuList(new(systemModels.SysMenuDQL), userId)
	bzc.SuccessData(m)
}
