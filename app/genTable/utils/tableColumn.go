package genUtils

// ColumnTypeStr 数据库字符串类型
var ColumnTypeStr = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

// ColumnTypeTime 数据库时间类型
var ColumnTypeTime = []string{"datetime", "time", "date", "timestamp"}

// ColumnTypeNumber 数据库数字类型
var ColumnTypeNumber = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"}

// ColumnNameNotEdit 页面不需要编辑字段
var ColumnNameNotEdit = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

// ColumnNameNotList 页面不需要显示的列表字段
var ColumnNameNotList = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

// ColumnNameNotQuery 页面不需要查询字段
var ColumnNameNotQuery = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time", "remark"}

// ColumnNameBaseEntity 在BaseEntity字段
var ColumnNameBaseEntity = []string{"create_by", "create_time", "del_flag", "update_by", "update_time"}

// IsExistInArray 判断string 是否存在在数组中
func IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// IsStringObject 判断是否是数据库字符串类型
func IsStringObject(value string) bool {
	return IsExistInArray(value, ColumnTypeStr)
}

// IsTimeObject 判断是否是数据库时间类型
func IsTimeObject(value string) bool {
	return IsExistInArray(value, ColumnTypeTime)
}

// IsNumberObject 判断是否是数据库数字类型
func IsNumberObject(value string) bool {
	return IsExistInArray(value, ColumnTypeNumber)
}

// IsNotEdit 页面不需要编辑字段
func IsNotEdit(value string) bool {
	return !IsExistInArray(value, ColumnNameNotEdit)
}

// IsNotList 页面不需要显示的列表字段
func IsNotList(value string) bool {
	return !IsExistInArray(value, ColumnNameNotList)
}

// IsNotQuery 页面不需要查询字段
func IsNotQuery(value string) bool {
	return !IsExistInArray(value, ColumnNameNotQuery)
}

// IsBaseEntity 在BaseEntity里面
func IsBaseEntity(value string) bool {
	return IsExistInArray(value, ColumnNameNotQuery)
}
