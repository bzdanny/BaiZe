package genTableServiceImpl

import (
	"baize/app/genTable/genTableDao"
	"baize/app/genTable/genTableDao/genTableDaoImpl"
	"baize/app/genTable/genTableModels"
	"baize/app/utils/snowflake"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"time"
)

var genTabletServiceImpl = &genTabletService{genTabletDao: genTableDaoImpl.GetGenTableDao(), genTabletColumnDao: genTableDaoImpl.GetGenTableColumnDao()}

type genTabletService struct {
	genTabletDao       genTableDao.IGenTable
	genTabletColumnDao genTableDao.IGenTableColumn
}

func GetGenTabletService() *genTabletService {
	return genTabletServiceImpl
}

func (genTabletService *genTabletService) SelectGenTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64) {
	return genTabletService.genTabletDao.SelectGenTableList(getTable)
}
func (genTabletService *genTabletService) SelectDbTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64) {
	return genTabletService.genTabletDao.SelectDbTableList(getTable)
}
func (genTabletService *genTabletService) SelectGenTableAll() (list []*genTableModels.GenTableVo) {
	return genTabletService.genTabletDao.SelectGenTableAll()
}
func (genTabletService *genTabletService) SelectGenTableById(id int64) (genTable *genTableModels.GenTableVo) {
	return genTabletService.genTabletDao.SelectGenTableById(id)
}
func (genTabletService *genTabletService) ImportTableSave(table []string, userName string) {
	tableList := genTabletService.genTabletDao.SelectDbTableListByNames(table)
	genTableList := make([]*genTableModels.GenTableDML, 0, len(tableList))
	genTableColumnList := make([]*genTableModels.GenTableColumnDML, 0, len(tableList)*2)
	for _, genTable := range tableList {
		tableId := snowflake.GenID()
		genTableList = append(genTableList, genTableModels.GetGenTableDML(genTable, tableId, userName))
		list := genTabletService.genTabletColumnDao.SelectDbTableColumnsByName(genTable.TableName)
		for _, column := range list {
			genTableColumnList = append(genTableColumnList, genTableModels.GetGenTableColumnDML(column, tableId, userName))
		}
	}
	genTabletService.genTabletDao.BatchInsertGenTable(genTableList)
	genTabletService.genTabletColumnDao.BatchInsertGenTableColumn(genTableColumnList)

}
func (genTabletService *genTabletService) UpdateGenTable(genTable *genTableModels.GenTableDML) (err error) {
	genTabletService.genTabletDao.UpdateGenTable(genTable)
	for _, cenTableColumn := range genTable.Columns {
		genTabletService.genTabletColumnDao.UpdateGenTableColumn(cenTableColumn)
	}
	return
}

func (genTabletService *genTabletService) DeleteGenTableByIds(ids []int64) (err error) {
	genTabletService.genTabletDao.DeleteGenTableByIds(ids)
	genTabletService.genTabletColumnDao.DeleteGenTableColumnByIds(ids)
	return nil
}
func (genTabletService *genTabletService) PreviewCode(tableId int64) (genTable *genTableModels.GenTableVo, err error) {
	genTable = genTabletService.genTabletDao.SelectGenTableById(tableId)
	genTable.Columns = genTabletService.genTabletColumnDao.SelectGenTableColumnListByTableId(tableId)

	genTable.GenerateTime = time.Now()
	jsons, _ := json.Marshal(genTable)
	fmt.Println(string(jsons))
	s := genTabletService.loadTemplate("./template/vm/go/dao.go.vm", genTable)
	fmt.Println(s)
	return genTable, nil
}

func (genTabletService *genTabletService) loadTemplate(templateName string, data interface{}) string {

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
