package systemDao

import (
	mysql "baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

func DeleteUserRole(ids []int64) {
	query, i, err := sqlx.In("delete from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = mysql.MysqlDb.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func BatchUserRole(users []*systemModels.SysUserRole) {
	_, err := mysql.MysqlDb.NamedExec("insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func DeleteUserRoleByUserId(userId int64) {
	_, err := mysql.MysqlDb.Exec("delete from sys_user_role where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}
func CountUserRoleByRoleId(ids []int64) int {
	var count = 0
	query, i, err := sqlx.In("select count(*) from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	err = mysql.MysqlDb.Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
