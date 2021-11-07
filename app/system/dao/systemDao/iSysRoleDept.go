package systemDao

import (
	"baize/app/system/models/systemModels"
)

type IRoleDeptDao interface {
	DeleteRoleDept(ids []int64)
	DeleteRoleDeptByRoleId(id int64)
	BatchRoleDept(list []*systemModels.SysRoleDept)
}
