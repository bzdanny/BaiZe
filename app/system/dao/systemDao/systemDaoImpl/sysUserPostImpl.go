package systemDaoImpl

import (
	"baize/app/common/datasource"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

var sysUserPostDaoImpl *sysUserPostDao

func init() {
	sysUserPostDaoImpl = &sysUserPostDao{}
}

type sysUserPostDao struct {
}

func GetSysUserPostDao() *sysUserPostDao {
	return sysUserPostDaoImpl
}

func (sysUserPostDao *sysUserPostDao) BatchUserPost(users []*systemModels.SysUserPost, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec("insert into sys_user_post(user_id, post_id) values (:user_id,:post_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *sysUserPostDao) DeleteUserPostByUserId(userId int64, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.Exec("delete from sys_user_post where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *sysUserPostDao) DeleteUserPost(ids []int64, tx ...datasource.Transaction) {
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	query, i, err := sqlx.In("delete from sys_user_post where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
