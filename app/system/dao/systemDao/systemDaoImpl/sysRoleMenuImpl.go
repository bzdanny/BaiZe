package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

var sysRoleMenuDaoImpl *sysRoleMenuDao = &sysRoleMenuDao{db: mysql.GetMysqlDb()}

type sysRoleMenuDao struct {
	db **sqlx.DB
}

func GetSysRoleMenuDao() *sysRoleMenuDao {
	return sysRoleMenuDaoImpl
}

func (sysRoleMenuDao *sysRoleMenuDao) getDb() *sqlx.DB {
	return *sysRoleMenuDao.db
}

func (sysRoleMenuDao *sysRoleMenuDao) BatchRoleMenu(list []*systemModels.SysRoleMenu) {
	_, err := sysRoleMenuDao.getDb().NamedExec("insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenu(ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_menu where role_id in", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysRoleMenuDao.getDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenuByRoleId(roleId int64) {
	_, err := sysRoleMenuDao.getDb().Exec("delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) CheckMenuExistRole(menuId int64) int {
	var count = 0
	err := sysRoleMenuDao.getDb().Get(&count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
