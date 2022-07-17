package systemServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/system/dao/systemDao"
	"github.com/bzdanny/BaiZe/app/system/dao/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/exceLize"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
	"strings"
)

type PostService struct {
	data    *datasource.Data
	postDao systemDao.IPostDao
}

func NewPostService(data *datasource.Data, pd *systemDaoImpl.SysPostDao) *PostService {
	return &PostService{
		data:    data,
		postDao: pd,
	}
}

//SelectPostAll 查询所有岗位
//@return 岗位列表
func (postService *PostService) SelectPostAll() (list []*systemModels.SysPostVo) {
	return postService.postDao.SelectPostAll(postService.data.GetSlaveDb())

}

//SelectPostListByUserId 根据用户ID获取岗位选择框列表
//@param userId 用户ID
//@return 选中岗位ID列表
func (postService *PostService) SelectPostListByUserId(userId int64) (list []int64) {
	return postService.postDao.SelectPostListByUserId(postService.data.GetSlaveDb(), userId)

}

func (postService *PostService) SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, count *int64) {
	return postService.postDao.SelectPostList(postService.data.GetSlaveDb(), post)

}
func (postService *PostService) PostExport(post *systemModels.SysPostDQL) (data []byte) {
	list, _ := postService.postDao.SelectPostList(postService.data.GetSlaveDb(), post)
	rows := systemModels.SysPostListToRows(list)
	return exceLize.SetRows(rows)

}

func (postService *PostService) SelectPostById(postId int64) (Post *systemModels.SysPostVo) {
	return postService.postDao.SelectPostById(postService.data.GetSlaveDb(), postId)

}

func (postService *PostService) InsertPost(post *systemModels.SysPostAdd) {
	post.PostId = snowflake.GenID()
	postService.postDao.InsertPost(postService.data.GetMasterDb(), post)
}

func (postService *PostService) UpdatePost(post *systemModels.SysPostEdit) {
	postService.postDao.UpdatePost(postService.data.GetMasterDb(), post)
}
func (postService *PostService) DeletePostByIds(postId []int64) {
	postService.postDao.DeletePostByIds(postService.data.GetSlaveDb(), postId)
	return
}
func (postService *PostService) SelectUserPostGroupByUserId(userId int64) string {

	return strings.Join(postService.postDao.SelectPostNameListByUserId(postService.data.GetSlaveDb(), userId), ",")

}
