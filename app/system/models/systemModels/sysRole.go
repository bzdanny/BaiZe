package systemModels

import "github.com/bzdanny/BaiZe/baize/baizeEntity"

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
	baizeEntity.BaseEntityDQL
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
	baizeEntity.BaseEntity
}

func SysRoleListToRows(roles []*SysRoleVo) (rows [][]string) {
	rows = make([][]string, 0, len(roles)+1)
	row1 := []string{"角色名称", "权限字符", "显示状态"}
	rows = append(rows, row1)
	for _, sysRole := range roles {
		row := make([]string, 3)
		row[0] = sysRole.RoleName
		row[1] = sysRole.RoleKey
		if sysRole.Status == "0" {
			row[3] = "正常"
		} else {
			row[3] = "停用"
		}
		rows = append(rows, row)

	}
	return
}

type SysRoleAdd struct {
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
	baizeEntity.BaseEntityAdd
}
type SysRoleEdit struct {
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
	baizeEntity.BaseEntityEdit
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
	RoleId      string `form:"roleId" db:"role_id"`
	UserName    string `form:"userName" db:"user_name"`
	Phonenumber string `form:"phonenumber" db:"phonenumber"`

	baizeEntity.BaseEntityDQL
}
