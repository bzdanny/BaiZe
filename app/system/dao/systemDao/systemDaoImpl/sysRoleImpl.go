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

var sysRoleDaoImpl *sysRoleDao

func init() {
	sysRoleDaoImpl = &sysRoleDao{
		selectSql: ` select distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly, r.status, r.del_flag, r.create_time, r.remark`,
		fromSql: ` from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
	        left join sys_dept d on u.dept_id = d.dept_id`,
	}
}

type sysRoleDao struct {
	selectSql string
	fromSql   string
}

func GetSysRoleDao() *sysRoleDao {
	return sysRoleDaoImpl
}

func (sysRoleDao *sysRoleDao) SelectRoleList(role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total *int64) {
	whereSql := " where r.del_flag = '0'"
	if role.RoleName != "" {
		whereSql += " AND r.role_name like concat('%', :role_name, '%')"
	}
	if role.Status != "" {
		whereSql += " AND r.status = :status"
	}
	if role.RoleKey != "" {
		whereSql += " AND r.role_key like concat('%', :roleKey, '%')"
	}

	if role.BeginTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
	}
	if role.EndTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
	}
	if role.DataScope != "" {
		whereSql += " AND " + role.DataScope
	}
	countSql := constants.MysqlCount + sysRoleDao.fromSql + whereSql
	countRow, err := datasource.GetMasterDb().NamedQuery(countSql, role)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	if *total > role.Offset {
		whereSql += " order by r.role_sort"
		roleList = make([]*systemModels.SysRoleVo, 0, 2)
		listRows, err := datasource.GetMasterDb().NamedQuery(sysRoleDao.selectSql+sysRoleDao.fromSql+whereSql, role)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			sysRole := new(systemModels.SysRoleVo)
			listRows.StructScan(sysRole)
			roleList = append(roleList, sysRole)
		}
	}
	return
}
func (sysRoleDao *sysRoleDao) SelectRoleById(roleId int64) (role *systemModels.SysRoleVo) {
	whereSql := ` where r.role_id = ?`
	role = new(systemModels.SysRoleVo)
	err := datasource.GetMasterDb().Get(role, sysRoleDao.selectSql+sysRoleDao.fromSql+whereSql, roleId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDao *sysRoleDao) SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole) {
	sqlStr := `select  r.role_id, r.role_name, r.role_key,r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*systemModels.SysRole, 0, 2)
	err := datasource.GetMasterDb().Select(&roles, sqlStr, userId)
	if err != nil {
		panic(err)
	}

	return
}

func (sysRoleDao *sysRoleDao) SelectRolePermissionByUserId(userId int64) (roles []string) {
	sqlStr := `select   r.role_key
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]string, 0, 1)
	err := datasource.GetMasterDb().Select(&roles, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
func (sysRoleDao *sysRoleDao) SelectRoleIdAndDataScopeByUserId(userId int64) (roles []*loginModels.Role) {
	sqlStr := `select  r.role_id, r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*loginModels.Role, 0, 2)
	err := datasource.GetMasterDb().Select(&roles, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDao *sysRoleDao) SelectRoleListByUserId(userId int64) (list []int64) {
	sqlStr := `select r.role_id
        from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := datasource.GetMasterDb().Select(&list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
	return
}

func (sysRoleDao *sysRoleDao) InsertRole(sysRole *systemModels.SysRoleDML, tx ...datasource.Transaction) {
	insertSQL := `insert into sys_role(role_id,role_name,role_key,role_sort,create_by,create_time,update_by,update_time %s)
					values(:role_id,:role_name,:role_key,:role_sort,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysRole.DataScope != "" {
		key += ",data_scope"
		value += ",:data_scope"
	}
	if sysRole.MenuCheckStrictly != nil {
		key += ",menu_check_strictly"
		value += ",:menu_check_strictly"
	}
	if sysRole.DeptCheckStrictly != nil {
		key += ",dept_check_strictly"
		value += ",:dept_check_strictly"
	}
	if sysRole.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if sysRole.Remake != "" {
		key += ",remake"
		value += ",:remake"
	}
	insertStr := fmt.Sprintf(insertSQL, key, value)
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec(insertStr, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDao *sysRoleDao) UpdateRole(sysRole *systemModels.SysRoleDML, tx ...datasource.Transaction) {
	updateSQL := `update sys_role set update_time = now() , update_by = :update_by`

	if sysRole.RoleName != "" {
		updateSQL += ",role_name = :role_name"
	}
	if sysRole.RoleKey != "" {
		updateSQL += ",role_key = :role_key"
	}
	if sysRole.RoleSort != -1 {
		updateSQL += ",role_sort = :role_sort"
	}
	if sysRole.DataScope != "" {
		updateSQL += ",data_scope = :data_scope"
	}
	if sysRole.MenuCheckStrictly != nil {
		updateSQL += ",menu_check_strictly = :menu_check_strictly"
	}
	if sysRole.DeptCheckStrictly != nil {
		updateSQL += ",dept_check_strictly = :dept_check_strictly"
	}
	if sysRole.Remake != "" {
		updateSQL += ",remake = :remake"
	}
	if sysRole.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where role_id = :role_id"
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err := db.NamedExec(updateSQL, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDao *sysRoleDao) DeleteRoleByIds(ids []int64, tx ...datasource.Transaction) {
	query, i, err := sqlx.In("update sys_role set del_flag = '2' where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	var db datasource.Transaction
	if len(tx) == 1 {
		db = tx[0]
	} else {
		db = datasource.GetMasterDb()
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysRoleDao *sysRoleDao) CheckRoleNameUnique(roleName string) int64 {
	var roleId int64 = 0
	err := datasource.GetMasterDb().Get(&roleId, "select role_id from sys_role where role_name = ?", roleName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func (sysRoleDao *sysRoleDao) CheckRoleKeyUnique(roleKey string) int64 {
	var roleId int64 = 0
	err := datasource.GetMasterDb().Get(&roleId, "select role_id from sys_role where role_key = ?", roleKey)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}
func (sysRoleDao *sysRoleDao) SelectAllocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	selectStr := ` select distinct u.user_id, u.dept_id, u.user_name, u.nick_name, u.email, u.phonenumber, u.status, u.create_time`

	whereSql := ` from sys_user u
			 left join sys_dept d on u.dept_id = d.dept_id
			 left join sys_user_role ur on u.user_id = ur.user_id
			 left join sys_role r on r.role_id = ur.role_id where u.del_flag = '0' and r.role_id =:role_id`
	if user.UserName != "" {
		whereSql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Phonenumber != "" {
		whereSql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	if user.DataScope != "" {
		whereSql += " AND " + user.DataScope
	}
	countRow, err := datasource.GetMasterDb().NamedQuery(constants.MysqlCount+whereSql, user)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	if *total > user.Offset {
		if user.Limit != "" {
			whereSql += user.Limit
		}
		list = make([]*systemModels.SysUserVo, 0, 2)
		listRows, err := datasource.GetMasterDb().NamedQuery(selectStr+whereSql, user)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			userVo := new(systemModels.SysUserVo)
			listRows.StructScan(userVo)
			list = append(list, userVo)
		}
	}
	return
}

func (sysRoleDao *sysRoleDao) SelectUnallocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	selectStr := ` select distinct u.user_id, u.dept_id, u.user_name, u.nick_name, u.email, u.phonenumber, u.status, u.create_time`

	whereSql := `  from sys_user u
			 left join sys_dept d on u.dept_id = d.dept_id
			 left join sys_user_role ur on u.user_id = ur.user_id
			 left join sys_role r on r.role_id = ur.role_id
	    where u.del_flag = '0' and (r.role_id != :role_id or r.role_id IS NULL)
	    and u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = :role_id)`
	if user.UserName != "" {
		whereSql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Phonenumber != "" {
		whereSql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	if user.DataScope != "" {
		whereSql += " AND " + user.DataScope
	}
	countRow, err := datasource.GetMasterDb().NamedQuery(constants.MysqlCount+whereSql, user)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	if *total > user.Offset {
		if user.Limit != "" {
			whereSql += user.Limit
		}
		list = make([]*systemModels.SysUserVo, 0, 2)
		listRows, err := datasource.GetMasterDb().NamedQuery(selectStr+whereSql, user)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			userVo := new(systemModels.SysUserVo)
			listRows.StructScan(userVo)
			list = append(list, userVo)
		}
	}
	return
}
