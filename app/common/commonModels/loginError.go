package commonModels

type ErrorData struct {
	Msg string
}

func (e ErrorData) Error() string {
	return e.Msg
}

func NewErrorData(msg string) *ErrorData {
	return &ErrorData{Msg: msg}
}
