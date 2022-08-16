package systemModels

import "github.com/bzdanny/BaiZe/baize/baizeEntity"

type SysDeptDQL struct {
	ParentId *int64 `form:"parentId,string" db:"parent_id"` //上级id
	DeptName string `form:"deptName" db:"dept_name"`        //部门名称
	Status   string `form:"status" db:"status"`             //状态
	baizeEntity.BaseEntityDQL
}

type SysDeptVo struct {
	DeptId    int64   `json:"deptId,string" db:"dept_id"`     //id
	ParentId  int64   `json:"parentId,string" db:"parent_id"` //上级id
	Ancestors *string `json:"ancestors" db:"ancestors"`       //祖级列表
	DeptName  *string `json:"deptName" db:"dept_name"`        //部门名称
	OrderNum  string  `json:"orderNum" db:"order_num"`        //排序
	Leader    *string `json:"leader" db:"leader"`             //负责人
	Phone     *string `json:"phone" db:"phone"`               //电话
	Email     *string `json:"email" db:"email"`               //邮箱
	Status    string  `json:"status" db:"status"`             //状态
	DelFlag   string  `json:"delFag" db:"del_flag"`           //删除标志
	baizeEntity.BaseEntity
}
type SysDeptAdd struct {
	DeptId    int64   `json:"DeptId,string" db:"dept_id"swaggerignore:"true"`
	ParentId  int64   `json:"parentId,string" db:"parent_id" binding:"required"` //上级id
	DeptName  string  `json:"deptName" db:"dept_name" binding:"required"`        //部门名称
	Ancestors string  `json:"ancestors" db:"ancestors"swaggerignore:"true"`
	OrderNum  *uint32 `json:"orderNum" db:"order_num"` //排序
	Leader    string  `json:"leader" db:"leader"`      //负责人
	Phone     string  `json:"phone" db:"phone"`        //电话
	Email     string  `json:"email" db:"email"`        //邮箱
	Status    string  `json:"status" db:"status"swaggerignore:"true"`
	baizeEntity.BaseEntityAdd
}
type SysDeptEdit struct {
	DeptId    int64   `json:"DeptId,string" db:"dept_id" binding:"required"` //id
	ParentId  int64   `json:"parentId,string" db:"parent_id"`                //上级id
	DeptName  string  `json:"deptName" db:"dept_name"`                       //部门名称
	Ancestors string  `json:"ancestors" db:"ancestors" swaggerignore:"true"`
	OrderNum  *uint32 `json:"orderNum" db:"order_num"` //排序
	Leader    string  `json:"leader" db:"leader"`      //负责人
	Phone     string  `json:"phone" db:"phone"`        //电话
	Email     string  `json:"email" db:"email"`        //邮箱
	Status    string  `json:"status" db:"status"`      //状态
	baizeEntity.BaseEntityEdit
}

type RoleDeptTree struct {
	CheckedKeys []string     `json:"checkedKeys"` //keys
	Depts       []*SysDeptVo `json:"depts"`       //部门
}
