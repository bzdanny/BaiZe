package monitorModels

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/constant/business"
	"baize/app/constant/constants"
	"baize/app/utils/unix"
	"github.com/gin-gonic/gin"
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

func SetLog(c *gin.Context, title string, businessTy business.BusinessType) {
	ol := new(SysOpenLog)
	ol.Title = title
	ol.BusinessType = businessTy.Msg()
	ol.Status = business.Success.Msg()
	ol.OperIp = c.ClientIP()
	ol.OperUrl = c.Request.URL.Path
	ol.RequestMethod = c.Request.Method
	loginUser := commonController.GetCurrentLoginUser(c)
	if loginUser != nil {
		ol.OperName = loginUser.User.UserName
	}
	c.Set(constants.LogKey, ol)
}
