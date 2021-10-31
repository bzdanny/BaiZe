package systemService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
	"strconv"
)

func SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, count *int64) {
	sysUserList, count = systemDao.SelectUserList(user)
	return
}
func UserExport(user *systemModels.SysUserDQL) (data []byte) {
	sysUserList, _ := systemDao.SelectUserList(user)
	rows := systemModels.SysUserDMLListToRows(sysUserList)
	return exceLize.SetRows(rows)
}

func SelectUserById(userId int64) (sysUser *systemModels.SysUserVo) {
	return systemDao.SelectUserById(userId)

}

func InsertUser(sysUser *systemModels.SysUserDML) {
	sysUser.UserId = snowflake.GenID()
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	systemDao.InsertUser(sysUser)

	InsertUserPost(sysUser)

	InsertUserRole(sysUser)

}

func UpdateUser(sysUser *systemModels.SysUserDML) {
	userId := sysUser.UserId

	systemDao.DeleteUserPostByUserId(userId)

	InsertUserPost(sysUser)

	systemDao.DeleteUserRoleByUserId(userId)

	InsertUserRole(sysUser)

	systemDao.UpdateUser(sysUser)

}

func UpdateuserStatus(sysUser *systemModels.SysUserDML) {
	systemDao.UpdateUser(sysUser)

}
func ResetPwd(sysUser *systemModels.SysUserDML) {
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	systemDao.UpdateUser(sysUser)

}

func InsertUserPost(user *systemModels.SysUserDML) {
	posts := user.PostIds
	if len(posts) != 0 {
		list := make([]*systemModels.SysUserPost, 0, len(posts))
		for _, postId := range posts {
			parseInt, _ := strconv.ParseInt(postId, 10, 64)
			post := systemModels.NewSysUserPost(user.UserId, parseInt)
			list = append(list, post)
		}
		systemDao.BatchUserPost(list)
	}

}

func InsertUserRole(user *systemModels.SysUserDML) {
	roles := user.RoleIds
	if len(roles) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roles))
		for _, roleId := range roles {
			parseInt, _ := strconv.ParseInt(roleId, 10, 64)
			role := systemModels.NewSysUserRole(user.UserId, parseInt)
			list = append(list, role)
		}
		systemDao.BatchUserRole(list)
	}

}

func CheckUserNameUnique(userName string) bool {
	return systemDao.CheckUserNameUnique(userName) > 0

}

func CheckPhoneUnique(user *systemModels.SysUserDML) bool {
	if user.Phonenumber == "" {
		return false
	}
	userId := systemDao.CheckPhoneUnique(user.Phonenumber)
	if userId == user.UserId || userId == 0 {
		return false
	}
	return true
}

func CheckEmailUnique(user *systemModels.SysUserDML) bool {
	if user.Email == "" {
		return false
	}
	userId := systemDao.CheckEmailUnique(user.Email)
	if userId == user.UserId || userId == 0 {
		return false
	}
	return true
}

func DeleteUserByIds(ids []int64) {

	systemDao.DeleteUserRole(ids)
	systemDao.DeleteUserPost(ids)
	systemDao.DeleteUserByIds(ids)

}

func UserImportData(rows [][]string, operName string, deptId *int64) (msg string, failureNum int) {
	successNum := 0
	list, failureMsg, failureNum := systemModels.RowsToSysUserDMLList(rows)
	password := bCryptPasswordEncoder.HashPassword("123456")
	for _, user := range list {
		unique := systemDao.CheckUserNameUnique(user.UserName)
		if unique < 1 {
			user.DeptId = deptId
			user.Password = password
			user.SetCreateBy(operName)
			systemDao.InsertUser(user)
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
func UpdateLoginInformation(userId int64, ip string) {
	systemDao.UpdateLoginInformation(userId, ip)
}
func UpdateUserAvatar(userId int64, avatar string) {
	systemDao.UpdateUserAvatar(userId, avatar)
}

func ResetUserPwd(userId int64, password string) {
	systemDao.ResetUserPwd(userId, bCryptPasswordEncoder.HashPassword(password))
}
func UpdateUserProfile(sysUser *systemModels.SysUserDML) {
	systemDao.UpdateUser(sysUser)

}
