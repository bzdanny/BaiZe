package systemDaoImpl

import (
	"baize/app/common/datasource"
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

func (sysRoleMenuDao *sysRoleMenuDao) BatchRoleMenu(list []*systemModels.SysRoleMenu, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec("insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenu(ids []int64, tx ...datasource.Transaction) {
	query, i, err := sqlx.In("delete from sys_role_menu where role_id in", ids)
	if err != nil {
		panic(err)
	}
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenuByRoleId(roleId int64, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.Exec("delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) CheckMenuExistRole(menuId int64) int {
	var count = 0
	err := datasource.GetMasterDb().Get(&count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
