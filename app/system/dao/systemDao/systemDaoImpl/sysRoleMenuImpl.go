package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

var sysRoleMenuDaoImpl *sysRoleMenuDao

func init() {
	sysRoleMenuDaoImpl = &sysRoleMenuDao{}
}

type sysRoleMenuDao struct {
}

func GetSysRoleMenuDao() *sysRoleMenuDao {
	return sysRoleMenuDaoImpl
}

func (sysRoleMenuDao *sysRoleMenuDao) BatchRoleMenu(list []*systemModels.SysRoleMenu, tx ...mysql.Transaction) {
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err := db.NamedExec("insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenu(ids []int64, tx ...mysql.Transaction) {
	query, i, err := sqlx.In("delete from sys_role_menu where role_id in", ids)
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
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenuByRoleId(roleId int64, tx ...mysql.Transaction) {
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err := db.Exec("delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) CheckMenuExistRole(menuId int64) int {
	var count = 0
	err := mysql.GetMasterMysqlDb().Get(&count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
