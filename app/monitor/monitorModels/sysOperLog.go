package monitorModels

import (
	"baize/app/common/commonModels"
	"baize/app/utils/excelDictionaries"
	"baize/app/utils/unix"
	"github.com/gogf/gf/util/gconv"
)

type SysOpenLog struct {
	OperId        int64           `json:"operId,string" db:"oper_id"`
	Title         string          `json:"title" db:"title"`
	BusinessType  int8            `json:"businessType" db:"business_type"`
	Method        string          `json:"method" db:"method"`
	RequestMethod string          `json:"requestMethod" db:"request_method"`
	OperatorType  int8            `json:"operatorType" db:"operator_type"`
	OperName      string          `json:"operName" db:"oper_name"`
	DeptName      *string         `json:"deptName" db:"dept_name"`
	OperUrl       string          `json:"operUrl" db:"oper_url"`
	OperIp        string          `json:"operIp" db:"oper_ip"`
	OperLocation  string          `json:"operLocation" db:"oper_location"`
	OperParam     string          `json:"operParam" db:"oper_param"`
	JsonResult    string          `json:"jsonResult" db:"json_result"`
	Status        int8            `json:"status" db:"status"`
	ErrorMsg      *string         `json:"errorMsg" db:"error_msg"`
	OperTime      *unix.BaiZeTime `json:"operTime" db:"oper_time"`
}

type SysOpenLogDQL struct {
	Title        string `form:"title" db:"title"`
	BusinessType *int8  `form:"businessType" db:"business_type"`
	Status       *int8  `form:"status" db:"status"`
	OperName     string `form:"operName" db:"oper_name"`
	BeginTime    string `form:"beginTime" db:"end_time"`
	EndTime      string `form:"endTime" db:"end_time"`
	commonModels.BaseEntityDQL
}

func SysOperLogListToRows(openLogs []*SysOpenLog) (rows [][]string) {
	rows = make([][]string, 0, len(openLogs)+1)
	row1 := []string{"系统模块", "操作类型", "请求方式", "操作人员", "主机", "操作状态", "操作时间"}
	rows = append(rows, row1)
	for _, openLog := range openLogs {
		row := make([]string, 7)
		row[0] = openLog.Title
		row[1] = excelDictionaries.ValueToLabel("sys_oper_type", gconv.String(openLog.BusinessType))
		row[2] = openLog.RequestMethod
		row[3] = openLog.OperName
		row[4] = openLog.OperIp
		row[5] = excelDictionaries.ValueToLabel("sys_common_status", gconv.String(openLog.Status))
		row[6] = openLog.OperTime.ToString()
		rows = append(rows, row)
	}
	return
}
