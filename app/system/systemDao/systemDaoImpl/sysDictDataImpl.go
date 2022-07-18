package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysDictDataDao struct {
	dictDataSql string
}

func NewSysDictDataDao() *SysDictDataDao {
	return &SysDictDataDao{
		dictDataSql: `select dict_code, dict_sort, dict_label, dict_value, dict_type, css_class, list_class, is_default , status , create_by, create_time, remark  from sys_dict_data`,
	}
}

func (sysDictDataDao *SysDictDataDao) SelectDictDataByType(db dataUtil.DB, dictType string) (SysDictDataList []*systemModels.SysDictDataVo) {
	whereSql := ` where status = '0' and dict_type = ? order by dict_sort asc`

	SysDictDataList = make([]*systemModels.SysDictDataVo, 0, 0)

	err := db.Select(&SysDictDataList, sysDictDataDao.dictDataSql+whereSql, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) SelectDictDataList(db dataUtil.DB, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total *int64) {
	whereSql := ``
	if dictData.DictType != "" {
		whereSql += " AND dict_type = :dict_type"
	}
	if dictData.Status != "" {
		whereSql += " AND  status = :status"
	}
	if dictData.DictLabel != "" {
		whereSql += " AND dict_label like concat('%', :dict_label, '%')"
	}

	if dictData.DataScope != "" {
		whereSql += " AND " + dictData.DataScope
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	return dataUtil.NamedQueryListAndTotal(db, list, dictData, sysDictDataDao.dictDataSql+whereSql, "", "")

}

func (sysDictDataDao *SysDictDataDao) SelectDictDataById(db dataUtil.DB, dictCode int64) (dictData *systemModels.SysDictDataVo) {

	dictData = new(systemModels.SysDictDataVo)
	err := db.Get(dictData, sysDictDataDao.dictDataSql+" where dict_code = ?", dictCode)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) InsertDictData(db dataUtil.DB, dictData *systemModels.SysDictDataAdd) {
	insertSQL := `insert into sys_dict_data(dict_code,dict_sort,dict_label,dict_value,dict_type,create_by,create_time,update_by,update_time %s)
					values(:dict_code,:dict_sort,:dict_label,:dict_value,:dict_type,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""

	if dictData.CssClass != "" {
		key += ",css_class"
		value += ",:css_class"
	}

	if dictData.ListClass != "" {
		key += ",list_class"
		value += ",:list_class"
	}

	if dictData.IsDefault != "" {
		key += ",is_default"
		value += ",:is_default"
	}

	if dictData.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if dictData.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) UpdateDictData(db dataUtil.DB, dictData *systemModels.SysDictDataEdit) {
	updateSQL := `update sys_dict_data set update_time = now() , update_by = :update_by`

	if dictData.DictSort != nil {
		updateSQL += ",dict_sort = :dict_sort"
	}

	if dictData.DictLabel != "" {
		updateSQL += ",dict_label = :dict_label"
	}
	if dictData.DictValue != "" {
		updateSQL += ",dict_value = :dict_value"
	}
	if dictData.DictType != "" {
		updateSQL += ",dict_type = :dict_type"
	}
	if dictData.CssClass != "" {
		updateSQL += ",css_class = :css_class"
	}
	if dictData.ListClass != "" {
		updateSQL += ",list_class = :list_class"
	}
	if dictData.IsDefault != "" {
		updateSQL += ",is_default = :is_default"
	}
	if dictData.Status != "" {
		updateSQL += ",status = :status"
	}
	if dictData.Status != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where dict_code = :dict_code"

	_, err := db.NamedExec(updateSQL, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) DeleteDictDataByIds(db dataUtil.DB, dictCodes []int64) {
	query, i, err := sqlx.In("delete from sys_dict_data where dict_code in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictDataDao *SysDictDataDao) CountDictDataByTypes(db dataUtil.DB, dictType []string) int {
	var count = 0
	query, i, err := sqlx.In("select count(*) from sys_dict_data where dict_type in(?)", dictType)
	if err != nil {
		panic(err)
	}
	err = db.Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
