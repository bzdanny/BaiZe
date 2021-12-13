package monitorModels

import (
	"baize/app/common/commonModels"
	"baize/app/utils/excelDictionaries"
	"baize/app/utils/unix"
	"github.com/gogf/gf/util/gconv"
)

type LogininforDQL struct {
	Status   string `form:"status" db:"status"`
	UserName string `form:"userName" db:"user_name"`
	IpAddr   string `form:"ipAddr" db:"ipaddr"`
	commonModels.BaseEntityDQL
}

type Logininfor struct {
	InfoId        int64          `json:"infoId,string" db:"info_id"`
	Status        int8           `json:"status" db:"status"`
	UserName      string         `json:"userName" db:"user_name"`
	Msg           string         `json:"msg" db:"msg"`
	IpAddr        string         `json:"ipAddr" db:"ipaddr"`
	LoginLocation string         `json:"loginLocation" db:"login_location"`
	Browser       string         `json:"browser" db:"browser"`
	Os            string         `json:"os" db:"os"`
	LoginTime     unix.BaiZeTime `json:"loginTime" db:"login_time"`
}

func SysLogininforToRows(logininfors []*Logininfor) (rows [][]string) {
	rows = make([][]string, 0, len(logininfors)+1)
	row1 := []string{"登录名称", "地址", "登录地点", "操作系统","浏览器","登录状态","描述","访问时间"}
	rows = append(rows, row1)
	for _, logininfor := range logininfors {
		row := make([]string, 8)
		row[0] = logininfor.UserName
		row[1] = logininfor.IpAddr
		row[2] = logininfor.LoginLocation
		row[3] = logininfor.Os
		row[4] = logininfor.Browser
		row[5] = excelDictionaries.ValueToLabel("sys_common_status", gconv.String(logininfor.Status))
		row[4] = logininfor.Msg
		row[6] = logininfor.LoginTime.ToString()
		rows = append(rows, row)

	}
	return
}
