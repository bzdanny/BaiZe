package genTableService

import (
	"baize/app/genTable/genTableDao"
	"baize/app/genTable/genTableModels"
	"baize/app/utils/snowflake"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"text/template"
	"time"
)

func SelectGenTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64) {
	return genTableDao.SelectGenTableList(getTable)
}
func SelectDbTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64) {
	return genTableDao.SelectDbTableList(getTable)
}
func SelectGenTableAll() (list []*genTableModels.GenTableVo) {
	return genTableDao.SelectGenTableAll()
}
func SelectGenTableById(id int64) (genTable *genTableModels.GenTableVo) {
	return genTableDao.SelectGenTableById(id)
}
func ImportTableSave(table []string, userName string) {
	tableList := genTableDao.SelectDbTableListByNames(table)
	genTableList := make([]*genTableModels.GenTableDML, 0, len(tableList))
	genTableColumnList := make([]*genTableModels.GenTableColumnDML, 0, len(tableList)*2)
	for _, genTable := range tableList {
		tableId := snowflake.GenID()
		genTableList = append(genTableList, genTableModels.GetGenTableDML(genTable, tableId, userName))
		list := genTableDao.SelectDbTableColumnsByName(genTable.TableName)
		for _, column := range list {
			genTableColumnList = append(genTableColumnList, genTableModels.GetGenTableColumnDML(column, tableId, userName))
		}
	}
	genTableDao.BatchInsertGenTable(genTableList)
	genTableDao.BatchInsertGenTableColumn(genTableColumnList)

}
func UpdateGenTable(genTable *genTableModels.GenTableDML) (err error) {
	genTableDao.UpdateGenTable(genTable)
	for _, cenTableColumn := range genTable.Columns {
		genTableDao.UpdateGenTableColumn(cenTableColumn)
	}
	return
}

func DeleteGenTableByIds(ids []int64) (err error) {
	genTableDao.DeleteGenTableByIds(ids)
	genTableDao.DeleteGenTableColumnByIds(ids)
	return nil
}
func PreviewCode(tableId int64) (genTable *genTableModels.GenTableVo, err error) {
	genTable = genTableDao.SelectGenTableById(tableId)
	genTable.Columns = genTableDao.SelectGenTableColumnListByTableId(tableId)

	genTable.GenerateTime=time.Now()
	jsons, _ := json.Marshal(genTable)
	fmt.Println(string(jsons))
	s := loadTemplate("./template/vm/go/dao.go.vm", genTable)
	fmt.Println(s)
	return genTable, nil
}

func loadTemplate(templateName string, data interface{}) string {

	b, err := ioutil.ReadFile(templateName)
	if err != nil {
		panic(err)
	}
	templateStr := string(b)

	tmpl, err := template.New(templateName).Parse(templateStr) //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, gin.H{"table": data}) //将string与模板合成，变量name的内容会替换掉{{.}}
	if err != nil {
		print(err)
	}
	return buffer.String()
}
