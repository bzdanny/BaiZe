package genTableDaoImpl

import (
	"fmt"
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type GenTableDao struct {
}

func NewGenTableDao() *GenTableDao {
	return new(GenTableDao)
}

func (genTableDao *GenTableDao) SelectGenTableList(db dataUtil.DB, GenTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64) {
	var selectSql = `select table_id, table_name, table_comment, sub_table_name, sub_table_fk_name, class_name,private_class_name, tpl_category, package_name, module_name, business_name, function_name, function_author, gen_type, gen_path, options, create_by, create_time, update_by, update_time, remark from gen_table`
	whereSql := ``
	if GenTable.TableName != "" {
		whereSql += " AND lower(table_name) like lower(concat('%', :table_name, '%'))"
	}
	if GenTable.TableComment != "" {
		whereSql += " AND lower(table_comment) like lower(concat('%', :table_comment, '%'))"
	}
	if GenTable.BeginTime != "" {
		whereSql += " AND date_format(create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
	}
	if GenTable.EndTime != "" {
		whereSql += " date_format(create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	return dataUtil.NamedQueryListAndTotal(db, list, GenTable, selectSql+whereSql, "", "")
}
func (genTableDao *GenTableDao) SelectDbTableList(db dataUtil.DB, GenTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64) {
	var selectSql = `select table_name , table_comment, create_time, update_time  from information_schema.tables`
	whereSql := ` where table_schema = (select database())
		AND table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'
		AND table_name NOT IN (select table_name from gen_table)`
	if GenTable.TableName != "" {
		whereSql += " AND lower(table_name) like lower(concat('%', :table_name, '%'))"
	}
	if GenTable.TableComment != "" {
		whereSql += " AND lower(table_comment) like lower(concat('%', :table_comment, '%'))"
	}
	if GenTable.BeginTime != "" {
		whereSql += " AND date_format(create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
	}
	if GenTable.EndTime != "" {
		whereSql += " date_format(create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
	}

	return dataUtil.NamedQueryListAndTotal(db, list, GenTable, selectSql+whereSql, "", "")
}

func (genTableDao *GenTableDao) SelectDbTableListByNames(db dataUtil.DB, tableNames []string) (list []*genTableModels.DBTableVo) {
	query, i, err := sqlx.In("select table_name, table_comment, create_time, update_time from information_schema.tables where table_name NOT LIKE 'gen_%' and table_schema = (select database()) and table_name in  (?)", tableNames)
	if err != nil {
		panic(err)
	}
	list = make([]*genTableModels.DBTableVo, 0)
	err = db.Select(&list, query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) SelectGenTableById(db dataUtil.DB, id int64) (genTable *genTableModels.GenTableVo) {
	genTable = new(genTableModels.GenTableVo)
	err := db.Get(genTable, `SELECT
       table_id, table_name, table_comment, sub_table_name,sub_table_fk_name, class_name, private_class_name,
      tpl_category, package_name,module_name, business_name,function_name, function_author,gen_type,gen_path, options, remark
		FROM gen_table 
		where table_id = ?`, id)
	if err != nil {
		panic(err)
	}
	return
}
func (genTableDao *GenTableDao) SelectGenTableByName(db dataUtil.DB, name string) (genTable *genTableModels.GenTableVo) {
	genTable = new(genTableModels.GenTableVo)
	err := db.Get(genTable, `SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, t.private_class_name,t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, t.options, t.remark
		FROM gen_table t
		where t.table_name = ? `, name)
	if err != nil {
		panic(err)
	}
	return
}
func (genTableDao *GenTableDao) SelectGenTableAll(db dataUtil.DB) (list []*genTableModels.GenTableVo) {
	list = make([]*genTableModels.GenTableVo, 0)
	err := db.Select(&list, `SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, t.private_class_name,t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, t.options, t.remark
		FROM gen_table t`)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) BatchInsertGenTable(db dataUtil.DB, genTables []*genTableModels.GenTableAdd) {

	_, err := db.NamedExec(`insert into gen_table(table_id,table_name,table_comment,class_name,private_class_name,tpl_category,package_name,module_name,business_name,function_name,function_author,gen_type,gen_path,create_by,create_time,update_by,update_time,remark)
							values(:table_id,:table_name,:table_comment,:class_name,:private_class_name,:tpl_category,:package_name,:module_name,:business_name,:function_name,:function_author,:gen_type,:gen_path,:create_by,now(),:update_by,now(),:remark)`,
		genTables)
	if err != nil {
		panic(err)
	}

}

func (genTableDao *GenTableDao) InsertGenTable(db dataUtil.DB, genTable *genTableModels.GenTableAdd) {
	insertSQL := `insert into gen_table(table_id,table_name,table_comment,class_name,private_class_name,tpl_category,package_name,module_name,business_name,function_name,function_author,gen_type,gen_path,create_by,create_time,update_by,update_time %s)
					values(:table_id,:table_name,:table_comment,:class_name,:private_class_name,:tpl_category,:package_name,:module_name,:business_name,:function_name,:function_author,:gen_type,:gen_path,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""

	if genTable.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, genTable)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) UpdateGenTable(db dataUtil.DB, genTable *genTableModels.GenTableEdit) {
	updateSQL := `update gen_table set update_time = now() , update_by = :update_by`
	if genTable.TableName != "" {
		updateSQL += ",table_name = :table_name"
	}
	if genTable.TableComment != "" {
		updateSQL += ",table_comment = :table_comment"
	}
	if genTable.SubTableName != "" {
		updateSQL += ",sub_table_name = :sub_table_name"
	}
	if genTable.SubTableFkName != "" {
		updateSQL += ",sub_table_fk_name = :sub_table_fk_name"
	}
	if genTable.ClassName != "" {
		updateSQL += ",class_name = :class_name"
	}
	if genTable.PrivateClassName != "" {
		updateSQL += ",private_class_name = :private_class_name"
	}
	if genTable.FunctionAuthor != "" {
		updateSQL += ",function_author = :function_author"
	}
	if genTable.GenType != "" {
		updateSQL += ",gen_type = :gen_type"
	}
	if genTable.GenPath != "" {
		updateSQL += ",gen_path = :gen_path"
	}
	if genTable.TplCategory != "" {
		updateSQL += ",tpl_category = :tpl_category"
	}
	if genTable.PackageName != "" {
		updateSQL += ",package_name = :package_name"
	}
	if genTable.ModuleName != "" {
		updateSQL += ",module_name = :module_name"
	}
	if genTable.BusinessName != "" {
		updateSQL += ",business_name = :business_name"
	}
	if genTable.FunctionName != "" {
		updateSQL += ",function_name = :function_name"
	}
	if genTable.Options != "" {
		updateSQL += ",options = :options"
	}
	if genTable.Remark != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where table_id = :table_id"

	_, err := db.NamedExec(updateSQL, genTable)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) DeleteGenTableByIds(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In(" delete from gen_table where table_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}

}
