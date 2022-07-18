package systemService

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
)

type IDeptService interface {
	SelectDeptList(dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo)
	SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(dept *systemModels.SysDeptAdd)
	UpdateDept(dept *systemModels.SysDeptEdit)
	DeleteDeptById(dept int64)
	CheckDeptNameUnique(id, parentId int64, deptName string) bool
	HasChildByDeptId(deptId int64) bool
	CheckDeptExistUser(deptId int64) bool
	SelectDeptListByRoleId(roleId int64) []string
}
