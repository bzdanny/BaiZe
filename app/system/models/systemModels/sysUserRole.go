package systemModels

type SysUserRole struct {
	UserId int64 `db:"user_id"`
	RoleId int64 `db:"role_id"`
}

func NewSysUserRole(userId int64, roleId int64) *SysUserRole {
	return &SysUserRole{UserId: userId, RoleId: roleId}
}
