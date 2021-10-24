package userStatus

const (
	OK      string = "0"
	Disable string = "1"
	Deleted string = "2"
)

var UserCodeMsgMap = map[string]string{
	OK:      "正常",
	Disable: "停用",
	Deleted: "删除",
}
