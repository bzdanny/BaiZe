package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IDeptDao interface {
	SelectDeptList(dept *systemModels.SysDeptDQL) (sysDeptList []*systemModels.SysDeptVo)
	SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(dept *systemModels.SysDeptDML)
	UpdateDept(dept *systemModels.SysDeptDML)
	DeleteDeptById(deptId int64)
	CheckDeptNameUnique(deptName string, parentId int64) int64
	HasChildByDeptId(deptId int64) int
	CheckDeptExistUser(deptId int64) int
	SelectDeptListByRoleId(roleId int64, deptCheckStrictly bool) (deptIds []string)
}
