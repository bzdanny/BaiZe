package systemDao

import (
	mysql "baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func BatchRoleMenu(list []*systemModels.SysRoleMenu) {
	_, err := mysql.MysqlDb.NamedExec("insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func DeleteRoleMenu(ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_menu where role_id in", ids)
	if err != nil {
		panic(err)
	}
	_, err = mysql.MysqlDb.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func DeleteRoleMenuByRoleId(roleId int64) {
	_, err := mysql.MysqlDb.Exec("delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func CheckMenuExistRole(menuId int64) int {
	var count = 0
	err := mysql.MysqlDb.Get(&count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
