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

//DeptList 部门列表查询
func (dc *DeptController) DeptList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	dept := new(systemModels.SysDeptDQL)
	_ = c.ShouldBind(dept)
	dept.SetDataScope(bzc.GetUser(), "d", "")
	bzc.SuccessData(dc.ds.SelectDeptList(dept))

}

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
func (dc *DeptController) RoleDeptTreeselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	m := make(map[string]interface{})
	m["checkedKeys"] = dc.ds.SelectDeptListByRoleId(roleId)
	m["depts"] = dc.ds.SelectDeptList(new(systemModels.SysDeptDQL))
	bzc.SuccessData(m)
}

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
