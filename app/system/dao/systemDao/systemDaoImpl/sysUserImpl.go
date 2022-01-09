package systemDaoImpl

import (
	"baize/app/common/datasource"
	"baize/app/constant/constants"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var sysUserDaoImpl *sysUserDao

type sysUserDao struct {
}

func init() {
	sysUserDaoImpl = &sysUserDao{}
}

func GetSysUserDao() *sysUserDao {
	return sysUserDaoImpl
}

func (userDao *sysUserDao) CheckUserNameUnique(userName string) int {
	var count = 0
	err := datasource.GetMasterDb().Get(&count, "select count(*) from sys_user where user_name = ?", userName)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic(err)
	}
	return count
}
func (userDao *sysUserDao) CheckPhoneUnique(phonenumber string) int64 {
	var userId int64 = 0
	err := datasource.GetMasterDb().Get(&userId, "select user_id from sys_user where phonenumber = ?", phonenumber)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic(err)
	}
	return userId
}

func (userDao *sysUserDao) CheckEmailUnique(email string) int64 {
	var userId int64 = 0
	err := datasource.GetMasterDb().Get(&userId, "select user_id from sys_user where email = ?", email)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic(err)
	}
	return userId
}

func (userDao *sysUserDao) InsertUser(sysUser *systemModels.SysUserDML, tx ...datasource.Transaction) {
	insertSQL := `insert into sys_user(user_id,user_name,nick_name,sex,password,status,create_by,create_time,update_by,update_time %s)
					values(:user_id,:user_name,:nick_name,:sex,:password,:status,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysUser.DeptId != nil {
		key += ",dept_id"
		value += ",:dept_id"
	}
	if sysUser.Email != "" {
		key += "email"
		value += ":email"
	}
	if sysUser.Avatar != "" {
		key += "avatar"
		value += ":avatar"
	}
	if sysUser.Phonenumber != "" {
		key += "phonenumber"
		value += ":phonenumber"
	}
	if sysUser.Remake != "" {
		key += "remake"
		value += ":remake"
	}
	insertStr := fmt.Sprintf(insertSQL, key, value)
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec(insertStr, sysUser)

	if err != nil {
		panic(err)
	}
}

func (userDao *sysUserDao) UpdateUser(sysUser *systemModels.SysUserDML, tx ...datasource.Transaction) {
	updateSQL := `update sys_user set update_time = now() , update_by = :update_by`

	if sysUser.DeptId != nil {
		updateSQL += ",dept_id = :dept_id"
	}
	if sysUser.Email != "" {
		updateSQL += ",email = :email"
	}
	if sysUser.Password != "" {
		updateSQL += ",password = :password"
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
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec(updateSQL, sysUser)
	if err != nil {
		panic(err)
	}
}

func (userDao *sysUserDao) SelectUserByUserName(userName string) (loginUser *loginModels.User) {
	sqlStr := `select u.user_id, u.dept_id, u.user_name, u.nick_name, u.email, u.avatar, u.phonenumber, u.password, u.sex, u.status, u.del_flag, u.login_ip, u.login_date,u.remark, u.create_time,
         d.parent_id, d.dept_name
        from sys_user u
		    left join sys_dept d on u.dept_id = d.dept_id
			where u.user_name = ?			
			`

	loginUser = new(loginModels.User)
	err := datasource.GetMasterDb().Get(loginUser, sqlStr, userName)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (userDao *sysUserDao) SelectUserById(userId int64) (sysUser *systemModels.SysUserVo) {
	sqlStr := `select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark, d.dept_name, d.leader,  
        r.role_id
        from sys_user u
		    left join sys_dept d on u.dept_id = d.dept_id
		    left join sys_user_role ur on u.user_id = ur.user_id
		    left join sys_role r on r.role_id = ur.role_id		
			where u.user_id = ?
			`

	sysUser = new(systemModels.SysUserVo)
	err := datasource.GetMasterDb().Get(sysUser, sqlStr, userId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (userDao *sysUserDao) SelectUserList(user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total *int64) {
	sql := "select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark, d.dept_name, d.leader"

	whereSql := ` from sys_user u left join sys_dept d on u.dept_id = d.dept_id where u.del_flag = '0'`
	if user.UserName != "" {
		whereSql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Status != "" {
		whereSql += " AND  u.status = :status"
	}
	if user.Phonenumber != "" {
		whereSql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	if user.BeginTime != "" {
		whereSql += " AND date_format(u.create_time,'%y%m%d') >= date_format(:begin_time,'%y%m%d')"
	}
	if user.EndTime != "" {
		whereSql += " AND date_format(u.create_time,'%y%m%d') <= date_format(:end_time,'%y%m%d')"
	}
	if user.DeptId != nil {
		whereSql += " AND (u.dept_id = :dept_id OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(:dept_id, ancestors) ))"
	}
	if user.DataScope != "" {
		whereSql += " AND " + user.DataScope
	}
	countSql := constants.MysqlCount + whereSql

	countRow, err := datasource.GetMasterDb().NamedQuery(countSql, user)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	if *total > user.Offset {
		sysUserList = make([]*systemModels.SysUserVo, 0, user.Size)

		if user.Limit != "" {
			whereSql += user.Limit
		}
		listRows, err := datasource.GetMasterDb().NamedQuery(sql+whereSql, user)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			sysUser := new(systemModels.SysUserVo)
			listRows.StructScan(sysUser)
			sysUserList = append(sysUserList, sysUser)
		}
		defer listRows.Close()
	} else {
		sysUserList = make([]*systemModels.SysUserVo, 0, 0)
	}

	return
}

func (userDao *sysUserDao) DeleteUserByIds(ids []int64, tx ...datasource.Transaction) {
	query, i, err := sqlx.In("update sys_user set del_flag = '2' where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = datasource.GetMasterDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (userDao *sysUserDao) UpdateLoginInformation(userId int64, ip string) {
	_, err := datasource.GetMasterDb().Exec(`update sys_user set login_date = now() , login_ip = ?  where user_id = ?`, ip, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *sysUserDao) UpdateUserAvatar(userId int64, avatar string) {
	_, err := datasource.GetMasterDb().Exec(`update sys_user set avatar = ?  where user_id = ?`, avatar, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *sysUserDao) ResetUserPwd(userId int64, password string) {
	_, err := datasource.GetMasterDb().Exec(`update sys_user set password = ?  where user_id = ?`, password, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *sysUserDao) SelectPasswordByUserId(userId int64) string {
	sqlStr := `select password
        from sys_user 
			where user_id = ?			
			`

	password := new(string)
	err := datasource.GetMasterDb().Get(password, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return *password
}
