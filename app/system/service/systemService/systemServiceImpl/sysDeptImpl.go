package systemServiceImpl

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/systemModels"
	"baize/app/utils/snowflake"
)

var deptServiceImpl *deptService = &deptService{deptDao: systemDaoImpl.GetSysDeptDao(), roleDao: systemDaoImpl.GetSysRoleDao()}

type deptService struct {
	deptDao systemDao.IDeptDao
	roleDao systemDao.IRoleDao
}

func GetDeptService() *deptService {
	return deptServiceImpl
}

func (deptService *deptService) SelectDeptList(dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo) {
	return deptService.deptDao.SelectDeptList(dept)

}

func (deptService *deptService) SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo) {
	return deptService.deptDao.SelectDeptById(deptId)

}

func (deptService *deptService) InsertDept(dept *systemModels.SysDeptDML) {
	dept.DeptId = snowflake.GenID()
	deptService.deptDao.InsertDept(dept)
	return
}

func (deptService *deptService) UpdateDept(dept *systemModels.SysDeptDML) {
	deptService.deptDao.UpdateDept(dept)
	return
}
func (deptService *deptService) DeleteDeptById(dept int64) {
	deptService.deptDao.DeleteDeptById(dept)
	return
}
func (deptService *deptService) CheckDeptNameUnique(dept *systemModels.SysDeptDML) bool {
	deptId := deptService.deptDao.CheckDeptNameUnique(dept.DeptName, dept.ParentId)
	if deptId == dept.DeptId || deptId == 0 {
		return false
	}
	return true
}
func (deptService *deptService) HasChildByDeptId(deptId int64) bool {
	return deptService.deptDao.HasChildByDeptId(deptId) > 0
}

func (deptService *deptService) CheckDeptExistUser(deptId int64) bool {
	return deptService.deptDao.CheckDeptExistUser(deptId) > 0
}
func (deptService *deptService) SelectDeptListByRoleId(roleId int64) []string {
	role := deptService.roleDao.SelectRoleById(roleId)
	return deptService.deptDao.SelectDeptListByRoleId(roleId, role.DeptCheckStrictly)
}
