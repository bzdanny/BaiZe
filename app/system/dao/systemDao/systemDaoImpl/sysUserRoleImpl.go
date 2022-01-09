package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

var sysUserRoleDaoImpl *sysUserRoleDao

type sysUserRoleDao struct {
}

func init() {
	sysUserRoleDaoImpl = &sysUserRoleDao{}
}

func GetSysUserRoleDao() *sysUserRoleDao {
	return sysUserRoleDaoImpl
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRole(ids []int64, tx ...mysql.Transaction) {
	query, i, err := sqlx.In("delete from sys_user_role where user_id in(?)", ids)
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

func (sysUserRoleDao *sysUserRoleDao) BatchUserRole(users []*systemModels.SysUserRole, tx ...mysql.Transaction) {
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err := db.NamedExec("insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleByUserId(userId int64, tx ...mysql.Transaction) {
	var db mysql.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = mysql.GetMasterMysqlDb()
	}
	_, err := db.Exec("delete from sys_user_role where user_id= ?", userId)
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
	err = mysql.GetMasterMysqlDb().Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleInfo(userRole *systemModels.SysUserRole) {
	_, err := mysql.GetMasterMysqlDb().NamedExec("delete from sys_user_role where user_id=:user_id and role_id=:role_id", userRole)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleInfos(roleId int64, userIds []int64) {
	query, i, err := sqlx.In("delete from sys_user_role where role_id=(?) and user_id in (?)", roleId, userIds)
	if err != nil {
		panic(err)
	}
	_, err = mysql.GetMasterMysqlDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
