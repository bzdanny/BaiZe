package genTableServiceImpl

import (
	"bytes"
	"fmt"
	"github.com/bzdanny/BaiZe/app/genTable/genTableDao"
	"github.com/bzdanny/BaiZe/app/genTable/genTableDao/genTableDaoImpl"
	"github.com/bzdanny/BaiZe/app/genTable/genTableModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/utils/snowflake"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"time"
)

type GenTabletService struct {
	data               *datasource.Data
	genTabletDao       genTableDao.IGenTable
	genTabletColumnDao genTableDao.IGenTableColumn
}

func GetGenTabletService(data *datasource.Data, gtc *genTableDaoImpl.GenTableColumnDao, gtd *genTableDaoImpl.GenTableDao) *GenTabletService {
	return &GenTabletService{data: data, genTabletColumnDao: gtc, genTabletDao: gtd}
}

func (g *GenTabletService) SelectGenTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.GenTableVo, total *int64) {
	return g.genTabletDao.SelectGenTableList(g.data.GetSlaveDb(), getTable)
}
func (g *GenTabletService) SelectDbTableList(getTable *genTableModels.GenTableDQL) (list []*genTableModels.DBTableVo, total *int64) {
	return g.genTabletDao.SelectDbTableList(g.data.GetSlaveDb(), getTable)
}
func (g *GenTabletService) SelectGenTableAll() (list []*genTableModels.GenTableVo) {
	return g.genTabletDao.SelectGenTableAll(g.data.GetSlaveDb())
}
func (g *GenTabletService) SelectGenTableById(id int64) (genTable *genTableModels.GenTableVo) {
	return g.genTabletDao.SelectGenTableById(g.data.GetSlaveDb(), id)
}
func (g *GenTabletService) ImportTableSave(table []string, userId int64) {
	tableList := g.genTabletDao.SelectDbTableListByNames(g.data.GetSlaveDb(), table)
	genTableList := make([]*genTableModels.GenTableAdd, 0, len(tableList))
	genTableColumnList := make([]*genTableModels.GenTableColumnDML, 0, len(tableList)*2)
	for _, genTable := range tableList {
		tableId := snowflake.GenID()
		genTableList = append(genTableList, genTableModels.GetGenTableDML(genTable, tableId, userId))
		list := g.genTabletColumnDao.SelectDbTableColumnsByName(g.data.GetSlaveDb(), genTable.TableName)
		for _, column := range list {
			genTableColumnList = append(genTableColumnList, genTableModels.GetGenTableColumnDML(column, tableId, userId))
		}
	}
	tx, err := g.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	g.genTabletDao.BatchInsertGenTable(tx, genTableList)
	g.genTabletColumnDao.BatchInsertGenTableColumn(tx, genTableColumnList)

}
func (g *GenTabletService) UpdateGenTable(genTable *genTableModels.GenTableEdit) (err error) {
	tx, err := g.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	g.genTabletDao.UpdateGenTable(tx, genTable)
	for _, cenTableColumn := range genTable.Columns {
		g.genTabletColumnDao.UpdateGenTableColumn(tx, cenTableColumn)
	}
	return
}

func (g *GenTabletService) DeleteGenTableByIds(ids []int64) (err error) {
	tx, err := g.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	g.genTabletDao.DeleteGenTableByIds(tx, ids)
	g.genTabletColumnDao.DeleteGenTableColumnByIds(tx, ids)
	return nil
}
func (g *GenTabletService) PreviewCode(tableId int64) (genTable *genTableModels.GenTableVo, err error) {
	genTable = g.genTabletDao.SelectGenTableById(g.data.GetSlaveDb(), tableId)
	genTable.Columns = g.genTabletColumnDao.SelectGenTableColumnListByTableId(g.data.GetSlaveDb(), tableId)

	genTable.GenerateTime = time.Now()

	s := g.loadTemplate("./template/vm/go/model/model.go.vm", genTable)
	fmt.Println(s)
	return genTable, nil
}

func (g *GenTabletService) loadTemplate(templateName string, data interface{}) string {

	b, err := os.ReadFile(templateName)
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
