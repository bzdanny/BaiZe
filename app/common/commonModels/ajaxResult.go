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

func SuccessListData(rows interface{}, total *int64) *ResponseData {

	responseData := &ResponseData{Code: httpStatus.Success, Msg: httpStatus.Success.Msg(), Data: ListData{Rows: rows, Total: total}}
	return responseData
}

func SuccessData(data interface{}) *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Success, Msg: httpStatus.Success.Msg(), Data: data}
	return responseData
}
func SuccessMsg(msg string) *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Success, Msg: msg}
	return responseData
}
func Success() *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Success, Msg: httpStatus.Success.Msg()}
	return responseData
}

func ParameterError() *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Parameter, Msg: httpStatus.Parameter.Msg()}
	return responseData
}

func Error() *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Error, Msg: httpStatus.Error.Msg()}
	return responseData
}
func ErrorMsg(msg string) *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Error, Msg: msg}
	return responseData
}

func Waring(msg string) *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Waring, Msg: msg}
	return responseData
}

func InvalidToken() *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Unauthorized, Msg: httpStatus.Unauthorized.Msg()}
	return responseData
}
func PermissionDenied() *ResponseData {
	responseData := &ResponseData{Code: httpStatus.Forbidden, Msg: httpStatus.Forbidden.Msg()}
	return responseData
}
