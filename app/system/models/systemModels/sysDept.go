package systemModels

import "github.com/bzdanny/BaiZe/baize/baizeEntity"

type SysDeptDQL struct {
	ParentId *int64 `form:"parentId,string" db:"parent_id"`
	DeptName string `form:"deptName" db:"dept_name"`
	Status   string `form:"status" db:"status"`
	baizeEntity.BaseEntityDQL
}

type SysDeptVo struct {
	DeptId    int64   `json:"deptId,string" db:"dept_id"`
	ParentId  int64   `json:"parentId,string" db:"parent_id"`
	Ancestors *string `json:"ancestors" db:"ancestors"`
	DeptName  *string `json:"deptName" db:"dept_name"`
	OrderNum  string  `json:"orderNum" db:"order_num"`
	Leader    *string `json:"leader" db:"leader"`
	Phone     string  `json:"phone" db:"phone"`
	Email     string  `json:"email" db:"email"`
	Status    string  `json:"status" db:"status"`
	DelFlag   string  `json:"delFag" db:"del_flag"`
	Remark    *string `json:"remark" db:"remark"`
	baizeEntity.BaseEntity
}
type SysDeptAdd struct {
	DeptId    int64  `json:"DeptId,string" db:"dept_id"`
	ParentId  int64  `json:"parentId,string" db:"parent_id"`
	DeptName  string `json:"deptName" db:"dept_name"`
	Ancestors string `json:"ancestors" db:"ancestors"`
	OrderNum  string `json:"orderNum" db:"order_num"`
	Leader    string `json:"leader" db:"leader"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
	Status    string `json:"status" db:"status"`
	baizeEntity.BaseEntityAdd
}
type SysDeptEdit struct {
	DeptId    int64  `json:"DeptId,string" db:"dept_id"`
	ParentId  int64  `json:"parentId,string" db:"parent_id"`
	DeptName  string `json:"deptName" db:"dept_name"`
	Ancestors string `json:"ancestors" db:"ancestors"`
	OrderNum  string `json:"orderNum" db:"order_num"`
	Leader    string `json:"leader" db:"leader"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
	Status    string `json:"status" db:"status"`
	baizeEntity.BaseEntityEdit
}
