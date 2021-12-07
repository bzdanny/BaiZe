package systemDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/constant/constants"
	"baize/app/system/models/systemModels"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var sysPostDaoImpl *sysPostDao = &sysPostDao{db: mysql.GetMysqlDb()}

type sysPostDao struct {
	db **sqlx.DB
}

func GetSysPostDao() *sysPostDao {
	return sysPostDaoImpl
}

func (postDao *sysPostDao) getDb() *sqlx.DB {
	return *postDao.db
}

var selectPostSql = `select post_id, post_code, post_name, post_sort, status, create_by, create_time, remark `
var fromPostSql = ` from sys_post`

func (postDao *sysPostDao) SelectPostAll() (sysPost []*systemModels.SysPostVo) {
	sysPost = make([]*systemModels.SysPostVo, 0, 2)
	err := postDao.getDb().Select(&sysPost, selectPostSql+fromPostSql)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) SelectPostListByUserId(userId int64) (list []int64) {
	sqlStr := `select p.post_id
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := postDao.getDb().Select(&list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) SelectPostList(post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total *int64) {
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
	countSql := constants.MysqlCount + fromPostSql + whereSql

	countRow, err := postDao.getDb().NamedQuery(countSql, post)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	list = make([]*systemModels.SysPostVo, 0, post.Size)
	if *total > post.Offset {
		if post.Limit != "" {
			whereSql += post.Limit
		}
		listRows, err := postDao.getDb().NamedQuery(selectPostSql+fromPostSql+whereSql, post)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			postVo := new(systemModels.SysPostVo)
			listRows.StructScan(postVo)
			list = append(list, postVo)
		}
		defer listRows.Close()
	}
	return
}

func (postDao *sysPostDao) SelectPostById(postId int64) (dictData *systemModels.SysPostVo) {

	dictData = new(systemModels.SysPostVo)
	err := postDao.getDb().Get(dictData, selectPostSql+fromPostSql+" where post_id = ?", postId)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) InsertPost(post *systemModels.SysPostDML) {
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
	_, err := postDao.getDb().NamedExec(insertStr, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) UpdatePost(post *systemModels.SysPostDML) {
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

	_, err := postDao.getDb().NamedExec(updateSQL, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) DeletePostByIds(dictCodes []int64) {
	query, i, err := sqlx.In("delete from sys_post where post_id in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = postDao.getDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (postDao *sysPostDao) SelectPostNameListByUserId(userId int64) (list []string) {
	sqlStr := `select p.post_name
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]string, 0, 1)
	err := postDao.getDb().Select(&list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
