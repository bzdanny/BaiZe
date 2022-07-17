package systemServiceImpl

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bzdanny/BaiZe/app/system/dao/systemDao"
	"github.com/bzdanny/BaiZe/app/system/dao/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/models/loginModels"
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/bzdanny/BaiZe/baize/utils/bCryptPasswordEncoder"
	"github.com/bzdanny/BaiZe/baize/utils/exceLize"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
	"strconv"
)

type UserService struct {
	data        *datasource.Data
	userDao     systemDao.IUserDao
	userPostDao systemDao.IUserPostDao
	userRoleDao systemDao.IUserRoleDao
}

func NewUserService(data *datasource.Data, ud *systemDaoImpl.SysUserDao, upd *systemDaoImpl.SysUserPostDao, urd *systemDaoImpl.SysUserRoleDao) *UserService {
	return &UserService{
		data:        data,
		userDao:     ud,
		userPostDao: upd,
		userRoleDao: urd,
	}
}

func (userService *UserService) SelectUserByUserName(userName string) *loginModels.User {
	return userService.userDao.SelectUserByUserName(userService.data.GetSlaveDb(), userName)

}
func (userService *UserService) SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, count *int64) {
	return userService.userDao.SelectUserList(userService.data.GetSlaveDb(), user)
}
func (userService *UserService) UserExport(user *systemModels.SysUserDQL) (data []byte) {
	sysUserList, _ := userService.userDao.SelectUserList(userService.data.GetSlaveDb(), user)
	return exceLize.SetRows(systemModels.SysUserListToRows(sysUserList))
}
func (userService *UserService) ImportTemplate() (data []byte) {
	f := excelize.NewFile()
	template := systemModels.SysUserImportTemplate()
	f.SetSheetRow("Sheet1", "A1", &template)
	buffer, _ := f.WriteToBuffer()
	return buffer.Bytes()

}

func (userService *UserService) SelectUserById(userId int64) (sysUser *systemModels.SysUserVo) {
	return userService.userDao.SelectUserById(userService.data.GetSlaveDb(), userId)

}

func (userService *UserService) InsertUser(sysUser *systemModels.SysUserDML) {
	sysUser.UserId = snowflake.GenID()
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	tx, err := userService.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	userService.userDao.InsertUser(tx, sysUser)
	userService.insertUserPost(tx, sysUser)
	userService.insertUserRole(tx, sysUser)

}

func (userService *UserService) UpdateUser(sysUser *systemModels.SysUserDML) {
	userId := sysUser.UserId
	tx, err := userService.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	userService.userPostDao.DeleteUserPostByUserId(tx, userId)
	userService.insertUserPost(tx, sysUser)
	userService.userRoleDao.DeleteUserRoleByUserId(tx, userId)
	userService.insertUserRole(tx, sysUser)
	userService.userDao.UpdateUser(tx, sysUser)

}

func (userService *UserService) UpdateuserStatus(sysUser *systemModels.SysUserDML) {
	userService.userDao.UpdateUser(userService.data.GetMasterDb(), sysUser)

}
func (userService *UserService) ResetPwd(sysUser *systemModels.SysUserDML) {
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	userService.userDao.UpdateUser(userService.data.GetMasterDb(), sysUser)

}

func (userService *UserService) insertUserPost(db dataUtil.DB, user *systemModels.SysUserDML) {
	posts := user.PostIds
	if len(posts) != 0 {
		list := make([]*systemModels.SysUserPost, 0, len(posts))
		for _, postId := range posts {
			parseInt, _ := strconv.ParseInt(postId, 10, 64)
			post := systemModels.NewSysUserPost(user.UserId, parseInt)
			list = append(list, post)
		}
		userService.userPostDao.BatchUserPost(db, list)
	}

}

func (userService *UserService) insertUserRole(db dataUtil.DB, user *systemModels.SysUserDML) {
	roles := user.RoleIds
	if len(roles) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roles))
		for _, roleId := range roles {
			parseInt, _ := strconv.ParseInt(roleId, 10, 64)
			role := systemModels.NewSysUserRole(user.UserId, parseInt)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(db, list)
	}

}

func (userService *UserService) CheckUserNameUnique(userName string) bool {
	return userService.userDao.CheckUserNameUnique(userService.data.GetSlaveDb(), userName) > 0

}

func (userService *UserService) CheckPhoneUnique(user *systemModels.SysUserDML) bool {
	if user.Phonenumber == "" {
		return false
	}
	userId := userService.userDao.CheckPhoneUnique(userService.data.GetSlaveDb(), user.Phonenumber)
	if userId == user.UserId || userId == 0 {
		return false
	}
	return true
}

func (userService *UserService) CheckEmailUnique(user *systemModels.SysUserDML) bool {
	if user.Email == "" {
		return false
	}
	userId := userService.userDao.CheckEmailUnique(userService.data.GetSlaveDb(), user.Email)
	if userId == user.UserId || userId == 0 {
		return false
	}
	return true
}

func (userService *UserService) DeleteUserByIds(ids []int64) {
	tx, err := userService.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	userService.userRoleDao.DeleteUserRole(tx, ids)
	userService.userPostDao.DeleteUserPost(tx, ids)
	userService.userDao.DeleteUserByIds(tx, ids)

}

func (userService *UserService) UserImportData(rows [][]string, operName string, deptId *int64) (msg string, failureNum int) {
	successNum := 0
	list, failureMsg, failureNum := systemModels.RowsToSysUserDMLList(rows)
	password := bCryptPasswordEncoder.HashPassword("123456")
	tx, err := userService.data.GetMasterDb().Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	for _, user := range list {
		unique := userService.userDao.CheckUserNameUnique(tx, user.UserName)
		if unique < 1 {
			user.DeptId = deptId
			user.Password = password
			//TODO
			//user.SetCreateBy(operName)
			userService.userDao.InsertUser(tx, user)
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
func (userService *UserService) UpdateLoginInformation(userId int64, ip string) {
	userService.userDao.UpdateLoginInformation(userService.data.GetMasterDb(), userId, ip)
}
func (userService *UserService) UpdateUserAvatar(userId int64, avatar string) {
	userService.userDao.UpdateUserAvatar(userService.data.GetMasterDb(), userId, avatar)
}

func (userService *UserService) ResetUserPwd(userId int64, password string) {
	userService.userDao.ResetUserPwd(userService.data.GetMasterDb(), userId, bCryptPasswordEncoder.HashPassword(password))
}
func (userService *UserService) UpdateUserProfile(sysUser *systemModels.SysUserDML) {
	userService.userDao.UpdateUser(userService.data.GetMasterDb(), sysUser)

}
func (userService *UserService) MatchesPassword(rawPassword string, userId int64) bool {

	return bCryptPasswordEncoder.CheckPasswordHash(rawPassword, userService.userDao.SelectPasswordByUserId(userService.data.GetSlaveDb(), userId))
}
func (userService *UserService) InsertUserAuth(userId int64, roleIds []int64) {
	userService.userRoleDao.DeleteUserRoleByUserId(userService.data.GetMasterDb(), userId)
	if len(roleIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roleIds))
		for _, roleId := range roleIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(userService.data.GetMasterDb(), list)
	}
}
