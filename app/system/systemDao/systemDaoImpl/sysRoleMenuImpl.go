package systemDaoImpl

import (
	"database/sql"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysRoleMenuDao struct {
}

func NewSysRoleMenuDao() *SysRoleMenuDao {
	return &SysRoleMenuDao{}
}

func (sysRoleMenuDao *SysRoleMenuDao) BatchRoleMenu(db dataUtil.DB, list []*systemModels.SysRoleMenu) {

	_, err := db.NamedExec("insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *SysRoleMenuDao) DeleteRoleMenu(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_menu where role_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *SysRoleMenuDao) DeleteRoleMenuByRoleId(db dataUtil.DB, roleId int64) {

	_, err := db.Exec("delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *SysRoleMenuDao) CheckMenuExistRole(db dataUtil.DB, menuId int64) int {
	var count = 0
	err := db.Get(&count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
