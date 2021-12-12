package systemServiceImpl

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/systemModels"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
	"strconv"
)

var userServiceImpl *userService = &userService{userDao: systemDaoImpl.GetSysUserDao(), userPostDao: systemDaoImpl.GetSysUserPostDao(), userRoleDao: systemDaoImpl.GetSysUserRoleDao()}

type userService struct {
	userDao     systemDao.IUserDao
	userPostDao systemDao.IUserPostDao
	userRoleDao systemDao.IUserRoleDao
}

func GetUserService() *userService {
	return userServiceImpl
}

func (userService *userService) SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, count *int64) {
	sysUserList, count = userService.userDao.SelectUserList(user)
	return
}
func (userService *userService) UserExport(user *systemModels.SysUserDQL) (data []byte) {
	sysUserList, _ := userService.userDao.SelectUserList(user)
	rows := systemModels.SysUserDMLListToRows(sysUserList)
	return exceLize.SetRows(rows)
}

func (userService *userService) SelectUserById(userId int64) (sysUser *systemModels.SysUserVo) {
	return userService.userDao.SelectUserById(userId)

}

func (userService *userService) InsertUser(sysUser *systemModels.SysUserDML) {
	sysUser.UserId = snowflake.GenID()
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	userService.userDao.InsertUser(sysUser)

	userService.InsertUserPost(sysUser)

	userService.insertUserRole(sysUser)

}

func (userService *userService) UpdateUser(sysUser *systemModels.SysUserDML) {
	userId := sysUser.UserId

	userService.userPostDao.DeleteUserPostByUserId(userId)

	userService.InsertUserPost(sysUser)

	userService.userRoleDao.DeleteUserRoleByUserId(userId)

	userService.insertUserRole(sysUser)

	userService.userDao.UpdateUser(sysUser)

}

func (userService *userService) UpdateuserStatus(sysUser *systemModels.SysUserDML) {
	userService.userDao.UpdateUser(sysUser)

}
func (userService *userService) ResetPwd(sysUser *systemModels.SysUserDML) {
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	userService.userDao.UpdateUser(sysUser)

}

func (userService *userService) InsertUserPost(user *systemModels.SysUserDML) {
	posts := user.PostIds
	if len(posts) != 0 {
		list := make([]*systemModels.SysUserPost, 0, len(posts))
		for _, postId := range posts {
			parseInt, _ := strconv.ParseInt(postId, 10, 64)
			post := systemModels.NewSysUserPost(user.UserId, parseInt)
			list = append(list, post)
		}
		userService.userPostDao.BatchUserPost(list)
	}

}

func (userService *userService) insertUserRole(user *systemModels.SysUserDML) {
	roles := user.RoleIds
	if len(roles) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roles))
		for _, roleId := range roles {
			parseInt, _ := strconv.ParseInt(roleId, 10, 64)
			role := systemModels.NewSysUserRole(user.UserId, parseInt)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(list)
	}

}

func (userService *userService) CheckUserNameUnique(userName string) bool {
	return userService.userDao.CheckUserNameUnique(userName) > 0

}

func (userService *userService) CheckPhoneUnique(user *systemModels.SysUserDML) bool {
	if user.Phonenumber == "" {
		return false
	}
	userId := userService.userDao.CheckPhoneUnique(user.Phonenumber)
	if userId == user.UserId || userId == 0 {
		return false
	}
	return true
}

func (userService *userService) CheckEmailUnique(user *systemModels.SysUserDML) bool {
	if user.Email == "" {
		return false
	}
	userId := userService.userDao.CheckEmailUnique(user.Email)
	if userId == user.UserId || userId == 0 {
		return false
	}
	return true
}

func (userService *userService) DeleteUserByIds(ids []int64) {

	userService.userRoleDao.DeleteUserRole(ids)
	userService.userPostDao.DeleteUserPost(ids)
	userService.userDao.DeleteUserByIds(ids)

}

func (userService *userService) UserImportData(rows [][]string, operName string, deptId *int64) (msg string, failureNum int) {
	successNum := 0
	list, failureMsg, failureNum := systemModels.RowsToSysUserDMLList(rows)
	password := bCryptPasswordEncoder.HashPassword("123456")
	for _, user := range list {
		unique := userService.userDao.CheckUserNameUnique(user.UserName)
		if unique < 1 {
			user.DeptId = deptId
			user.Password = password
			user.SetCreateBy(operName)
			userService.userDao.InsertUser(user)
			successNum++
		} else {
			failureNum++
			failureMsg += "<br/>账号 " + user.UserName + " 已存在"
		}
	}
	if failureNum > 0 {
		failureMsg = "很抱歉，导入失败！共 " + strconv.Itoa(failureNum) + " 条数据格式不正确，错误如下：" + failureMsg
		return failureMsg, failureNum
	}
	return "恭喜您，数据已全部导入成功！共 " + strconv.Itoa(successNum) + " 条。", 0
}
func (userService *userService) UpdateLoginInformation(userId int64, ip string) {
	userService.userDao.UpdateLoginInformation(userId, ip)
}
func (userService *userService) UpdateUserAvatar(userId int64, avatar string) {
	userService.userDao.UpdateUserAvatar(userId, avatar)
}

func (userService *userService) ResetUserPwd(userId int64, password string) {
	userService.userDao.ResetUserPwd(userId, bCryptPasswordEncoder.HashPassword(password))
}
func (userService *userService) UpdateUserProfile(sysUser *systemModels.SysUserDML) {
	userService.userDao.UpdateUser(sysUser)

}
func (userService *userService) MatchesPassword(rawPassword string, userId int64) bool {

	return bCryptPasswordEncoder.CheckPasswordHash(rawPassword, userService.userDao.SelectPasswordByUserId(userId))
}
func (userService *userService) InsertUserAuth(userId int64, roleIds []int64)  {
	userService.userRoleDao.DeleteUserRoleByUserId(userId)
	if len(roleIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roleIds))
		for _, roleId := range roleIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(list)
	}
}
