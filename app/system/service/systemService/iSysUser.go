package systemService

import (
	"baize/app/system/models/systemModels"
)

type IUserService interface {
	SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, count *int64)
	UserExport(user *systemModels.SysUserDQL) (data []byte)
	SelectUserById(userId int64) (sysUser *systemModels.SysUserVo)
	InsertUser(sysUser *systemModels.SysUserDML)
	UpdateUser(sysUser *systemModels.SysUserDML)
	UpdateuserStatus(sysUser *systemModels.SysUserDML)
	ResetPwd(sysUser *systemModels.SysUserDML)
	CheckUserNameUnique(userName string) bool
	CheckPhoneUnique(user *systemModels.SysUserDML) bool
	CheckEmailUnique(user *systemModels.SysUserDML) bool
	DeleteUserByIds(ids []int64)
	UserImportData(rows [][]string, operName string, deptId *int64) (msg string, failureNum int)
	UpdateLoginInformation(userId int64, ip string)
	UpdateUserAvatar(userId int64, avatar string)
	ResetUserPwd(userId int64, password string)
	UpdateUserProfile(sysUser *systemModels.SysUserDML)
	MatchesPassword(rawPassword string, userId int64) bool
	InsertUserAuth(userId int64, roleIds []int64)
	ImportTemplate() (data []byte)
}
