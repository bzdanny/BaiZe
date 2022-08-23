package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type SysUserDao struct {
}

func GetSysUserDao() *SysUserDao {
	return &SysUserDao{}
}

func (userDao *SysUserDao) CheckUserNameUnique(db dataUtil.DB, userName string) int {
	var count = 0
	err := db.Get(&count, "select count(*) from sys_user where user_name = ?", userName)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic(err)
	}
	return count
}
func (userDao *SysUserDao) CheckPhoneUnique(db dataUtil.DB, phonenumber string) int64 {
	var userId int64 = 0
	err := db.Get(&userId, "select user_id from sys_user where phonenumber = ?", phonenumber)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic(err)
	}
	return userId
}

func (userDao *SysUserDao) CheckEmailUnique(db dataUtil.DB, email string) int64 {
	var userId int64 = 0
	err := db.Get(&userId, "select user_id from sys_user where email = ?", email)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic(err)
	}
	return userId
}

func (userDao *SysUserDao) InsertUser(db dataUtil.DB, sysUser *systemModels.SysUserAdd) {
	insertSQL := `insert into sys_user(user_id,user_name,nick_name,sex,password,status,create_by,create_time,update_by,update_time %s)
					values(:user_id,:user_name,:nick_name,:sex,:password,:status,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysUser.DeptId != nil {
		key += ",dept_id"
		value += ",:dept_id"
	}
	if sysUser.Email != "" {
		key += ",email"
		value += ",:email"
	}
	if sysUser.Avatar != "" {
		key += ",avatar"
		value += ",:avatar"
	}
	if sysUser.Phonenumber != "" {
		key += ",phonenumber"
		value += ",:phonenumber"
	}
	if sysUser.Remake != "" {
		key += ",remake"
		value += ",:remake"
	}
	insertStr := fmt.Sprintf(insertSQL, key, value)

	_, err := db.NamedExec(insertStr, sysUser)

	if err != nil {
		panic(err)
	}
}

func (userDao *SysUserDao) UpdateUser(db dataUtil.DB, sysUser *systemModels.SysUserEdit) {
	updateSQL := `update sys_user set update_time = now() , update_by = :update_by`

	if sysUser.DeptId != nil {
		updateSQL += ",dept_id = :dept_id"
	}
	if sysUser.Email != "" {
		updateSQL += ",email = :email"
	}

	if sysUser.Avatar != "" {
		updateSQL += ",avatar = :avatar"
	}
	if sysUser.Phonenumber != "" {
		updateSQL += ",phonenumber = :phonenumber"
	}
	if sysUser.Remake != "" {
		updateSQL += ",remake = :remake"
	}
	if sysUser.NickName != "" {
		updateSQL += ",nick_name = :nick_name"
	}
	if sysUser.Sex != "" {
		updateSQL += ",sex = :sex"
	}
	if sysUser.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where user_id = :user_id"

	_, err := db.NamedExec(updateSQL, sysUser)
	if err != nil {
		panic(err)
	}
}

func (userDao *SysUserDao) SelectUserByUserName(db dataUtil.DB, userName string) (loginUser *systemModels.User) {
	sqlStr := `select u.user_id, u.dept_id, u.user_name, u.nick_name, u.email, u.avatar, u.phonenumber, u.password, u.sex, u.status, u.del_flag, u.login_ip, u.login_date,u.remark, u.create_time,
         d.parent_id, d.dept_name
        from sys_user u
		    left join sys_dept d on u.dept_id = d.dept_id
			where u.user_name = ?			
			`

	loginUser = new(systemModels.User)
	err := db.Get(loginUser, sqlStr, userName)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (userDao *SysUserDao) SelectUserById(db dataUtil.DB, userId int64) (sysUser *systemModels.SysUserVo) {
	sqlStr := `select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark, d.dept_name, d.leader,  
        r.role_id
        from sys_user u
		    left join sys_dept d on u.dept_id = d.dept_id
		    left join sys_user_role ur on u.user_id = ur.user_id
		    left join sys_role r on r.role_id = ur.role_id		
			where u.user_id = ?
			`

	sysUser = new(systemModels.SysUserVo)
	err := db.Get(sysUser, sqlStr, userId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (userDao *SysUserDao) SelectUserList(db dataUtil.DB, user *systemModels.SysUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	sql := `select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark, d.dept_name, d.leader
			 from sys_user u left join sys_dept d on u.dept_id = d.dept_id where u.del_flag = '0'`
	if user.UserName != "" {
		sql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Status != "" {
		sql += " AND  u.status = :status"
	}
	if user.Phonenumber != "" {
		sql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	if user.BeginTime != "" {
		sql += " AND date_format(u.create_time,'%y%m%d') >= date_format(:begin_time,'%y%m%d')"
	}
	if user.EndTime != "" {
		sql += " AND date_format(u.create_time,'%y%m%d') <= date_format(:end_time,'%y%m%d')"
	}
	if user.DeptId != nil {
		sql += " AND (u.dept_id = :dept_id OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(:dept_id, ancestors) ))"
	}
	if user.DataScope != "" {
		sql += " AND " + user.DataScope
	}
	return dataUtil.NamedQueryListAndTotal(db, list, user, sql, "", "")

}

func (userDao *SysUserDao) DeleteUserByIds(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("update sys_user set del_flag = '2' where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (userDao *SysUserDao) UpdateLoginInformation(db dataUtil.DB, userId int64, ip string) {
	_, err := db.Exec(`update sys_user set login_date = now() , login_ip = ?  where user_id = ?`, ip, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) UpdateUserAvatar(db dataUtil.DB, userId int64, avatar string) {
	_, err := db.Exec(`update sys_user set avatar = ?  where user_id = ?`, avatar, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) ResetUserPwd(db dataUtil.DB, userId int64, password string) {
	_, err := db.Exec(`update sys_user set password = ?  where user_id = ?`, password, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) SelectPasswordByUserId(db dataUtil.DB, userId int64) string {
	sqlStr := `select password
        from sys_user 
			where user_id = ?			
			`

	password := new(string)
	err := db.Get(password, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return *password
}
