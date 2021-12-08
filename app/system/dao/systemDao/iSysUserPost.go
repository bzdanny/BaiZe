package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IUserPostDao interface {
	BatchUserPost(users []*systemModels.SysUserPost)
	DeleteUserPostByUserId(userId int64)
	DeleteUserPost(ids []int64)
}
