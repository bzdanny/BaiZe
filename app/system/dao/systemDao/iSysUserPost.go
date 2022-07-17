package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IUserPostDao interface {
	BatchUserPost(db dataUtil.DB, users []*systemModels.SysUserPost)
	DeleteUserPostByUserId(db dataUtil.DB, userId int64)
	DeleteUserPost(db dataUtil.DB, ids []int64)
}
