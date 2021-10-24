package systemService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
	"baize/app/utils/snowflake"
	"strings"
)

/**
 * 查询所有岗位
 *
 * @return 岗位列表
 */
func SelectPostAll() (list []*systemModels.SysPostVo) {
	return systemDao.SelectPostAll()

}

/**
 * 根据用户ID获取岗位选择框列表
 *
 * @param userId 用户ID
 * @return 选中岗位ID列表
 */
func SelectPostListByUserId(userId int64) (list []int64) {
	return systemDao.SelectPostListByUserId(userId)

}

func SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, count *int64) {
	return systemDao.SelectPostList(post)

}

func SelectPostById(postId int64) (Post *systemModels.SysPostVo) {
	return systemDao.SelectPostById(postId)

}

func InsertPost(post *systemModels.SysPostDML) {
	post.PostId = snowflake.GenID()
	systemDao.InsertPost(post)
}

func UpdatePost(post *systemModels.SysPostDML) {
	systemDao.UpdatePost(post)
}
func DeletePostByIds(postId []int64) {
	systemDao.DeletePostByIds(postId)
	return
}
func SelectUserPostGroupByUserId(userId int64) string {

	return strings.Join(systemDao.SelectPostNameListByUserId(userId), ",")

}
