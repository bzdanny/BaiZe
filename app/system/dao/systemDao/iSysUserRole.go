package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IUserRoleDao interface {
	DeleteUserRole(db dataUtil.DB, ids []int64)
	BatchUserRole(db dataUtil.DB, users []*systemModels.SysUserRole)
	DeleteUserRoleByUserId(db dataUtil.DB, userId int64)
	CountUserRoleByRoleId(db dataUtil.DB, ids []int64) int
	DeleteUserRoleInfo(db dataUtil.DB, userRole *systemModels.SysUserRole)
	DeleteUserRoleInfos(db dataUtil.DB, roleId int64, userIds []int64)
}
