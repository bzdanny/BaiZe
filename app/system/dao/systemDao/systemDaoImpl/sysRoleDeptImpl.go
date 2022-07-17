package systemDaoImpl

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysRoleDeptDao struct {
}

func NewSysRoleDeptDao() *SysRoleDeptDao {
	return &SysRoleDeptDao{}
}

func (sysRoleDeptDao *SysRoleDeptDao) DeleteRoleDept(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_dept where role_id in", ids)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDeptDao *SysRoleDeptDao) DeleteRoleDeptByRoleId(db dataUtil.DB, id int64) {

	_, err := db.Exec("delete from sys_role_dept where role_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func (sysRoleDeptDao *SysRoleDeptDao) BatchRoleDept(db dataUtil.DB, list []*systemModels.SysRoleDept) {

	_, err := db.NamedExec("insert into sys_role_dept(role_id, dept_id) values (:role_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
