package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PermissionController struct {
	ms systemService.IPermissionService
}

func NewPermissionController(ms *systemServiceImpl.PermissionService) *PermissionController {
	return &PermissionController{ms: ms}
}

func (mc *PermissionController) PermissionList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	Permission := new(systemModels.SysPermissionDQL)
	_ = c.ShouldBind(Permission)
	list := mc.ms.SelectPermissionList(Permission, bzc.GetUserId())
	bzc.SuccessData(list)
}
func (mc *PermissionController) PermissionGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	PermissionId := bzc.ParamInt64("PermissionId")
	if PermissionId == 0 {
		bzc.ParameterError()
		return
	}
	Permission := mc.ms.SelectPermissionById(PermissionId)
	bzc.SuccessData(Permission)
}
func (mc *PermissionController) PermissionTreeSelect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.GetUserId()
	bzc.SuccessData(mc.ms.SelectPermissionList(new(systemModels.SysPermissionDQL), userId))
}
func (mc *PermissionController) PermissionAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysPermission := new(systemModels.SysPermissionDML)
	c.ShouldBind(sysPermission)
	if mc.ms.CheckPermissionNameUnique(sysPermission) {
		bzc.Waring("新增菜单'" + sysPermission.PermissionName + "'失败，菜单名称已存在")
		return
	}
	sysPermission.SetCreateBy(bzc.GetUserId())
	mc.ms.InsertPermission(sysPermission)
	bzc.Success()
}
func (mc *PermissionController) PermissionEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysPermission := new(systemModels.SysPermissionDML)
	if mc.ms.CheckPermissionNameUnique(sysPermission) {
		bzc.Waring("修改菜单'" + sysPermission.PermissionName + "'失败，菜单名称已存在")
		return
	}
	c.ShouldBind(sysPermission)
	sysPermission.SetCreateBy(bzc.GetUserId())
	mc.ms.UpdatePermission(sysPermission)
	bzc.Success()
}
func (mc *PermissionController) PermissionRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	PermissionId := bzc.ParamInt64("PermissionId")
	if PermissionId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	if mc.ms.HasChildByPermissionId(PermissionId) {
		bzc.Waring("存在子菜单,不允许删除")
		return
	}
	if mc.ms.CheckPermissionExistRole(PermissionId) {
		bzc.Waring("菜单已分配,不允许删除")
		return
	}
	mc.ms.DeletePermissionById(PermissionId)
	bzc.Success()
}
func (mc *PermissionController) RolePermissionTreeselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	userId := bzc.GetUserId()
	m := make(map[string]interface{})
	m["checkedKeys"] = mc.ms.SelectPermissionListByRoleId(roleId)
	m["permissions"] = mc.ms.SelectPermissionList(new(systemModels.SysPermissionDQL), userId)
	bzc.SuccessData(m)
}
