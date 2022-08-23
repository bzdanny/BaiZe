package systemModels

type SysUserRole struct {
	UserId int64 `json:"userId,string" db:"user_id"`
	RoleId int64 `json:"roleId,string" db:"role_id"`
}

func NewSysUserRole(userId int64, roleId int64) *SysUserRole {
	return &SysUserRole{UserId: userId, RoleId: roleId}
}
