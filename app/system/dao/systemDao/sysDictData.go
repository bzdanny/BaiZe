package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/constant/constants"
	"baize/app/system/models/systemModels"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var selectDictDataSql = `select dict_code, dict_sort, dict_label, dict_value, dict_type, css_class, list_class, is_default , status , create_by, create_time, remark `
var fromDictDataSql = ` from sys_dict_data`

func SelectDictDataByType(dictType string) (SysDictDataList []*systemModels.SysDictDataVo) {
	whereSql := ` where status = '0' and dict_type = ? order by dict_sort asc`

	SysDictDataList = make([]*systemModels.SysDictDataVo, 0, 0)

	err := mysql.MysqlDb.Select(&SysDictDataList, selectDictDataSql+fromDictDataSql+whereSql, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func SelectDictDataList(dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total *int64) {
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
	countSql := constants.MysqlCount + fromDictDataSql + whereSql

	countRow, err := mysql.MysqlDb.NamedQuery(countSql, dictData)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	list = make([]*systemModels.SysDictDataVo, 0, dictData.Size)
	if *total > dictData.Offset {
		if dictData.Limit != "" {
			whereSql += dictData.Limit
		}
		listRows, err := mysql.MysqlDb.NamedQuery(selectDictDataSql+fromDictDataSql+whereSql, dictData)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			dictDataVo := new(systemModels.SysDictDataVo)
			listRows.StructScan(dictDataVo)
			list = append(list, dictDataVo)
		}
		defer listRows.Close()
	}
	return
}

func SelectDictDataById(dictCode int64) (dictData *systemModels.SysDictDataVo) {

	dictData = new(systemModels.SysDictDataVo)
	err := mysql.MysqlDb.Get(dictData, selectDictDataSql+fromDictDataSql+" where dict_code = ？", dictCode)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func InsertDictData(dictData *systemModels.SysDictDataDML) {
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
	_, err := mysql.MysqlDb.NamedExec(insertStr, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func UpdateDictData(dictData *systemModels.SysDictDataDML) {
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

	_, err := mysql.MysqlDb.NamedExec(updateSQL, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func DeleteDictDataByIds(dictCodes []int64) {
	query, i, err := sqlx.In("delete from sys_dict_data where dict_code in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = mysql.MysqlDb.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func CountDictDataByTypes(dictType []string) int {
	var count = 0
	query, i, err := sqlx.In("select count(*) from sys_dict_data where dict_type in(?)", dictType)
	if err != nil {
		panic(err)
	}
	err = mysql.MysqlDb.Get(&count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
