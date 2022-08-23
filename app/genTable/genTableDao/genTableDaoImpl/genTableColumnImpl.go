package genTableDaoImpl

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type GenTableColumnDao struct {
}

func NewGenTableColumnDao() *GenTableColumnDao {
	return &GenTableColumnDao{}
}

func (genTableColumnDao *GenTableColumnDao) SelectDbTableColumnsByName(db dataUtil.DB, tableName string) (list []*genTableModels.InformationSchemaColumn) {
	list = make([]*genTableModels.InformationSchemaColumn, 0, 0)
	db.Select(&list, `
		select column_name, (case when (is_nullable = 'no'  &&  column_key != 'PRI') then '1' else '0' end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment,  column_type
		from information_schema.columns where table_schema = (select database()) and table_name = (?)
		order by ordinal_position
`, tableName)

	return
}

func (genTableColumnDao *GenTableColumnDao) SelectGenTableColumnListByTableId(db dataUtil.DB, tableId int64) (list []*genTableModels.GenTableColumnVo) {
	list = make([]*genTableModels.GenTableColumnVo, 0, 0)
	db.Select(&list, `select column_id, table_id, column_name, column_comment, column_type, go_type, go_field,html_field, is_pk,  is_required, is_insert, is_edit, is_list, is_query, is_entity,query_type, html_type, dict_type, sort, create_by, create_time, update_by, update_time 
	from gen_table_column
       where table_id = ?
        order by sort`, tableId)

	return
}

func (genTableColumnDao *GenTableColumnDao) BatchInsertGenTableColumn(db dataUtil.DB, genTables []*genTableModels.GenTableColumnDML) {

	_, err := db.NamedExec(`insert into gen_table_column(column_id,table_id,column_name,column_comment,column_type,go_type,go_field,html_field,is_pk,is_required,is_insert,is_edit,is_list, is_query,is_entity, query_type, html_type, dict_type, sort,create_by,create_time,update_by,update_time)
							values(:column_id,:table_id,:column_name,:column_comment,:column_type,:go_type,:go_field,:html_field,:is_pk,:is_required,:is_insert,:is_edit,:is_list, :is_query, :is_entity, :query_type, :html_type, :dict_type, :sort,:create_by,now(),:update_by,now())`,
		genTables)
	if err != nil {
		panic(err)
	}

}

func (genTableColumnDao *GenTableColumnDao) UpdateGenTableColumn(db dataUtil.DB, column *genTableModels.GenTableColumnDML) {
	_, err := db.NamedExec("update gen_table_column set column_comment=:column_comment,go_type=:go_type,go_field=:go_field,html_field=:html_field,is_insert=:is_insert, is_edit=:is_edit,is_list=:is_list,is_query=:is_query,is_required=:is_required,query_type=:query_type,html_type=:html_type,dict_type=:dict_type,sort=:sort, update_by = :update_by,update_time = now()  where column_id = :column_id", column)
	if err != nil {
		panic(err)
	}
}

func (genTableColumnDao *GenTableColumnDao) DeleteGenTableColumnByIds(db dataUtil.DB, ids []int64) {

	query, i, err := sqlx.In("  delete from gen_table_column where table_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
