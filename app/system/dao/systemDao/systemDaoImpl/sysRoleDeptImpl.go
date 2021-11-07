package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

var sysRoleDeptDaoImpl *sysRoleDeptDao = &sysRoleDeptDao{db: mysql.GetMysqlDb()}

type sysRoleDeptDao struct {
	db **sqlx.DB
}

func GetSysRoleDeptDao() *sysRoleDeptDao {
	return sysRoleDeptDaoImpl
}
func (sysRoleDeptDao *sysRoleDeptDao) getDb() *sqlx.DB {
	return *sysRoleDeptDao.db
}

func (sysRoleDeptDao *sysRoleDeptDao) DeleteRoleDept(ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_dept where role_id in", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysRoleDeptDao.getDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDeptDao *sysRoleDeptDao) DeleteRoleDeptByRoleId(id int64) {
	_, err := sysRoleDeptDao.getDb().Exec("delete from sys_role_dept where role_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func (sysRoleDeptDao *sysRoleDeptDao) BatchRoleDept(list []*systemModels.SysRoleDept) {
	_, err := sysRoleDeptDao.getDb().NamedExec("insert into sys_role_dept(role_id, dept_id) values (:role_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
