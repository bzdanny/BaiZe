package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysDictTypeDao struct {
	dictTypeSql string
}

func NewSysDictTypeDao() *SysDictTypeDao {
	return &SysDictTypeDao{
		dictTypeSql: `select dict_id, dict_name, dict_type, status, create_by, create_time, remark   from sys_dict_type`,
	}
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeList(db dataUtil.DB, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total *int64) {
	whereSql := ``
	if dictType.DictName != "" {
		whereSql += " AND dict_name like concat('%', :dictName, '%')"
	}
	if dictType.Status != "" {
		whereSql += " AND  status = :status"
	}
	if dictType.DictType != "" {
		whereSql += " AND dict_type like concat('%', :dictType, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	return dataUtil.NamedQueryListAndTotal(db, list, dictType, sysDictTypeDao.dictTypeSql+whereSql, "", "")

}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeAll(db dataUtil.DB) (list []*systemModels.SysDictTypeVo) {

	list = make([]*systemModels.SysDictTypeVo, 0)
	err := db.Select(&list, sysDictTypeDao.dictTypeSql)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeById(db dataUtil.DB, dictId int64) (dictType *systemModels.SysDictTypeVo) {

	dictType = new(systemModels.SysDictTypeVo)
	err := db.Get(dictType, sysDictTypeDao.dictTypeSql+" where dict_id = ?", dictId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeByIds(db dataUtil.DB, dictId []int64) (dictTypes []string) {
	dictTypes = make([]string, 0)
	query, args, err := sqlx.In("select dict_type from sys_dict_type where dict_id in(?)", dictId)
	if err != nil {
		panic(err)
	}

	err = db.Select(&dictTypes, query, args...)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) InsertDictType(db dataUtil.DB, dictType *systemModels.SysDictTypeAdd) {
	insertSQL := `insert into sys_dict_type(dict_id,dict_name,dict_type,create_by,create_time,update_by,update_time %s)
					values(:dict_id,:dict_name,:dict_type,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""

	if dictType.Status != "" {
		key += ",status"
		value += ",:status"
	}

	if dictType.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) UpdateDictType(db dataUtil.DB, dictType *systemModels.SysDictTypeEdit) {
	updateSQL := `update sys_dict_type set update_time = now() , update_by = :update_by`

	if dictType.DictName != "" {
		updateSQL += ",dict_name = :dict_name"
	}
	if dictType.DictType != "" {
		updateSQL += ",dict_type = :dict_type"
	}
	if dictType.Status != "" {
		updateSQL += ",status = :status"
	}
	if dictType.Remark != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where dict_id = :dict_id"

	_, err := db.NamedExec(updateSQL, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) DeleteDictTypeByIds(db dataUtil.DB, dictIds []int64) (err error) {
	query, i, err := sqlx.In("delete from sys_dict_type where dict_id in (?)", dictIds)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictTypeDao *SysDictTypeDao) CheckDictTypeUnique(db dataUtil.DB, dictType string) int64 {
	var dictId int64 = 0
	err := db.Get(&dictId, "select dict_id from sys_dict_type where dict_type = ?", dictType)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return dictId
}
