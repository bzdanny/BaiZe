package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

var sysUserRoleDaoImpl *sysUserRoleDao = &sysUserRoleDao{db: mysql.GetMysqlDb()}

type sysUserRoleDao struct {
	db **sqlx.DB
}

func GetSysUserRoleDao() *sysUserRoleDao {
	return sysUserRoleDaoImpl
}

func (sysUserRoleDao *sysUserRoleDao) getDb() *sqlx.DB {
	return *sysUserRoleDao.db
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRole(ids []int64) {
	query, i, err := sqlx.In("delete from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysUserRoleDao.getDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *sysUserRoleDao) BatchUserRole(users []*systemModels.SysUserRole) {
	_, err := sysUserRoleDao.getDb().NamedExec("insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleByUserId(userId int64) {
	_, err := sysUserRoleDao.getDb().Exec("delete from sys_user_role where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *sysUserRoleDao) CountUserRoleByRoleId(ids []int64) int {
	var count = 0
	query, i, err := sqlx.In("select count(*) from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	err = sysUserRoleDao.getDb().Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
