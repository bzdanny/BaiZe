package systemController

import (
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeptController struct {
	ds systemService.IDeptService
}

func NewDeptController(ds *systemServiceImpl.DeptService) *DeptController {
	return &DeptController{ds: ds}
}

// DeptList 查询部门列表查询
// @Summary 查询部门列表查询
// @Description 查询部门列表查询
// @Tags 部门相关
// @Param  object query systemModels.SysDeptDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=commonModels.ListData{Rows=[]systemModels.SysDeptVo}}  "成功"
// @Router /system/dept  [get]
func (dc *DeptController) DeptList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dept := new(systemModels.SysDeptDQL)
	_ = c.ShouldBind(dept)
	dept.SetDataScope(bzc.GetUser(), "d", "")
	bzc.SuccessData(dc.ds.SelectDeptList(dept))

}

// DeptGetInfo 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags 部门相关
// @Param id path string true "deptId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=systemModels.SysDeptVo}  "成功"
// @Router /system/dept/{deptId}  [get]
func (dc *DeptController) DeptGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	deptId := bzc.ParamInt64("deptId")
	if deptId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	bzc.SuccessData(dc.ds.SelectDeptById(deptId))
}

// RoleDeptTreeSelect 获取角色部门
// @Summary 获取角色部门
// @Description 获取角色部门
// @Tags 部门相关
// @Param id path string true "roleId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=systemModels.RoleDeptTree}  "成功"
// @Router /system/dept/roleDeptTreeSelect/{roleId}  [get]
func (dc *DeptController) RoleDeptTreeSelect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	rdt := new(systemModels.RoleDeptTree)
	rdt.CheckedKeys = dc.ds.SelectDeptListByRoleId(roleId)
	rdt.Depts = dc.ds.SelectDeptList(new(systemModels.SysDeptDQL))
	bzc.SuccessData(rdt)
}

// DeptAdd 添加部门
// @Summary 添加部门
// @Description 添加部门
// @Tags 部门相关
// @Param  object body systemModels.SysDeptAdd true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData "成功"
// @Router /system/dept  [post]
func (dc *DeptController) DeptAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysDept := new(systemModels.SysDeptAdd)
	if err := c.ShouldBindJSON(sysDept); err != nil {
		fmt.Println(err)
		bzc.ParameterError()
		return
	}
	if dc.ds.CheckDeptNameUnique(0, sysDept.ParentId, sysDept.DeptName) {
		bzc.Waring("新增部门'" + sysDept.DeptName + "'失败，部门名称已存在")
		return
	}
	sysDept.SetCreateBy(bzc.GetUserId())
	dc.ds.InsertDept(sysDept)
	bzc.Success()
}

// DeptEdit 修改部门
// @Summary 修改部门
// @Description 修改部门
// @Tags 部门相关
// @Param  object body systemModels.SysDeptEdit true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData "成功"
// @Router /system/dept  [put]
func (dc *DeptController) DeptEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysDept := new(systemModels.SysDeptEdit)
	if err := c.ShouldBindJSON(sysDept); err != nil {
		bzc.ParameterError()
		return
	}
	if dc.ds.CheckDeptNameUnique(sysDept.DeptId, sysDept.ParentId, sysDept.DeptName) {
		bzc.Waring("修改部门'" + sysDept.DeptName + "'失败，部门名称已存在")
		return
	}
	sysDept.SetUpdateBy(bzc.GetUserId())
	dc.ds.UpdateDept(sysDept)
	bzc.Success()
}

// DeptRemove 删除部门
// @Summary 删除部门
// @Description 删除部门
// @Tags 部门相关
// @Param id path string true "deptId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData "成功"
// @Router /system/dept/{deptId}  [delete]
func (dc *DeptController) DeptRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	deptId := bzc.ParamInt64("deptId")
	if deptId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	if dc.ds.HasChildByDeptId(deptId) {
		bzc.Waring("存在下级部门,不允许删除")
		return
	}
	if dc.ds.CheckDeptExistUser(deptId) {
		bzc.Waring("部门存在用户,不允许删除")
		return
	}
	dc.ds.DeleteDeptById(deptId)
	bzc.Success()
}
