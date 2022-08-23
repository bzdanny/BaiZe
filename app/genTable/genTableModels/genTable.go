package genTableModels

import (
	genUtils "github.com/bzdanny/BaiZe/app/genTable/utils"
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
	"github.com/bzdanny/BaiZe/baize/utils/stringUtils"
	"strings"
	"time"
)

type GenTableDQL struct {
	TableName    string `form:"tableName" db:"table_name"`
	TableComment string `form:"tableComment" db:"table_comment"`
	BeginTime    string `form:"beginTime" db:"begin_time"`
	EndTime      string `form:"endTime" db:"end_time"`
	baizeEntity.BaseEntityDQL
}

type GenTableAdd struct {
	TableId          int64                `json:"tableId,string" db:"table_id"`
	TableName        string               `json:"tableName" db:"table_name"`
	TableComment     string               `json:"tableComment" db:"table_comment"`
	SubTableName     string               `json:"subTableName" db:"sub_table_name"`
	SubTableFkName   string               `json:"subTableFkName" db:"sub_table_fk_name"`
	ClassName        string               `json:"className" db:"class_name"`
	PrivateClassName string               `json:"privateClassName" db:"private_class_name"`
	TplCategory      string               `json:"tplCategory" db:"tpl_category"`
	PackageName      string               `json:"packageName" db:"package_name"`
	ModuleName       string               `json:"moduleName" db:"module_name"`
	BusinessName     string               `json:"businessName" db:"business_name"`
	FunctionName     string               `json:"functionName" db:"function_name"`
	FunctionAuthor   string               `json:"functionAuthor" db:"function_author"`
	GenType          string               `json:"genType" db:"gen_type"`
	GenPath          string               `json:"genPath" db:"gen_path"`
	Options          string               `json:"options" db:"options"`
	Remark           string               `json:"remark" db:"remark"`
	Columns          []*GenTableColumnDML `json:"columns"`
	baizeEntity.BaseEntityAdd
}
type GenTableEdit struct {
	TableId          int64                `json:"tableId,string" db:"table_id"`
	TableName        string               `json:"tableName" db:"table_name"`
	TableComment     string               `json:"tableComment" db:"table_comment"`
	SubTableName     string               `json:"subTableName" db:"sub_table_name"`
	SubTableFkName   string               `json:"subTableFkName" db:"sub_table_fk_name"`
	ClassName        string               `json:"className" db:"class_name"`
	PrivateClassName string               `json:"privateClassName" db:"private_class_name"`
	TplCategory      string               `json:"tplCategory" db:"tpl_category"`
	PackageName      string               `json:"packageName" db:"package_name"`
	ModuleName       string               `json:"moduleName" db:"module_name"`
	BusinessName     string               `json:"businessName" db:"business_name"`
	FunctionName     string               `json:"functionName" db:"function_name"`
	FunctionAuthor   string               `json:"functionAuthor" db:"function_author"`
	GenType          string               `json:"genType" db:"gen_type"`
	GenPath          string               `json:"genPath" db:"gen_path"`
	Options          string               `json:"options" db:"options"`
	Remark           string               `json:"remark" db:"remark"`
	Columns          []*GenTableColumnDML `json:"columns"`
	baizeEntity.BaseEntityEdit
}

func GetGenTableDML(table *DBTableVo, tableId int64, userId int64) *GenTableAdd {
	gen := new(GenTableAdd)
	gen.TableId = tableId
	gen.TableName = table.TableName
	gen.TableComment = table.TableComment
	gen.ClassName = stringUtils.ConvertToBigCamelCase(genUtils.ConvertClassName(table.TableName))
	gen.PrivateClassName = stringUtils.ConvertToLittleCamelCase(genUtils.ConvertClassName(table.TableName))
	gen.PackageName = "baize"
	gen.ModuleName = "module"
	gen.BusinessName = genUtils.GetBusinessName(table.TableName)
	gen.FunctionName = strings.ReplaceAll(table.TableComment, "表", "")
	gen.FunctionAuthor = "baize"
	gen.CreateBy = userId
	gen.TplCategory = "crud"
	return gen
}

type GenTableVo struct {
	TableId          int64               `json:"tableId,string" db:"table_id"`
	TableName        string              `json:"tableName" db:"table_name"`
	TableComment     string              `json:"tableComment" db:"table_comment"`
	SubTableName     *string             `json:"subTableName" db:"sub_table_name"`
	SubTableFkName   *string             `json:"subTableFkName" db:"sub_table_fk_name"`
	ClassName        string              `json:"className" db:"class_name"`
	PrivateClassName string              `json:"privateClassName" db:"private_class_name"`
	TplCategory      string              `json:"tplCategory" db:"tpl_category"`
	PackageName      string              `json:"packageName" db:"package_name"`
	ModuleName       string              `json:"moduleName" db:"module_name"`
	BusinessName     string              `json:"businessName" db:"business_name"`
	FunctionName     string              `json:"functionName" db:"function_name"`
	FunctionAuthor   string              `json:"functionAuthor" db:"function_author"`
	GenType          string              `json:"genType" db:"gen_type"`
	GenPath          string              `json:"genPath" db:"gen_path"`
	Options          *string             `json:"options" db:"options"`
	Remark           *string             `json:"remark" db:"remark"`
	Columns          []*GenTableColumnVo `json:"column"`
	GenerateTime     time.Time
	baizeEntity.BaseEntity
}

type DBTableVo struct {
	TableName    string                 `json:"tableName" db:"TABLE_NAME"`
	TableComment string                 `json:"tableComment" db:"TABLE_COMMENT"`
	CreateTime   *baizeEntity.BaiZeTime `json:"createTime" db:"CREATE_TIME"`
	UpdateTime   *baizeEntity.BaiZeTime `json:"updateTime" db:"UPDATE_TIME"`
}
