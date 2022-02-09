package commonModels

import (
	"baize/app/constant/httpStatus"
)

type ListData struct {
	Rows  interface{} `json:"rows"`
	Total *int64      `json:"total"`
}

type ResponseData struct {
	Code httpStatus.ResCode `json:"code"`           //相应状态码
	Msg  string             `json:"msg"`            //提示信息
	Data interface{}        `json:"data,omitempty"` //数据
}
