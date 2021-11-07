package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func RoleList(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(page)
	role.SetLimit(page)
	role.SetDataScope(loginUser, "d", "")
	list, count := iRole.SelectRoleList(role)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func RoleExport(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetDataScope(loginUser, "d", "")
	data := iRole.RoleExport(role)
	c.Header("Content-Type", "application/vnd.ms-excel")
	c.Data(http.StatusOK, "application/vnd.ms-excel", data)

}
func RoleGetInfo(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	sysUser := iRole.SelectRoleById(roleId)

	c.JSON(http.StatusOK, commonModels.SuccessData(sysUser))
}
func RoleAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysRole := new(systemModels.SysRoleDML)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	if iRole.CheckRoleNameUnique(sysRole) {
		c.JSON(http.StatusOK, commonModels.Waring("新增角色'"+sysRole.RoleName+"'失败，角色名称已存在"))
		return
	}
	if iRole.CheckRoleKeyUnique(sysRole) {
		c.JSON(http.StatusOK, commonModels.Waring("新增角色'"+sysRole.RoleKey+"'失败，角色权限已存在"))
		return
	}

	sysRole.SetCreateBy(loginUser.User.UserName)
	iRole.InsertRole(sysRole)

	c.JSON(http.StatusOK, commonModels.Success())

}
func RoleEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysRole := new(systemModels.SysRoleDML)
	if err := c.ShouldBindJSON(sysRole); err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	if iRole.CheckRoleNameUnique(sysRole) {
		c.JSON(http.StatusOK, commonModels.Waring("新增角色'"+sysRole.RoleName+"'失败，角色名称已存在"))
		return
	}
	if iRole.CheckRoleKeyUnique(sysRole) {
		c.JSON(http.StatusOK, commonModels.Waring("新增角色'"+sysRole.RoleKey+"'失败，角色权限已存在"))
		return
	}

	sysRole.SetUpdateBy(loginUser.User.UserName)
	iRole.UpdateRole(sysRole)

	c.JSON(http.StatusOK, commonModels.Success())

}
func RoleDataScope(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysRole := new(systemModels.SysRoleDML)
	c.ShouldBindJSON(sysRole)
	sysRole.SetUpdateBy(loginUser.User.UserName)
	iRole.AuthDataScope(sysRole)
	c.JSON(http.StatusOK, commonModels.Success())

}
func RoleChangeStatus(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysRole := new(systemModels.SysRoleDML)
	c.ShouldBindJSON(sysRole)

	sysRole.SetUpdateBy(loginUser.User.UserName)
	iRole.UpdateRoleStatus(sysRole)

	c.JSON(http.StatusOK, commonModels.Success())
}
func RoleRemove(c *gin.Context) {

	var s slicesUtils.Slices = strings.Split(c.Param("rolesIds"), ",")
	ids := s.StrSlicesToInt()
	if iRole.CountUserRoleByRoleId(ids) {
		c.JSON(http.StatusOK, commonModels.Waring("角色已分配，不能删除"))
		return
	}
	iRole.DeleteRoleByIds(ids)

	c.JSON(http.StatusOK, commonModels.Success())
}
