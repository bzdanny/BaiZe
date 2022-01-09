package systemDao

import (
	"baize/app/common/datasource"
	"baize/app/system/models/systemModels"
)

type IUserPostDao interface {
	BatchUserPost(users []*systemModels.SysUserPost, tx ...datasource.Transaction)
	DeleteUserPostByUserId(userId int64, tx ...datasource.Transaction)
	DeleteUserPost(ids []int64, tx ...datasource.Transaction)
}
