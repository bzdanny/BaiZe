package baizeEntity

import (
	"fmt"
	"github.com/bzdanny/BaiZe/app/constant/dataScopeAspect"
	"github.com/bzdanny/BaiZe/baize/utils/stringUtils"
	"math"
	"strconv"
)

type BaseEntity struct {
	CreateBy   int64      `json:"createBy" db:"create_by"`                           //创建人
	CreateTime *BaiZeTime `json:"createTime" db:"create_time" swaggertype:"integer"` //创建时间
	UpdateBy   int64      `json:"updateBy" db:"update_by"`                           //修改人
	UpdateTime *BaiZeTime `json:"updateTime" db:"update_time" swaggertype:"integer"` //修改时间
}
type BaseEntityAdd struct {
	CreateBy   int64      `db:"create_by" swaggerignore:"true"`
	CreateTime *BaiZeTime `db:"create_time" swaggerignore:"true" swaggertype:"integer"`
	UpdateBy   int64      `db:"update_by" swaggerignore:"true"`
	UpdateTime *BaiZeTime `db:"update_time" swaggerignore:"true" swaggertype:"integer"`
}

type BaseEntityEdit struct {
	UpdateBy   int64      `db:"update_by" swaggerignore:"true"`
	UpdateTime *BaiZeTime `db:"update_time" swaggerignore:"true"`
}

func (b *BaseEntityAdd) SetCreateBy(userId int64) {
	b.CreateBy = userId
	kt := NewBaiZeTime()
	b.CreateTime = kt
	b.UpdateBy = userId
	b.UpdateTime = kt
}

func (b *BaseEntityEdit) SetUpdateBy(userId int64) {
	b.UpdateBy = userId
	b.UpdateTime = NewBaiZeTime()
}

type BaseEntityDQL struct {
	DataScope string `swaggerignore:"true"`
	OrderBy   string `form:"orderBy" `          //排序字段
	IsAsc     string `form:"isAsc" `            //排序规则  降序desc   asc升序
	Page      int64  `form:"page" default:"1"`  //第几页
	Size      int64  `form:"size" default:"10"` //数量
}

func (b *BaseEntityDQL) GetLimit() string {
	if b.Page < 1 {
		b.Page = 1
	}
	if b.Size < 1 {
		b.Size = math.MaxInt
	} else if b.Size > math.MaxInt {
		b.Size = math.MaxInt
	}
	return " limit " + strconv.FormatInt((b.Page-1)*b.Size, 10) + "," + strconv.FormatInt(b.Size, 10)

}
func (b *BaseEntityDQL) GetOrder() string {
	if b.OrderBy != "" {
		return " order by " + stringUtils.ToUnderScoreCase(b.OrderBy) + " " + b.IsAsc
	}
	return ""
}

func (b *BaseEntityDQL) GetSize() int64 {
	if b.Size < 1 {
		b.Size = 10
	}
	if b.Size > 10000 {
		b.Size = 10000
	}
	return b.Size
}
func (b *BaseEntityDQL) GetOffset() int64 {
	if b.Page < 1 {
		b.Page = 1
	}
	return (b.Page - 1) * b.Size
}

type User interface {
	GetRoles() []*Role
	GetDeptId() int64
}
type Role struct {
	RoleId    int64  `db:"role_id"`
	DataScope string `db:"data_scope"`
}

func (b *BaseEntityDQL) SetDataScope(user User, deptAlias string, userAlias string) {
	roles := user.GetRoles()
	var sqlString string
	for _, role := range roles {

		switch role.DataScope {
		case dataScopeAspect.DataScopeAll:
			sqlString = ""
			break
		case dataScopeAspect.DataScopeCustom:
			sqlString += fmt.Sprintf(" OR %s.dept_id IN ( SELECT dept_id FROM sys_role_dept WHERE role_id = %d ) ", deptAlias, role.RoleId)
		case dataScopeAspect.DataScopeDept:
			sqlString += fmt.Sprintf(" OR %s.dept_id = %d ", deptAlias, user.GetDeptId())
		case dataScopeAspect.DataScopeDeptAndChild:
			sqlString += fmt.Sprintf(" OR %s.dept_id IN ( SELECT dept_id FROM sys_dept WHERE dept_id = %d or find_in_set( %d , ancestors ) ) ", deptAlias, user.GetDeptId(), user.GetDeptId())
		case dataScopeAspect.DataScopeSelf:
			if userAlias != "" {
				sqlString += fmt.Sprintf(" OR %s.user_id = %d ", userAlias, user.GetDeptId())
			} else {
				sqlString += " OR 1=0 "
			}

		}

	}
	if sqlString != "" {
		b.DataScope = " (" + sqlString[4:] + ")"
	}

}
