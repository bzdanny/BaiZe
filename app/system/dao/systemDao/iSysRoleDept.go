package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
)

type IRoleDeptDao interface {
	DeleteRoleDept(ids []int64, tx ...mysql.Transaction)
	DeleteRoleDeptByRoleId(id int64, tx ...mysql.Transaction)
	BatchRoleDept(list []*systemModels.SysRoleDept, tx ...mysql.Transaction)
}
