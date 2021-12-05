package systemService

import (
	"baize/app/system/models/systemModels"
)

type IPostService interface {
	SelectPostAll() (list []*systemModels.SysPostVo)
	SelectPostListByUserId(userId int64) (list []int64)
	SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, count *int64)
	SelectPostById(postId int64) (Post *systemModels.SysPostVo)
	InsertPost(post *systemModels.SysPostDML)
	UpdatePost(post *systemModels.SysPostDML)
	DeletePostByIds(postId []int64)
	SelectUserPostGroupByUserId(userId int64) string
}
