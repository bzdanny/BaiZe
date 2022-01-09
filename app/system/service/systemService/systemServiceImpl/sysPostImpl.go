package systemServiceImpl

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/systemModels"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
	"strings"
)

var postServiceImpl *postService

func init() {
	postServiceImpl = &postService{postDao: systemDaoImpl.GetSysPostDao()}
}

type postService struct {
	postDao systemDao.IPostDao
}

func GetPostService() *postService {
	return postServiceImpl
}

//SelectPostAll 查询所有岗位
//@return 岗位列表
func (postService *postService) SelectPostAll() (list []*systemModels.SysPostVo) {
	return postService.postDao.SelectPostAll()

}

//SelectPostListByUserId 根据用户ID获取岗位选择框列表
//@param userId 用户ID
//@return 选中岗位ID列表
func (postService *postService) SelectPostListByUserId(userId int64) (list []int64) {
	return postService.postDao.SelectPostListByUserId(userId)

}

func (postService *postService) SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, count *int64) {
	return postService.postDao.SelectPostList(post)

}
func (postService *postService) PostExport(post *systemModels.SysPostDQL) (data []byte) {
	list, _ := postService.postDao.SelectPostList(post)
	rows := systemModels.SysPostListToRows(list)
	return exceLize.SetRows(rows)

}

func (postService *postService) SelectPostById(postId int64) (Post *systemModels.SysPostVo) {
	return postService.postDao.SelectPostById(postId)

}

func (postService *postService) InsertPost(post *systemModels.SysPostDML) {
	post.PostId = snowflake.GenID()
	postService.postDao.InsertPost(post)
}

func (postService *postService) UpdatePost(post *systemModels.SysPostDML) {
	postService.postDao.UpdatePost(post)
}
func (postService *postService) DeletePostByIds(postId []int64) {
	postService.postDao.DeletePostByIds(postId)
	return
}
func (postService *postService) SelectUserPostGroupByUserId(userId int64) string {

	return strings.Join(postService.postDao.SelectPostNameListByUserId(userId), ",")

}
