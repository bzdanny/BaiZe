package systemModels

import (
	commonModels "baize/app/common/commonModels"
	"fmt"
)

type SysRole struct {
	RoleId    int64  `db:"role_id"`
	RoleName  string `db:"role_name"`
	RoleKey   string `db:"role_key"`
	DataScope string `db:"data_scope"`
}

type SysRoleDQL struct {
	RoleName  string `form:"roleName" db:"role_name"`
	Status    string `form:"status" db:"status"`
	RoleKey   string `form:"roleKey" db:"role_key"`
	BeginTime string `form:"beginTime" db:"begin_time"`
	EndTime   string `form:"endTime" db:"end_time"`
	commonModels.BaseEntityDQL
}

type SysRoleVo struct {
	RoleId            int64   `json:"roleId,string" db:"role_id"`
	RoleName          string  `json:"roleName" db:"role_name"`
	RoleKey           string  `json:"roleKey" db:"role_key"`
	RoleSort          int     `json:"roleSort" db:"role_sort"`
	DataScope         string  `json:"dataScope" db:"data_scope"`
	MenuCheckStrictly bool    `json:"menuCheckStrictly" db:"menu_check_strictly"`
	DeptCheckStrictly bool    `json:"deptCheckStrictly" db:"dept_check_strictly"`
	Status            string  `json:"status"  db:"status"`
	DelFlag           string  `json:"delFlag" db:"del_flag"`
	Remark            *string `json:"remark" db:"remark"`
	commonModels.BaseEntity
}

func SysRoleDMLListToRows(roles []*SysRoleVo) (rows [][]string) {
	rows = make([][]string, 0, len(roles)+1)
	row1 := []string{"登录名称", "用户姓名", "用户邮箱", "手机号码", "用户性别", "帐号状态", "部门名称"}
	rows = append(rows, row1)
	for _, sysRole := range roles {
		row := make([]string, 7, 7)
		fmt.Println(sysRole)
		rows = append(rows, row)

	}
	return
}

type SysRoleDML struct {
	RoleId            int64    `json:"RoleId,string" db:"role_id"`
	RoleName          string   `json:"roleName" db:"role_name"`
	RoleKey           string   `json:"roleKey" db:"role_key"`
	RoleSort          int      `json:"roleSort" db:"role_sort"`
	DataScope         string   `json:"dataScope" db:"data_scope"`
	MenuCheckStrictly *bool    `json:"menuCheckStrictly" db:"menu_check_strictly"`
	DeptCheckStrictly *bool    `json:"deptCheckStrictly" db:"dept_check_strictly" `
	Status            string   `json:"status" db:"status"`
	Remake            string   `json:"remake" db:"remake "`
	MenuIds           []string `json:"menuIds"`
	DeptIds           []string `json:"deptIds"`
	commonModels.BaseEntityDML
}

type SysRoleMenu struct {
	RoleId int64 `db:"role_id"`
	MenuId int64 `db:"menu_id"`
}

type SysRoleDept struct {
	RoleId int64 `db:"role_id"`
	DeptId int64 `db:"dept_id"`
}

type SysRoleAndUserDQL struct {
	RoleId  string `form:"roleId" db:"role_id"`
	UserName    string `form:"userName" db:"user_name"`
	Phonenumber   string `form:"phonenumber" db:"phonenumber"`

	commonModels.BaseEntityDQL
}
