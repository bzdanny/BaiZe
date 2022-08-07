package systemDaoImpl

import (
	"database/sql"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysRolePermissionDao struct {
}

func NewSysRolePermissionDao() *SysRolePermissionDao {
	return &SysRolePermissionDao{}
}

func (sysRolePermissionDao *SysRolePermissionDao) BatchRolePermission(db dataUtil.DB, list []*systemModels.SysRolePermission) {

	_, err := db.NamedExec("insert into sys_role_Permission(role_id, permission_id) values (:role_id,:permission_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRolePermissionDao *SysRolePermissionDao) DeleteRolePermission(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("delete from sys_role_permission where role_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRolePermissionDao *SysRolePermissionDao) DeleteRolePermissionByRoleId(db dataUtil.DB, roleId int64) {

	_, err := db.Exec("delete from sys_role_permission where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRolePermissionDao *SysRolePermissionDao) CheckPermissionExistRole(db dataUtil.DB, permissionId int64) int {
	var count = 0
	err := db.Get(&count, "select count(1) from sys_role_permission where permission_id = ?", permissionId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
