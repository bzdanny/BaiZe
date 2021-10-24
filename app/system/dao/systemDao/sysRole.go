package systemDao

import (
	mysql "baize/app/common/mysql"
	"baize/app/constant/constants"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var selectRoleSql = ` select distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,
            r.status, r.del_flag, r.create_time, r.remark`
var fromRoleSql = ` from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
	        left join sys_dept d on u.dept_id = d.dept_id`

func SelectRoleList(role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total *int64) {
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
	countSql := constants.MysqlCount + fromRoleSql + whereSql
	countRow, err := mysql.MysqlDb.NamedQuery(countSql, role)
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
		listRows, err := mysql.MysqlDb.NamedQuery(selectRoleSql+fromRoleSql+whereSql, role)
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
func SelectRoleById(roleId int64) (role *systemModels.SysRoleVo) {
	whereSql := ` where r.role_id = ?`
	role = new(systemModels.SysRoleVo)
	err := mysql.MysqlDb.Get(role, selectRoleSql+fromRoleSql+whereSql, roleId)
	if err != nil {
		panic(err)
	}
	return
}

func SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole) {
	sqlStr := `select  r.role_id, r.role_name, r.role_key,r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*systemModels.SysRole, 0, 2)
	err := mysql.MysqlDb.Select(&roles, sqlStr, userId)
	if err != nil {
		panic(err)
	}

	return
}

func SelectRolePermissionByUserId(userId int64) (roles []string) {
	sqlStr := `select   r.role_key
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]string, 0, 1)
	err := mysql.MysqlDb.Select(&roles, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
func SelectRoleIdAndDataScopeByUserId(userId int64) (roles []*loginModels.Role) {
	sqlStr := `select  r.role_id, r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*loginModels.Role, 0, 2)
	err := mysql.MysqlDb.Select(&roles, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func SelectRoleListByUserId(userId int64) (list []int64) {
	sqlStr := `select r.role_id
        from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := mysql.MysqlDb.Select(&list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
	return
}

func InsertRole(sysRole *systemModels.SysRoleDML) {
	insertSQL := `insert into sys_role(role_id,role_name,role_sort,create_by,create_time,update_by,update_time %s)
					values(:role_id,:role_name,:role_sort,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysRole.DataScope != "" {
		key += ",data_scope"
		value += ",:data_scope"
	}
	if sysRole.MenuCheckStrictly != "" {
		key += ",menu_check_strictly"
		value += ",:menu_check_strictly"
	}
	if sysRole.DeptCheckStrictly != "" {
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
	_, err := mysql.MysqlDb.NamedExec(insertStr, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func UpdateRole(sysRole *systemModels.SysRoleDML) {
	updateSQL := `update sys_role set update_time = now() , update_by = :update_by`

	if sysRole.RoleName != "" {
		updateSQL += ",role_name = :role_name"
	}
	if sysRole.RoleKey != "" {
		updateSQL += ",role_key = :role_key"
	}
	if sysRole.RoleSort != "" {
		updateSQL += ",role_sort = :role_sort"
	}
	if sysRole.DataScope != "" {
		updateSQL += ",data_scope = :data_scope"
	}
	if sysRole.MenuCheckStrictly != "" {
		updateSQL += ",menu_check_strictly = :menu_check_strictly"
	}
	if sysRole.DeptCheckStrictly != "" {
		updateSQL += ",dept_check_strictly = :dept_check_strictly"
	}
	if sysRole.Remake != "" {
		updateSQL += ",remake = :remake"
	}
	if sysRole.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where user_id = :user_id"

	_, err := mysql.MysqlDb.NamedExec(updateSQL, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func DeleteRoleByIds(ids []int64) {
	query, i, err := sqlx.In("update sys_role set del_flag = '2' where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = mysql.MysqlDb.Exec(query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func CheckRoleNameUnique(roleName string) int64 {
	var roleId int64 = 0
	err := mysql.MysqlDb.Get(&roleId, "select role_id from sys_role where role_name = ?", roleName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func CheckRoleKeyUnique(roleKey string) int64 {
	var roleId int64 = 0
	err := mysql.MysqlDb.Get(&roleId, "select role_id from sys_role where role_key = ?", roleKey)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}
