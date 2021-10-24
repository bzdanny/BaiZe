package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
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
	list := systemService.SelectDeptList(dept)
	c.JSON(http.StatusOK, commonModels.SuccessData(list))

}

func DeptGetInfo(c *gin.Context) {
	deptId, err := strconv.ParseInt(c.Param("deptId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	menu := systemService.SelectDeptById(deptId)
	c.JSON(http.StatusOK, commonModels.SuccessData(menu))
}

func DeptAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysDept := new(systemModels.SysDeptDML)
	c.ShouldBind(sysDept)
	if systemService.CheckDeptNameUnique(sysDept) {
		c.JSON(http.StatusOK, commonModels.Waring("新增部门'"+sysDept.DeptName+"'失败，部门名称已存在"))
		return
	}

	sysDept.SetCreateBy(loginUser.User.UserName)
	systemService.InsertDept(sysDept)

	c.JSON(http.StatusOK, commonModels.Success())
}
func DeptEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysDept := new(systemModels.SysDeptDML)
	if systemService.CheckDeptNameUnique(sysDept) {
		c.JSON(http.StatusOK, commonModels.Waring("修改部门'"+sysDept.DeptName+"'失败，部门名称已存在"))
		return
	}
	c.ShouldBind(sysDept)
	sysDept.SetCreateBy(loginUser.User.UserName)
	systemService.UpdateDept(sysDept)
	c.JSON(http.StatusOK, commonModels.Success())
}
func DeptRemove(c *gin.Context) {
	deptId, err := strconv.ParseInt(c.Param("deptId"), 10, 64)
	if err != nil {
		zap.L().Error("登录参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	if systemService.HasChildByDeptId(deptId) {
		c.JSON(http.StatusOK, commonModels.Waring("存在下级部门,不允许删除"))
		return
	}
	if systemService.CheckDeptExistUser(deptId) {
		c.JSON(http.StatusOK, commonModels.Waring("部门存在用户,不允许删除"))
		return
	}
	systemService.DeleteDeptById(deptId)

	c.JSON(http.StatusOK, commonModels.Success())
}
