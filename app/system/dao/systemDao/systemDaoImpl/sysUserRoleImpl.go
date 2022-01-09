package systemDaoImpl

import (
	"baize/app/common/datasource"
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

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRole(ids []int64, tx ...datasource.Transaction) {
	query, i, err := sqlx.In("delete from sys_user_role where user_id in(?)", ids)
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

func (sysUserRoleDao *sysUserRoleDao) BatchUserRole(users []*systemModels.SysUserRole, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec("insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleByUserId(userId int64, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
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
	err = datasource.GetMasterDb().Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleInfo(userRole *systemModels.SysUserRole) {
	_, err := datasource.GetMasterDb().NamedExec("delete from sys_user_role where user_id=:user_id and role_id=:role_id", userRole)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleInfos(roleId int64, userIds []int64) {
	query, i, err := sqlx.In("delete from sys_user_role where role_id=(?) and user_id in (?)", roleId, userIds)
	if err != nil {
		panic(err)
	}
	_, err = datasource.GetMasterDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
