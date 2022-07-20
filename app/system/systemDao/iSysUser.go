package systemDao

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IUserDao interface {
	CheckUserNameUnique(db dataUtil.DB, userName string) int
	CheckPhoneUnique(db dataUtil.DB, phonenumber string) int64
	CheckEmailUnique(db dataUtil.DB, email string) int64
	InsertUser(db dataUtil.DB, sysUser *systemModels.SysUserAdd)
	UpdateUser(db dataUtil.DB, sysUser *systemModels.SysUserEdit)
	SelectUserByUserName(db dataUtil.DB, userName string) (loginUser *systemModels.User)
	SelectUserById(db dataUtil.DB, userId int64) (sysUser *systemModels.SysUserVo)
	SelectUserList(db dataUtil.DB, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total *int64)
	DeleteUserByIds(db dataUtil.DB, ids []int64)
	UpdateLoginInformation(db dataUtil.DB, userId int64, ip string)
	UpdateUserAvatar(db dataUtil.DB, userId int64, avatar string)
	ResetUserPwd(db dataUtil.DB, userId int64, password string)
	SelectPasswordByUserId(db dataUtil.DB, userId int64) string
}
