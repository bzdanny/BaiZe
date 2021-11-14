package commonModels

import (
	"baize/app/constant/dataScopeAspect"
	"baize/app/system/models/loginModels"
	"baize/app/utils/stringUtils"
	"baize/app/utils/unix"
	"fmt"
	"strconv"
)

type BaseEntity struct {
	CreateBy   string          `json:"createBy" db:"create_by"`
	CreateTime *unix.BaiZeTime `json:"createTime" db:"create_time"`
	UpdateBy   string          `json:"updateBy" db:"update_by"`
	UpdateTime *unix.BaiZeTime `json:"updateTime" db:"update_time"`
}

type BaseEntityDML struct {
	CreateBy string `db:"create_by"`
	UpdateBy string `db:"update_by"`
}

func (b *BaseEntityDML) SetCreateBy(userName string) {
	b.CreateBy = userName
	b.UpdateBy = userName
}

func (b *BaseEntityDML) SetUpdateBy(userName string) {
	b.UpdateBy = userName
}

type BaseEntityDQL struct {
	CreateTime *unix.BaiZeTime `db:"create_time"`
	UpdateTime *unix.BaiZeTime `db:"update_time"`
	CreateBy   string          `db:"create_by"`
	UpdateBy   string          `db:"update_by"`
	DataScope  string
	Limit      string
	Offset     int64
	Size       int64
}

func (b *BaseEntityDQL) SetLimit(p *pageDomain) {
	var limitString string = " "
	if p.OrderBy != "" {
		limitString += stringUtils.ToUnderScoreCase(p.OrderBy) + " " + p.IsAsc
	}

	limitString += " limit " + strconv.FormatInt((p.Page-1)*p.Size, 10) + "," + strconv.FormatInt(p.Size, 10)
	b.Limit = limitString
	b.Offset = (p.Page - 1) * p.Size
	b.Size = p.Size
}

func (b *BaseEntityDQL) SetDataScope(user *loginModels.LoginUser, deptAlias string, userAlias string) {
	var sqlString string
	for _, role := range user.User.Roles {

		switch role.DataScope {
		case dataScopeAspect.DataScopeAll:
			sqlString = ""
			break
		case dataScopeAspect.DataScopeCustom:
			sqlString += fmt.Sprintf(" OR %s.dept_id IN ( SELECT dept_id FROM sys_role_dept WHERE role_id = %d ) ", deptAlias, role.RoleId)
		case dataScopeAspect.DataScopeDept:
			sqlString += fmt.Sprintf(" OR %s.dept_id = %d ", deptAlias, user.User.DeptId)
		case dataScopeAspect.DataScopeDeptAndChild:
			sqlString += fmt.Sprintf(" OR %s.dept_id IN ( SELECT dept_id FROM sys_dept WHERE dept_id = %d or find_in_set( %d , ancestors ) ) ", deptAlias, user.User.DeptId, user.User.DeptId)
		case dataScopeAspect.DataScopeSelf:
			if userAlias != "" {
				sqlString += fmt.Sprintf(" OR %s.user_id = %d ", userAlias, user.User.UserId)
			} else {
				sqlString += " OR 1=0 "
			}

		}

	}
	if sqlString != "" {
		b.DataScope = " (" + sqlString[4:] + ")"
	}

}
