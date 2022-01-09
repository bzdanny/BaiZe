package systemDaoImpl

import (
	"baize/app/common/datasource"
	"baize/app/constant/constants"
	"baize/app/system/models/systemModels"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var sysDictTypeDaoImpl *sysDictTypeDao

func init() {
	sysDictTypeDaoImpl = &sysDictTypeDao{
		selectDictTypeSql: `select dict_id, dict_name, dict_type, status, create_by, create_time, remark  `,
		fromDictTypeSql:   ` from sys_dict_type`,
	}
}

type sysDictTypeDao struct {
	selectDictTypeSql string
	fromDictTypeSql   string
}

func GetSysDictTypeDao() *sysDictTypeDao {
	return sysDictTypeDaoImpl
}

func (sysDictTypeDao *sysDictTypeDao) SelectDictTypeList(dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total *int64) {
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

	countRow, err := datasource.GetMasterDb().NamedQuery(constants.MysqlCount+sysDictTypeDao.fromDictTypeSql+whereSql, dictType)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	list = make([]*systemModels.SysDictTypeVo, 0, dictType.Size)
	if *total > dictType.Offset {
		if dictType.Limit != "" {
			whereSql += dictType.Limit
		}
		listRows, err := datasource.GetMasterDb().NamedQuery(sysDictTypeDao.selectDictTypeSql+sysDictTypeDao.fromDictTypeSql+whereSql, dictType)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			dictTypeVo := new(systemModels.SysDictTypeVo)
			listRows.StructScan(dictTypeVo)
			list = append(list, dictTypeVo)
		}
		defer listRows.Close()
	}
	return
}

func (sysDictTypeDao *sysDictTypeDao) SelectDictTypeAll() (list []*systemModels.SysDictTypeVo) {

	list = make([]*systemModels.SysDictTypeVo, 0, 2)
	err := datasource.GetMasterDb().Select(&list, sysDictTypeDao.selectDictTypeSql+sysDictTypeDao.fromDictTypeSql)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *sysDictTypeDao) SelectDictTypeById(dictId int64) (dictType *systemModels.SysDictTypeVo) {

	dictType = new(systemModels.SysDictTypeVo)
	err := datasource.GetMasterDb().Get(dictType, sysDictTypeDao.selectDictTypeSql+sysDictTypeDao.fromDictTypeSql+" where dict_id = ?", dictId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *sysDictTypeDao) SelectDictTypeByIds(dictId []int64) (dictTypes []string) {

	dictTypes = make([]string, 0, 2)

	query, i, err := sqlx.In("select dict_type from sys_dict_type where dict_id in(?)", dictId)
	if err != nil {
		panic(err)
	}
	return
	err = datasource.GetMasterDb().Select(&dictTypes, query, i)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *sysDictTypeDao) InsertDictType(dictType *systemModels.SysDictTypeDML) {
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
	_, err := datasource.GetMasterDb().NamedExec(insertStr, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *sysDictTypeDao) UpdateDictType(dictType *systemModels.SysDictTypeDML) {
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

	_, err := datasource.GetMasterDb().NamedExec(updateSQL, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *sysDictTypeDao) DeleteDictTypeByIds(dictIds []int64) (err error) {
	query, i, err := sqlx.In("delete from sys_dict_type where dict_id in (?)", dictIds)
	if err != nil {
		panic(err)
	}
	_, err = datasource.GetMasterDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictTypeDao *sysDictTypeDao) CheckDictTypeUnique(dictType string) int64 {
	var dictId int64 = 0
	err := datasource.GetMasterDb().Get(&dictId, "select dict_id from sys_dict_type where dict_type = ?", dictType)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return dictId
}
