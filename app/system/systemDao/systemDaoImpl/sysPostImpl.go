package systemDaoImpl

import (
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysPostDao struct {
	postSql string
}

func NewSysPostDao() *SysPostDao {
	return &SysPostDao{
		postSql: `select post_id, post_code, post_name, post_sort, status, create_by, create_time, remark  from sys_post`,
	}
}

func (postDao *SysPostDao) SelectPostAll(db dataUtil.DB) (sysPost []*systemModels.SysPostVo) {
	sysPost = make([]*systemModels.SysPostVo, 0)
	err := db.Select(&sysPost, postDao.postSql)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) SelectPostListByUserId(db dataUtil.DB, userId int64) (list []int64) {
	sqlStr := `select p.post_id
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := db.Select(&list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) SelectPostList(db dataUtil.DB, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total *int64) {
	whereSql := ``
	if post.PostCode != "" {
		whereSql += " AND post_code like concat('%', :post_code, '%')"
	}
	if post.Status != "" {
		whereSql += " AND  status = :status"
	}
	if post.PostName != "" {
		whereSql += " AND post_name like concat('%', :post_name, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	return dataUtil.NamedQueryListAndTotal(db, list, post, postDao.postSql+whereSql, "", "")

}

func (postDao *SysPostDao) SelectPostById(db dataUtil.DB, postId int64) (dictData *systemModels.SysPostVo) {

	dictData = new(systemModels.SysPostVo)
	err := db.Get(dictData, postDao.postSql+" where post_id = ?", postId)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) InsertPost(db dataUtil.DB, post *systemModels.SysPostAdd) {
	insertSQL := `insert into sys_post(post_id,post_code,post_name,create_by,create_time,update_by,update_time %s)
					values(:post_id,:post_code,:post_name,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""

	if post.PostSort != nil {
		key += ",post_sort"
		value += ",:post_sort"
	}

	if post.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if post.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) UpdatePost(db dataUtil.DB, post *systemModels.SysPostEdit) {
	updateSQL := `update sys_post set update_time = now() , update_by = :update_by`

	if post.PostCode != "" {
		updateSQL += ",post_code = :post_code"
	}

	if post.PostName != "" {
		updateSQL += ",post_name = :post_name"
	}
	if post.PostSort != nil {
		updateSQL += ",post_sort = :post_sort"
	}
	if post.Status != "" {
		updateSQL += ",status = :status"
	}
	if post.Status != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where post_id = :post_id"

	_, err := db.NamedExec(updateSQL, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) DeletePostByIds(db dataUtil.DB, dictCodes []int64) {
	query, i, err := sqlx.In("delete from sys_post where post_id in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (postDao *SysPostDao) SelectPostNameListByUserId(db dataUtil.DB, userId int64) (list []string) {
	sqlStr := `select p.post_name
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]string, 0, 1)
	err := db.Select(&list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
