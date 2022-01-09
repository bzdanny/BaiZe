package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

var sysRoleDeptDaoImpl *sysRoleDeptDao

func init() {
	sysRoleDeptDaoImpl = &sysRoleDeptDao{}
}

type sysRoleDeptDao struct {
}

func GetSysRoleDeptDao() *sysRoleDeptDao {
	return sysRoleDeptDaoImpl
}

func (sysRoleDeptDao *sysRoleDeptDao) DeleteRoleDept(ids []int64, tx ...mysql.Transaction) {
	query, i, err := sqlx.In("delete from sys_role_dept where role_id in", ids)
	if err != nil {
		panic(err)
	}
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDeptDao *sysRoleDeptDao) DeleteRoleDeptByRoleId(id int64, tx ...mysql.Transaction) {
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err := db.Exec("delete from sys_role_dept where role_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func (sysRoleDeptDao *sysRoleDeptDao) BatchRoleDept(list []*systemModels.SysRoleDept, tx ...mysql.Transaction) {
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err := db.NamedExec("insert into sys_role_dept(role_id, dept_id) values (:role_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
