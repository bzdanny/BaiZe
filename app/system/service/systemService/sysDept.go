package systemService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
	"baize/app/utils/snowflake"
)

func SelectDeptList(dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo) {
	return systemDao.SelectDeptList(dept)

}

func SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo) {
	return systemDao.SelectDeptById(deptId)

}

func InsertDept(dept *systemModels.SysDeptDML) {
	dept.DeptId = snowflake.GenID()
	systemDao.InsertDept(dept)
	return
}

func UpdateDept(dept *systemModels.SysDeptDML) {
	systemDao.UpdateDept(dept)
	return
}
func DeleteDeptById(dept int64) {
	systemDao.DeleteDeptById(dept)
	return
}
func CheckDeptNameUnique(dept *systemModels.SysDeptDML) bool {
	deptId := systemDao.CheckDeptNameUnique(dept.DeptName, dept.ParentId)
	if deptId == dept.DeptId || deptId == 0 {
		return false
	}
	return true
}
func HasChildByDeptId(deptId int64) bool {
	return systemDao.HasChildByDeptId(deptId) > 0
}

func CheckDeptExistUser(deptId int64) bool {
	return systemDao.CheckDeptExistUser(deptId) > 0
}
