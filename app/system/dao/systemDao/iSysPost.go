package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IPostDao interface {
	SelectPostAll() (sysPost []*systemModels.SysPostVo)
	SelectPostListByUserId(userId int64) (list []int64)
	SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total *int64)
	SelectPostById(postId int64) (dictData *systemModels.SysPostVo)
	InsertPost(post *systemModels.SysPostDML)
	UpdatePost(post *systemModels.SysPostDML)
	DeletePostByIds(dictCodes []int64)
	SelectPostNameListByUserId(userId int64) (list []string)
}
