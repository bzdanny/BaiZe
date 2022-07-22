package monitorModels

import "github.com/bzdanny/BaiZe/app/system/systemModels"

type SysUserOnline struct {
	TokenId       string  `json:"tokenId"`
	UserName      string  `json:"userName"`
	Ipaddr        string  `json:"ipaddr"`
	LoginLocation string  `json:"loginLocation"`
	Browser       string  `json:"browser"`
	Os            string  `json:"os"`
	LoginTime     int64   `json:"loginTime"`
	DeptName      *string `json:"deptName"`
}

func GetSysUserOnlineByUser(loginUser *systemModels.LoginUser) *SysUserOnline {
	uo := new(SysUserOnline)
	uo.TokenId = loginUser.Token
	uo.UserName = loginUser.User.UserName
	uo.Ipaddr = loginUser.IpAddr
	uo.LoginLocation = loginUser.LoginLocation
	uo.Browser = loginUser.Browser
	uo.Os = loginUser.Os
	uo.LoginTime = loginUser.LoginTime
	uo.DeptName = loginUser.User.DeptName
	return uo
}
