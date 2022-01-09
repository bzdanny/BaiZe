package systemDao

import (
	"baize/app/common/mysql"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
)

type IUserDao interface {
	CheckUserNameUnique(userName string) int
	CheckPhoneUnique(phonenumber string) int64
	CheckEmailUnique(email string) int64
	InsertUser(sysUser *systemModels.SysUserDML, tx ...mysql.Transaction)
	UpdateUser(sysUser *systemModels.SysUserDML, tx ...mysql.Transaction)
	SelectUserByUserName(userName string) (loginUser *loginModels.User)
	SelectUserById(userId int64) (sysUser *systemModels.SysUserVo)
	SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total *int64)
	DeleteUserByIds(ids []int64, tx ...mysql.Transaction)
	UpdateLoginInformation(userId int64, ip string)
	UpdateUserAvatar(userId int64, avatar string)
	ResetUserPwd(userId int64, password string)
	SelectPasswordByUserId(userId int64) string
}
