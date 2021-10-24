package systemDao

import (
	mysql "baize/app/common/mysql"
	"baize/app/system/models/systemModels"
	"github.com/jmoiron/sqlx"
)

func BatchUserPost(users []*systemModels.SysUserPost) {
	_, err := mysql.MysqlDb.NamedExec("insert into sys_user_post(user_id, post_id) values (:user_id,:post_id)", users)
	if err != nil {
		panic(err)
	}
}

func DeleteUserPostByUserId(userId int64) {
	_, err := mysql.MysqlDb.Exec("delete from sys_user_post where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}

func DeleteUserPost(ids []int64) {
	query, i, err := sqlx.In("delete from sys_user_post where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = mysql.MysqlDb.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
