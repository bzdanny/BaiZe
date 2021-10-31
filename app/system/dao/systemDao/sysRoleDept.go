package systemDao

import (
	mysql "baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

func DeleteRoleDept(ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_dept where role_id in", ids)
	if err != nil {
		panic(err)
	}
	_, err = mysql.MysqlDb.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func DeleteRoleDeptByRoleId(id int64) {
	_, err := mysql.MysqlDb.Exec("delete from sys_role_dept where role_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func BatchRoleDept(list []*systemModels.SysRoleDept) {
	_, err := mysql.MysqlDb.NamedExec("insert into sys_role_dept(role_id, dept_id) values (:role_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
