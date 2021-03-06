package systemModels

import (
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
	"strconv"
	"time"
)

type SysUserDQL struct {
	UserName    string `form:"userName" db:"user_name"`
	Status      string `form:"status" db:"status"`
	Phonenumber string `form:"phonenumber" db:"phonenumber"`
	BeginTime   string `form:"beginTime" db:"begin_time"`
	EndTime     string `form:"endTime" db:"end_time"`
	DeptId      *int64 `form:"deptId" db:"dept_id"`
	baizeEntity.BaseEntityDQL
}

type SysUserAdd struct {
	UserId      int64    `json:"userId,string" db:"user_id"`
	DeptId      *int64   `json:"deptId,string" db:"dept_id"`
	UserName    string   `json:"userName" db:"user_name"`
	NickName    string   `json:"nickName" db:"nick_name"`
	Email       string   `json:"email" db:"email"`
	Avatar      string   `json:"avatar" db:"avatar"`
	Phonenumber string   `json:"phonenumber" db:"phonenumber"`
	Sex         string   `json:"sex" db:"sex"  binding:"required"`
	Password    string   `json:"password" db:"password"`
	Status      string   `json:"status" db:"status"`
	Remake      string   `json:"remake" db:"remake "`
	PostIds     []string `json:"postIds"`
	RoleIds     []string `json:"roleIds"`
	baizeEntity.BaseEntityAdd
}
type SysUserEdit struct {
	UserId      int64    `json:"userId,string" db:"user_id"`
	DeptId      *int64   `json:"deptId,string" db:"dept_id"`
	UserName    string   `json:"userName" db:"user_name"`
	NickName    string   `json:"nickName" db:"nick_name"`
	Email       string   `json:"email" db:"email"`
	Avatar      string   `json:"avatar" db:"avatar"`
	Phonenumber string   `json:"phonenumber" db:"phonenumber"`
	Sex         string   `json:"sex" db:"sex"  binding:"required"`
	Password    string   `json:"password" db:"password"`
	Status      string   `json:"status" db:"status"`
	Remake      string   `json:"remake" db:"remake "`
	PostIds     []string `json:"postIds"`
	RoleIds     []string `json:"roleIds"`
	baizeEntity.BaseEntityEdit
}

func RowsToSysUserDMLList(rows [][]string) (list []*SysUserAdd, str string, failureNum int) {
	list = make([]*SysUserAdd, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if row[0] == "" || row[1] == "" {
			str += "<br/>???" + strconv.Itoa(i+1) + "?????????????????????"
			failureNum++
			continue
		}
		sysUser := new(SysUserAdd)
		sysUser.UserName = row[0]
		sysUser.NickName = row[1]
		sysUser.Email = row[2]
		sysUser.Phonenumber = row[3]
		sysUser.Sex = row[4]
		sysUser.Status = row[5]
		list = append(list, sysUser)
	}
	return
}

type SysUserVo struct {
	UserId      int64      `json:"userId,string" db:"user_id"`
	UserName    string     `json:"userName" db:"user_name"`
	NickName    string     `json:"nickName" db:"nick_name"`
	Sex         string     `json:"sex" db:"sex"`
	Status      string     `json:"status" db:"status"`
	DelFlag     string     `json:"delFlag" db:"del_flag"`
	LoginIp     string     `json:"loginIp" db:"login_ip"`
	DeptId      *int64     `json:"deptId" db:"dept_id"`
	LoginDate   *time.Time `json:"loginDate" db:"login_date"`
	DeptName    *string    `json:"deptName" db:"dept_name"`
	Leader      *string    `json:"leader" db:"leader"`
	Email       *string    `json:"email" db:"email"`
	Phonenumber *string    `json:"phonenumber"db:"phonenumber"`
	Avatar      *string    `json:"avatar" db:"avatar"`
	RoleId      *int64     `json:"roleId" db:"role_id"`
	Remark      *string    `json:"remark" db:"remark"`
	Roles       []*SysRole
	baizeEntity.BaseEntity
}

//sysUser.UserName = row[0]
//sysUser.NickName = row[1]
//sysUser.Email = row[2]
//sysUser.Phonenumber = row[3]
//sysUser.Sex = row[4]
//sysUser.Status = row[5]

func SysUserImportTemplate() (row []string) {
	row = []string{"????????????(*)", "????????????(*)", "????????????", "????????????", "????????????(0??? 1???)", "????????????(0?????? 1??????)", "*???????????????"}
	return
}

func SysUserListToRows(sysUsers []*SysUserVo) (rows [][]string) {
	rows = make([][]string, 0, len(sysUsers)+1)
	row1 := []string{"????????????", "????????????", "????????????", "????????????", "????????????", "????????????", "????????????"}
	rows = append(rows, row1)
	for _, sysUser := range sysUsers {
		row := make([]string, 7)
		row[0] = sysUser.UserName
		row[1] = sysUser.NickName
		if sysUser.Email != nil {
			row[2] = *sysUser.Email
		}
		if sysUser.Phonenumber != nil {
			row[3] = *sysUser.Phonenumber
		}
		if sysUser.Sex == "0" {
			row[4] = "???"
		} else if sysUser.Sex == "1" {
			row[4] = "???"
		} else {
			row[4] = "??????"
		}
		if sysUser.Status == "0" {
			row[5] = "??????"
		} else {
			row[5] = "??????"
		}

		if sysUser.DeptName != nil {
			row[6] = *sysUser.DeptName
		}
		rows = append(rows, row)

	}
	return
}
