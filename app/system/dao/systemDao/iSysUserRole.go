package systemDao

import (
	"baize/app/common/datasource"
	"baize/app/system/models/systemModels"
)

type IUserRoleDao interface {
	DeleteUserRole(ids []int64, tx ...datasource.Transaction)
	BatchUserRole(users []*systemModels.SysUserRole, tx ...datasource.Transaction)
	DeleteUserRoleByUserId(userId int64, tx ...datasource.Transaction)
	CountUserRoleByRoleId(ids []int64) int
	DeleteUserRoleInfo(userRole *systemModels.SysUserRole)
	DeleteUserRoleInfos(roleId int64, userIds []int64)
}
