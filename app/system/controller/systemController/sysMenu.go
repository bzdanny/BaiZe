package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func MenuList(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	menu := new(systemModels.SysMenuDQL)
	c.ShouldBind(menu)
	list := iMenu.SelectMenuList(menu, loginUser.User.UserId)

	c.JSON(http.StatusOK, commonModels.SuccessData(list))

}
func MenuGetInfo(c *gin.Context) {
	menuId, err := strconv.ParseInt(c.Param("menuId"), 10, 64)
	if err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	menu := iMenu.SelectMenuById(menuId)

	c.JSON(http.StatusOK, commonModels.SuccessData(menu))
}
func MenuTreeSelect(c *gin.Context) {

}
func MenuAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysMenu := new(systemModels.SysMenuDML)
	c.ShouldBind(sysMenu)
	if iMenu.CheckMenuNameUnique(sysMenu) {
		c.JSON(http.StatusOK, commonModels.Waring("新增菜单'"+sysMenu.MenuName+"'失败，菜单名称已存在"))
		return
	}
	sysMenu.SetCreateBy(loginUser.User.UserName)
	iMenu.InsertMenu(sysMenu)

	c.JSON(http.StatusOK, commonModels.Success())
}
func MenuEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysMenu := new(systemModels.SysMenuDML)
	if iMenu.CheckMenuNameUnique(sysMenu) {
		c.JSON(http.StatusOK, commonModels.Waring("修改菜单'"+sysMenu.MenuName+"'失败，菜单名称已存在"))
		return
	}
	c.ShouldBind(sysMenu)
	sysMenu.SetCreateBy(loginUser.User.UserName)
	iMenu.UpdateMenu(sysMenu)

	c.JSON(http.StatusOK, commonModels.Success())
}
func MenuRemove(c *gin.Context) {
	menuId, err := strconv.ParseInt(c.Param("menuId"), 10, 64)
	if err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
	}
	if iMenu.HasChildByMenuId(menuId) {
		c.JSON(http.StatusOK, commonModels.Waring("存在子菜单,不允许删除"))
		return
	}
	if iMenu.CheckMenuExistRole(menuId) {
		c.JSON(http.StatusOK, commonModels.Waring("菜单已分配,不允许删除"))
		return
	}

	iMenu.DeleteMenuById(menuId)

	c.JSON(http.StatusOK, commonModels.Success())
}
func RoleMenuTreeselect(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
	}
	userId := commonController.GetCurrentLoginUser(c).User.UserId

	m := make(map[string]interface{})
	m["checkedKeys"] = iMenu.SelectMenuListByRoleId(roleId)
	m["menus"] = iMenu.SelectMenuList(new(systemModels.SysMenuDQL), userId)
	c.JSON(http.StatusOK, commonModels.SuccessData(m))
}
