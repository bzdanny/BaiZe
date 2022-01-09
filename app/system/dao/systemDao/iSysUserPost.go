package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
)

type IUserPostDao interface {
	BatchUserPost(users []*systemModels.SysUserPost, tx ...mysql.Transaction)
	DeleteUserPostByUserId(userId int64, tx ...mysql.Transaction)
	DeleteUserPost(ids []int64, tx ...mysql.Transaction)
}
