package genUtils

//数据库字符串类型
var COLUMNTYPE_STR = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

//数据库时间类型
var COLUMNTYPE_TIME = []string{"datetime", "time", "date", "timestamp"}

//数据库数字类型
var COLUMNTYPE_NUMBER = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"}

//页面不需要编辑字段
var COLUMNNAME_NOT_EDIT = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

//页面不需要显示的列表字段
var COLUMNNAME_NOT_LIST = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time"}

//页面不需要查询字段
var COLUMNNAME_NOT_QUERY = []string{"id", "create_by", "create_time", "del_flag", "update_by", "update_time", "remark"}

//判断string 是否存在在数组中
func IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

//判断是否是数据库字符串类型
func IsStringObject(value string) bool {
	return IsExistInArray(value, COLUMNTYPE_STR)
}

//判断是否是数据库时间类型
func IsTimeObject(value string) bool {
	return IsExistInArray(value, COLUMNTYPE_TIME)
}

//判断是否是数据库数字类型
func IsNumberObject(value string) bool {
	return IsExistInArray(value, COLUMNTYPE_NUMBER)
}

//页面不需要编辑字段
func IsNotEdit(value string) bool {
	return !IsExistInArray(value, COLUMNNAME_NOT_EDIT)
}

//页面不需要显示的列表字段
func IsNotList(value string) bool {
	return !IsExistInArray(value, COLUMNNAME_NOT_LIST)
}

//页面不需要查询字段
func IsNotQuery(value string) bool {
	return !IsExistInArray(value, COLUMNNAME_NOT_QUERY)
}
