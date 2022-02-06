package DeptController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var iDept systemService.IDeptService = systemServiceImpl.GetDeptService()

//DeptList 部门列表查询
func DeptList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	loginUser := bzc.GetCurrentLoginUser()
	dept := new(systemModels.SysDeptDQL)
	c.ShouldBind(dept)
	dept.SetDataScope(loginUser, "d", "")
	list := iDept.SelectDeptList(dept)
	bzc.SuccessData(list)

}

func DeptGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	deptId := bzc.ParamInt64("deptId")
	if deptId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	menu := iDept.SelectDeptById(deptId)
	bzc.SuccessData(menu)
}
func RoleDeptTreeselect(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	roleId := bzc.ParamInt64("roleId")
	if roleId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
	}
	m := make(map[string]interface{})
	m["checkedKeys"] = iDept.SelectDeptListByRoleId(roleId)
	m["depts"] = iDept.SelectDeptList(new(systemModels.SysDeptDQL))
	bzc.SuccessData(m)
}

func DeptAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("部门管理", "INSERT")
	loginUser := bzc.GetCurrentLoginUser()
	sysDept := new(systemModels.SysDeptDML)
	c.ShouldBind(sysDept)
	if iDept.CheckDeptNameUnique(sysDept) {
		bzc.Waring("新增部门'" + sysDept.DeptName + "'失败，部门名称已存在")
		return
	}
	sysDept.SetCreateBy(loginUser.User.UserName)
	iDept.InsertDept(sysDept)
	bzc.Success()
}
func DeptEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("部门管理", "UPDATE")
	loginUser := bzc.GetCurrentLoginUser()
	sysDept := new(systemModels.SysDeptDML)
	if iDept.CheckDeptNameUnique(sysDept) {
		bzc.Waring("修改部门'" + sysDept.DeptName + "'失败，部门名称已存在")
		return
	}
	c.ShouldBind(sysDept)
	sysDept.SetCreateBy(loginUser.User.UserName)
	iDept.UpdateDept(sysDept)
	bzc.Success()
}
func DeptRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("部门管理", "DELETE")
	deptId := bzc.ParamInt64("deptId")
	if deptId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	if iDept.HasChildByDeptId(deptId) {
		bzc.Waring("存在下级部门,不允许删除")
		return
	}
	if iDept.CheckDeptExistUser(deptId) {
		bzc.Waring("部门存在用户,不允许删除")
		return
	}
	iDept.DeleteDeptById(deptId)

	bzc.Success()
}
