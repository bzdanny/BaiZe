package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/system/models/systemModels"
)

type IUserRoleDao interface {
	DeleteUserRole(ids []int64, tx ...mysql.Transaction)
	BatchUserRole(users []*systemModels.SysUserRole, tx ...mysql.Transaction)
	DeleteUserRoleByUserId(userId int64, tx ...mysql.Transaction)
	CountUserRoleByRoleId(ids []int64) int
	DeleteUserRoleInfo(userRole *systemModels.SysUserRole)
	DeleteUserRoleInfos(roleId int64, userIds []int64)
}
