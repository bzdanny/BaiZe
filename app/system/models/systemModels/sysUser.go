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

type SysUserDML struct {
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

func RowsToSysUserDMLList(rows [][]string) (list []*SysUserDML, str string, failureNum int) {
	list = make([]*SysUserDML, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if row[0] == "" || row[1] == "" {
			str += "<br/>第" + strconv.Itoa(i+1) + "行数据格式有误"
			failureNum++
			continue
		}
		sysUser := new(SysUserDML)
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
	row = []string{"登录名称(*)", "用户姓名(*)", "用户邮箱", "手机号码", "用户性别(0男 1女)", "帐号状态(0正常 1停用)", "*号为必填项"}
	return
}

func SysUserListToRows(sysUsers []*SysUserVo) (rows [][]string) {
	rows = make([][]string, 0, len(sysUsers)+1)
	row1 := []string{"登录名称", "用户姓名", "用户邮箱", "手机号码", "用户性别", "帐号状态", "部门名称"}
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
			row[4] = "男"
		} else if sysUser.Sex == "1" {
			row[4] = "女"
		} else {
			row[4] = "未知"
		}
		if sysUser.Status == "0" {
			row[5] = "正常"
		} else {
			row[5] = "停用"
		}

		if sysUser.DeptName != nil {
			row[6] = *sysUser.DeptName
		}
		rows = append(rows, row)

	}
	return
}
