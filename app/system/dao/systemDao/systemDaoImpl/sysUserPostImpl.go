package systemDaoImpl

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysUserPostDao struct {
}

func NewSysUserPostDao() *SysUserPostDao {
	return &SysUserPostDao{}
}

func (sysUserPostDao *SysUserPostDao) BatchUserPost(db dataUtil.DB, users []*systemModels.SysUserPost) {

	_, err := db.NamedExec("insert into sys_user_post(user_id, post_id) values (:user_id,:post_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *SysUserPostDao) DeleteUserPostByUserId(db dataUtil.DB, userId int64) {

	_, err := db.Exec("delete from sys_user_post where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *SysUserPostDao) DeleteUserPost(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("delete from sys_user_post where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
