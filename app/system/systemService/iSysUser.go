package systemService

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"mime/multipart"
)

type IUserService interface {
	SelectUserByUserName(userName string) *systemModels.User
	SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, count *int64)
	UserExport(user *systemModels.SysUserDQL) (data []byte)
	SelectUserById(userId int64) (sysUser *systemModels.SysUserVo)
	InsertUser(sysUser *systemModels.SysUserAdd)
	UpdateUser(sysUser *systemModels.SysUserEdit)
	UpdateUserStatus(sysUser *systemModels.EditUserStatus)
	ResetPwd(userId int64, password string)
	CheckUserNameUnique(userName string) bool
	CheckPhoneUnique(id int64, phonenumber string) bool
	CheckEmailUnique(id int64, email string) bool
	DeleteUserByIds(ids []int64)
	UserImportData(rows [][]string, operName string, deptId *int64) (msg string, failureNum int)
	UpdateLoginInformation(userId int64, ip string)
	UpdateUserAvatar(loginUser *systemModels.LoginUser, file *multipart.FileHeader) string
	ResetUserPwd(userId int64, password string)
	UpdateUserProfile(sysUser *systemModels.SysUserEdit)
	MatchesPassword(rawPassword string, userId int64) bool
	InsertUserAuth(userId int64, roleIds []int64)
	ImportTemplate() (data []byte)
}
