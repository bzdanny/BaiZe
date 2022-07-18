package systemServiceImpl

import (
	systemDao2 "github.com/bzdanny/BaiZe/app/system/systemDao"
	systemDaoImpl2 "github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
)

type DeptService struct {
	data    *datasource.Data
	deptDao systemDao2.IDeptDao
	roleDao systemDao2.IRoleDao
}

func NewDeptService(data *datasource.Data, dd *systemDaoImpl2.SysDeptDao, rd *systemDaoImpl2.SysRoleDao) *DeptService {
	return &DeptService{data: data, deptDao: dd, roleDao: rd}
}

func (ds *DeptService) SelectDeptList(dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo) {
	return ds.deptDao.SelectDeptList(ds.data.GetSlaveDb(), dept)

}

func (ds *DeptService) SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo) {
	return ds.deptDao.SelectDeptById(ds.data.GetMasterDb(), deptId)

}

func (ds *DeptService) InsertDept(dept *systemModels.SysDeptAdd) {
	dept.DeptId = snowflake.GenID()
	ds.deptDao.InsertDept(ds.data.GetMasterDb(), dept)
	return
}

func (ds *DeptService) UpdateDept(dept *systemModels.SysDeptEdit) {
	ds.deptDao.UpdateDept(ds.data.GetMasterDb(), dept)
}
func (ds *DeptService) DeleteDeptById(dept int64) {
	ds.deptDao.DeleteDeptById(ds.data.GetMasterDb(), dept)
	return
}
func (ds *DeptService) CheckDeptNameUnique(id, parentId int64, deptName string) bool {
	deptId := ds.deptDao.CheckDeptNameUnique(ds.data.GetMasterDb(), deptName, parentId)
	if deptId == id || deptId == 0 {
		return false
	}
	return true
}
func (ds *DeptService) HasChildByDeptId(deptId int64) bool {
	return ds.deptDao.HasChildByDeptId(ds.data.GetMasterDb(), deptId) > 0
}

func (ds *DeptService) CheckDeptExistUser(deptId int64) bool {
	return ds.deptDao.CheckDeptExistUser(ds.data.GetMasterDb(), deptId) > 0
}
func (ds *DeptService) SelectDeptListByRoleId(roleId int64) []string {
	role := ds.roleDao.SelectRoleById(ds.data.GetMasterDb(), roleId)
	return ds.deptDao.SelectDeptListByRoleId(ds.data.GetMasterDb(), roleId, role.DeptCheckStrictly)
}
