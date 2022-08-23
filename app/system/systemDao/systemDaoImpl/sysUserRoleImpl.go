package systemDaoImpl

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysUserRoleDao struct {
}

func NewSysUserRoleDao() *SysUserRoleDao {
	return &SysUserRoleDao{}
}

func (sysUserRoleDao *SysUserRoleDao) DeleteUserRole(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("delete from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *SysUserRoleDao) BatchUserRole(db dataUtil.DB, users []*systemModels.SysUserRole) {

	_, err := db.NamedExec("insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *SysUserRoleDao) DeleteUserRoleByUserId(db dataUtil.DB, userId int64) {

	_, err := db.Exec("delete from sys_user_role where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *SysUserRoleDao) CountUserRoleByRoleId(db dataUtil.DB, ids []int64) int {
	var count = 0
	query, i, err := sqlx.In("select count(*) from sys_user_role where role_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	err = db.Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
func (sysUserRoleDao *SysUserRoleDao) DeleteUserRoleInfo(db dataUtil.DB, userRole *systemModels.SysUserRole) {
	_, err := db.NamedExec("delete from sys_user_role where user_id=:user_id and role_id=:role_id", userRole)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *SysUserRoleDao) DeleteUserRoleInfos(db dataUtil.DB, roleId int64, userIds []int64) {
	query, i, err := sqlx.In("delete from sys_user_role where role_id=(?) and user_id in (?)", roleId, userIds)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
