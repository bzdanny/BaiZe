package systemService

import (
	"baize/app/system/models/systemModels"
)

type IDeptService interface {
	SelectDeptList(dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo)
	SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(dept *systemModels.SysDeptDML)
	UpdateDept(dept *systemModels.SysDeptDML)
	DeleteDeptById(dept int64)
	CheckDeptNameUnique(dept *systemModels.SysDeptDML) bool
	HasChildByDeptId(deptId int64) bool
	CheckDeptExistUser(deptId int64) bool
	SelectDeptListByRoleId(roleId int64) []string
}
