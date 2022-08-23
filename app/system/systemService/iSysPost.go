package systemService

import "github.com/bzdanny/BaiZe/app/system/systemModels"

type IPostService interface {
	SelectPostAll() (list []*systemModels.SysPostVo)
	PostExport(role *systemModels.SysPostDQL) (data []byte)
	SelectPostListByUserId(userId int64) (list []int64)
	SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, count *int64)
	SelectPostById(postId int64) (Post *systemModels.SysPostVo)
	InsertPost(post *systemModels.SysPostAdd)
	UpdatePost(post *systemModels.SysPostEdit)
	DeletePostByIds(postId []int64)
	SelectUserPostGroupByUserId(userId int64) string
}
