package httpStatus

type ResCode int64

const (
	Success      ResCode = 200
	Unauthorized ResCode = 401
	Forbidden    ResCode = 403
	Parameter    ResCode = 412
	Error        ResCode = 500
	Waring       ResCode = 600
)

var codeMsgMap = map[ResCode]string{
	Success:      "success",
	Unauthorized: "无效的令牌",
	Forbidden:    "没有权限，请联系管理员授权",
	Parameter:    "参数错误",
	Error:        "系统异常",
	Waring:       "警告",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[Error]
	}
	return msg
}
