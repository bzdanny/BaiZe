package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		selectSql: ` select distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.Permission_check_strictly, r.dept_check_strictly, r.status, r.del_flag, r.create_time, r.remark from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
	        left join sys_dept d on u.dept_id = d.dept_id`,
	}
}

type SysRoleDao struct {
	selectSql string
}

func (rd *SysRoleDao) SelectRoleList(db dataUtil.DB, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, total *int64) {
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

	return dataUtil.NamedQueryListAndTotal(db, list, role, rd.selectSql+whereSql, "", "")
}
func (rd *SysRoleDao) SelectRoleById(db dataUtil.DB, roleId int64) (role *systemModels.SysRoleVo) {
	whereSql := ` where r.role_id = ?`
	role = new(systemModels.SysRoleVo)
	err := db.Get(role, rd.selectSql+whereSql, roleId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectBasicRolesByUserId(db dataUtil.DB, userId int64) (roles []*systemModels.SysRole) {
	sqlStr := `select  r.role_id, r.role_name, r.role_key,r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*systemModels.SysRole, 0, 2)
	err := db.Select(&roles, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectRolePermissionByUserId(db dataUtil.DB, userId int64) (roles []string) {
	sqlStr := `select   r.role_key
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]string, 0, 1)
	err := db.Select(&roles, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) SelectRoleIdAndDataScopeByUserId(db dataUtil.DB, userId int64) (roles []*baizeEntity.Role) {
	sqlStr := `select  r.role_id, r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*baizeEntity.Role, 0, 2)
	err := db.Select(&roles, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectRoleListByUserId(db dataUtil.DB, userId int64) (list []int64) {
	sqlStr := `select r.role_id
        from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := db.Select(&list, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) InsertRole(db dataUtil.DB, sysRole *systemModels.SysRoleAdd) {
	insertSQL := `insert into sys_role(role_id,role_name,role_key,role_sort,create_by,create_time,update_by,update_time %s)
					values(:role_id,:role_name,:role_key,:role_sort,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysRole.DataScope != "" {
		key += ",data_scope"
		value += ",:data_scope"
	}
	if sysRole.PermissionCheckStrictly != nil {
		key += ",Permission_check_strictly"
		value += ",:Permission_check_strictly"
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

	_, err := db.NamedExec(insertStr, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) UpdateRole(db dataUtil.DB, sysRole *systemModels.SysRoleEdit) {
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
	if sysRole.PermissionCheckStrictly != nil {
		updateSQL += ",permission_check_strictly = :permission_check_strictly"
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
	_, err := db.NamedExec(updateSQL, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) DeleteRoleByIds(db dataUtil.DB, ids []int64) {
	query, i, err := sqlx.In("update sys_role set del_flag = '2',role_name = concat(role_name,'(delete)')  where role_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) CheckRoleNameUnique(db dataUtil.DB, roleName string) int64 {
	var roleId int64 = 0
	err := db.Get(&roleId, "select role_id from sys_role where role_name = ?", roleName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func (rd *SysRoleDao) CheckRoleKeyUnique(db dataUtil.DB, roleKey string) int64 {
	var roleId int64 = 0
	err := db.Get(&roleId, "select role_id from sys_role where role_key = ?", roleKey)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}
func (rd *SysRoleDao) SelectAllocatedList(db dataUtil.DB, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
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

	return dataUtil.NamedQueryListAndTotal(db, list, user, selectStr+whereSql, "", "")

}

func (rd *SysRoleDao) SelectUnallocatedList(db dataUtil.DB, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
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

	return dataUtil.NamedQueryListAndTotal(db, list, user, selectStr+whereSql, "", "")

}
