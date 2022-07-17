package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IRoleDeptDao interface {
	DeleteRoleDept(db dataUtil.DB, ids []int64)
	DeleteRoleDeptByRoleId(db dataUtil.DB, id int64)
	BatchRoleDept(db dataUtil.DB, list []*systemModels.SysRoleDept)
}
