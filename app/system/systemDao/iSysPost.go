package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IPostDao interface {
	SelectPostAll(db dataUtil.DB) (sysPost []*systemModels.SysPostVo)
	SelectPostListByUserId(db dataUtil.DB, userId int64) (list []int64)
	SelectPostList(db dataUtil.DB, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total *int64)
	SelectPostById(db dataUtil.DB, postId int64) (dictData *systemModels.SysPostVo)
	InsertPost(db dataUtil.DB, post *systemModels.SysPostAdd)
	UpdatePost(db dataUtil.DB, post *systemModels.SysPostEdit)
	DeletePostByIds(db dataUtil.DB, dictCodes []int64)
	SelectPostNameListByUserId(db dataUtil.DB, userId int64) (list []string)
}
