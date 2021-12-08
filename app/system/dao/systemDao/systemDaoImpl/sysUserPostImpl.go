package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

var sysUserPostDaoImpl *sysUserPostDao = &sysUserPostDao{db: mysql.GetMysqlDb()}

type sysUserPostDao struct {
	db **sqlx.DB
}

func GetSysUserPostDao() *sysUserPostDao {
	return sysUserPostDaoImpl
}

func (sysUserPostDao *sysUserPostDao) getDb() *sqlx.DB {
	return *sysUserPostDao.db
}

func (sysUserPostDao *sysUserPostDao) BatchUserPost(users []*systemModels.SysUserPost) {
	_, err := sysUserPostDao.getDb().NamedExec("insert into sys_user_post(user_id, post_id) values (:user_id,:post_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *sysUserPostDao) DeleteUserPostByUserId(userId int64) {
	_, err := sysUserPostDao.getDb().Exec("delete from sys_user_post where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *sysUserPostDao) DeleteUserPost(ids []int64) {
	query, i, err := sqlx.In("delete from sys_user_post where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysUserPostDao.getDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
