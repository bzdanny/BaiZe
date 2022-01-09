package systemDao

import (
	"baize/app/common/datasource"
	"baize/app/system/models/systemModels"
)

type IRoleDeptDao interface {
	DeleteRoleDept(ids []int64, tx ...datasource.Transaction)
	DeleteRoleDeptByRoleId(id int64, tx ...datasource.Transaction)
	BatchRoleDept(list []*systemModels.SysRoleDept, tx ...datasource.Transaction)
}
