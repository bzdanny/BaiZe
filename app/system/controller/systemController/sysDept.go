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

func DeptList(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	dept := new(systemModels.SysDeptDQL)
	c.ShouldBind(dept)
	dept.SetDataScope(loginUser, "d", "")
	list := iDept.SelectDeptList(dept)
	c.JSON(http.StatusOK, commonModels.SuccessData(list))

}

func DeptGetInfo(c *gin.Context) {
	deptId, err := strconv.ParseInt(c.Param("deptId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	menu := iDept.SelectDeptById(deptId)
	c.JSON(http.StatusOK, commonModels.SuccessData(menu))
}
func RoleDeptTreeselect(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
	}
	m := make(map[string]interface{})
	m["checkedKeys"] = iDept.SelectDeptListByRoleId(roleId)
	m["depts"] = iDept.SelectDeptList(new(systemModels.SysDeptDQL))
	c.JSON(http.StatusOK, commonModels.SuccessData(m))
}

func DeptAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysDept := new(systemModels.SysDeptDML)
	c.ShouldBind(sysDept)
	if iDept.CheckDeptNameUnique(sysDept) {
		c.JSON(http.StatusOK, commonModels.Waring("新增部门'"+sysDept.DeptName+"'失败，部门名称已存在"))
		return
	}

	sysDept.SetCreateBy(loginUser.User.UserName)
	iDept.InsertDept(sysDept)

	c.JSON(http.StatusOK, commonModels.Success())
}
func DeptEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysDept := new(systemModels.SysDeptDML)
	if iDept.CheckDeptNameUnique(sysDept) {
		c.JSON(http.StatusOK, commonModels.Waring("修改部门'"+sysDept.DeptName+"'失败，部门名称已存在"))
		return
	}
	c.ShouldBind(sysDept)
	sysDept.SetCreateBy(loginUser.User.UserName)
	iDept.UpdateDept(sysDept)
	c.JSON(http.StatusOK, commonModels.Success())
}
func DeptRemove(c *gin.Context) {
	deptId, err := strconv.ParseInt(c.Param("deptId"), 10, 64)
	if err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	if iDept.HasChildByDeptId(deptId) {
		c.JSON(http.StatusOK, commonModels.Waring("存在下级部门,不允许删除"))
		return
	}
	if iDept.CheckDeptExistUser(deptId) {
		c.JSON(http.StatusOK, commonModels.Waring("部门存在用户,不允许删除"))
		return
	}
	iDept.DeleteDeptById(deptId)

	c.JSON(http.StatusOK, commonModels.Success())
}
