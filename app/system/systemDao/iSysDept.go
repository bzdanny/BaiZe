package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IDeptDao interface {
	SelectDeptList(db dataUtil.DB, dept *systemModels.SysDeptDQL) (sysDeptList []*systemModels.SysDeptVo)
	SelectDeptById(db dataUtil.DB, deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(db dataUtil.DB, dept *systemModels.SysDeptAdd)
	UpdateDept(db dataUtil.DB, dept *systemModels.SysDeptEdit)
	DeleteDeptById(db dataUtil.DB, deptId int64)
	CheckDeptNameUnique(db dataUtil.DB, deptName string, parentId int64) int64
	HasChildByDeptId(db dataUtil.DB, deptId int64) int
	CheckDeptExistUser(db dataUtil.DB, deptId int64) int
	SelectDeptListByRoleId(db dataUtil.DB, roleId int64, deptCheckStrictly bool) (deptIds []string)
}
