package genTableModels

import (
	"baize/app/common/baize/baizeUnix"
	"baize/app/common/commonModels"
	genUtils "baize/app/genTable/utils"
	"baize/app/utils/stringUtils"
	"fmt"
	"strings"
	"time"
)

type GenTableDQL struct {
	TableName    string `form:"tableName" db:"table_name"`
	TableComment string `form:"tableComment" db:"table_comment"`
	BeginTime    string `form:"beginTime" db:"begin_time"`
	EndTime      string `form:"endTime" db:"end_time"`
	commonModels.BaseEntityDQL
}

type GenTableDML struct {
	TableId        int64                `json:"tableId,string" db:"table_id"`
	ParentMenuId   int64                `json:"parentMenuId,string" db:"parent_menu_id"`
	TableName      string               `json:"tableName" db:"table_name"`
	TableComment   string               `json:"tableComment" db:"table_comment"`
	SubTableName   string               `json:"subTableName" db:"sub_table_name"`
	SubTableFkName string               `json:"subTableFkName" db:"sub_table_fk_name"`
	ClassName      string               `json:"className" db:"class_name"`
	TplCategory    string               `json:"tplCategory" db:"tpl_category"`
	PackageName    string               `json:"packageName" db:"package_name"`
	ModuleName     string               `json:"moduleName" db:"module_name"`
	BusinessName   string               `json:"businessName" db:"business_name"`
	FunctionName   string               `json:"functionName" db:"function_name"`
	FunctionAuthor string               `json:"functionAuthor" db:"function_author"`
	GenType        string               `json:"genType" db:"gen_type"`
	GenPath        string               `json:"genPath" db:"gen_path"`
	Remark         string               `json:"remark" db:"remark"`
	Columns        []*GenTableColumnDML `json:"columns"`
	commonModels.BaseEntityDML
}

func GetGenTableDML(table *DBTableVo, tableId int64, operName string) *GenTableDML {
	gen := new(GenTableDML)
	gen.TableId = tableId
	gen.TableName = table.TableName
	gen.TableComment = table.TableComment
	gen.ClassName = stringUtils.ConvertToBigCamelCase(genUtils.ConvertClassName(table.TableName))
	gen.PackageName = "baize"
	gen.ModuleName = "module"
	gen.BusinessName = genUtils.GetBusinessName(table.TableName)
	gen.FunctionName = strings.ReplaceAll(table.TableComment, "表", "")
	gen.FunctionAuthor = "baize"
	gen.CreateBy = operName
	gen.TplCategory = "crud"
	gen.ParentMenuId = 3
	return gen
}

type GenTableVo struct {
	TableId        int64               `json:"tableId,string" db:"table_id"`
	ParentMenuId   int64               `json:"parentMenuId,string" db:"parent_menu_id"`
	TableName      string              `json:"tableName" db:"table_name"`
	TableComment   string              `json:"tableComment" db:"table_comment"`
	SubTableName   *string             `json:"subTableName" db:"sub_table_name"`
	SubTableFkName *string             `json:"subTableFkName" db:"sub_table_fk_name"`
	ClassName      string              `json:"className" db:"class_name"`
	TplCategory    string              `json:"tplCategory" db:"tpl_category"`
	PackageName    string              `json:"packageName" db:"package_name"`
	ModuleName     string              `json:"moduleName" db:"module_name"`
	BusinessName   string              `json:"businessName" db:"business_name"`
	FunctionName   string              `json:"functionName" db:"function_name"`
	FunctionAuthor string              `json:"functionAuthor" db:"function_author"`
	GenType        string              `json:"genType" db:"gen_type"`
	GenPath        string              `json:"genPath" db:"gen_path"`
	Remark         *string             `json:"remark" db:"remark"`
	Columns        []*GenTableColumnVo `json:"column"`
	GenerateTime   time.Time
	commonModels.BaseEntity
}

type DBTableVo struct {
	TableName    string               `json:"tableName" db:"TABLE_NAME"`
	TableComment string               `json:"tableComment" db:"TABLE_COMMENT"`
	CreateTime   *baizeUnix.BaiZeTime `json:"createTime" db:"CREATE_TIME"`
	UpdateTime   *baizeUnix.BaiZeTime `json:"updateTime" db:"UPDATE_TIME"`
}

func (genTable *GenTableVo) GetFileName(template string) string {
	if strings.HasSuffix(template, "model.go.vm") {
		return fmt.Sprintf("/go/%s/%sModels/%s.go", genTable.ModuleName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "controller.go.vm") {
		return fmt.Sprintf("/go/%s/%sController/%s.go", genTable.ModuleName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "iService.go.vm") {
		return fmt.Sprintf("/go/%s/%sService/i%s.go", genTable.ModuleName, genTable.PackageName, stringUtils.Capitalize(genTable.BusinessName))
	}
	if strings.HasSuffix(template, "serviceImpl.go.vm") {
		return fmt.Sprintf("/go/%s/%sService/%sServiceImpl/%sImpl.go", genTable.ModuleName, genTable.PackageName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "iDao.go.vm") {
		return fmt.Sprintf("/go/%s/%sDao/i%s.go", genTable.ModuleName, genTable.PackageName, stringUtils.Capitalize(genTable.BusinessName))
	}
	if strings.HasSuffix(template, "daoImpl.go.vm") {
		return fmt.Sprintf("/go/%s/%sDao/%sDaoImpl/%sImpl.go", genTable.ModuleName, genTable.PackageName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "router.go.vm") {
		return fmt.Sprintf("/go/routes/%sRoutes/%sRouter.go", genTable.ModuleName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "sql.sql.vm") {
		return fmt.Sprintf("/go/%s.sql", genTable.BusinessName)
	}
	if strings.HasSuffix(template, "api.js.vm") {
		return fmt.Sprintf("/vue/api/%s/%s.js", genTable.ModuleName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "index.vue.vm") {
		return fmt.Sprintf("/vue/views/%s/%s/index.vue", genTable.ModuleName, genTable.BusinessName)
	}
	return ""
}
