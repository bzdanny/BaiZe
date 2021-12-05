package business

type BusinessStatus string

const (
	Success BusinessStatus = "SUCCESS"
	Fail    BusinessStatus = "FAIL"
)

var businessStatusMap = map[BusinessStatus]int8{
	Success: 0,
	Fail:    1,
}

func (c BusinessStatus) Msg() int8 {
	msg, _ := businessStatusMap[c]
	return msg
}
