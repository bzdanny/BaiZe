package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type SysPermissionDao struct {
	PermissionSql string
}

func NewSysPermissionDao() *SysPermissionDao {
	return &SysPermissionDao{
		PermissionSql: `select distinct m.permission_id, m.parent_id, m.permission_name, m.status, ifnull(m.perms,'') as perms, m.create_time from sys_Permission m`,
	}
}

func (md *SysPermissionDao) SelectPermissionById(db dataUtil.DB, PermissionId int64) (Permission *systemModels.SysPermissionVo) {
	Permission = new(systemModels.SysPermissionVo)
	err := db.Get(Permission, md.PermissionSql+` where permission_id = ?`, PermissionId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (md *SysPermissionDao) SelectPermissionList(db dataUtil.DB, Permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo) {
	whereSql := ``
	if Permission.PermissionName != "" {
		whereSql += " AND permission_name like concat('%', :permission_name, '%')"
	}
	if Permission.PermissionName != "" {
		whereSql += "  AND status = :status"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	whereSql += " order by m.permission_id"

	return dataUtil.NamedQueryList(db, list, Permission, md.PermissionSql+whereSql, "", "")

}

func (md *SysPermissionDao) SelectPermissionListByUserId(db dataUtil.DB, Permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo) {
	whereSql := ` left join sys_role_Permission rm on m.Permission_id = rm.Permission_id
		left join sys_user_role ur on rm.role_id = ur.role_id
		left join sys_role ro on ur.role_id = ro.role_id
		where ur.user_id = :user_id`
	if Permission.PermissionName != "" {
		whereSql += " AND Permission_name like concat('%', :Permission_name, '%')"
	}
	if Permission.PermissionName != "" {
		whereSql += "  AND status = :status"
	}
	whereSql += " m.parent_id"
	list = make([]*systemModels.SysPermissionVo, 0, 2)

	return dataUtil.NamedQueryList(db, list, Permission, md.PermissionSql+whereSql, "", "")

}

func (md *SysPermissionDao) InsertPermission(db dataUtil.DB, Permission *systemModels.SysPermissionDML) {
	insertSQL := `insert into sys_Permission(permission_id,permission_name,parent_id,perms,create_by,create_time,update_by,update_time %s)
					values(:permission_id,:permission_name,:parent_id,:perms,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""

	if Permission.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if Permission.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, Permission)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysPermissionDao) UpdatePermission(db dataUtil.DB, Permission *systemModels.SysPermissionDML) {
	updateSQL := `update sys_Permission set update_time = now() , update_by = :update_by`

	if Permission.ParentId != 0 {
		updateSQL += ",parent_id = :parent_id"
	}
	if Permission.Status != "" {
		updateSQL += ",status = :status"
	}
	if Permission.Perms != "" {
		updateSQL += ",perms = :perms"
	}
	if Permission.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where permission_id = :permission_id"

	_, err := db.NamedExec(updateSQL, Permission)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysPermissionDao) DeletePermissionById(db dataUtil.DB, PermissionId int64) {
	_, err := db.Exec("delete from sys_Permission where Permission_id = ?", PermissionId)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysPermissionDao) SelectPermissionPermsByUserId(db dataUtil.DB, userId int64) (perms []string) {
	sqlStr := `	select distinct m.perms
				from sys_Permission m
				left join sys_role_Permission rm on m.Permission_id = rm.Permission_id
				left join sys_user_role ur on rm.role_id = ur.role_id
				left join sys_role r on r.role_id = ur.role_id
				where m.status = '0' and r.status = '0' and ur.user_id =  ?`
	perms = make([]string, 0, 2)
	err := db.Select(&perms, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysPermissionDao) SelectPermissionTreeAll(db dataUtil.DB) (sysPermissions []*systemModels.SysPermissionVo) {
	whereSql := ` where m.status = 0
		order by m.parent_id`
	sysPermissions = make([]*systemModels.SysPermissionVo, 0, 2)
	err := db.Select(&sysPermissions, md.PermissionSql+whereSql)
	if err != nil {
		panic(err)
	}
	return
}
func (md *SysPermissionDao) SelectPermissionTreeByUserId(db dataUtil.DB, userId int64) (sysPermissions []*systemModels.SysPermissionVo) {
	whereSql := ` left join sys_role_Permission rm on m.Permission_id = rm.Permission_id
			 left join sys_user_role ur on rm.role_id = ur.role_id
			 left join sys_role ro on ur.role_id = ro.role_id
			 left join sys_user u on ur.user_id = u.user_id
		where u.user_id = ? and m.status = 0  AND ro.status = 0
		order by m.parent_id`
	sysPermissions = make([]*systemModels.SysPermissionVo, 0, 2)
	err := db.Select(&sysPermissions, md.PermissionSql+whereSql, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysPermissionDao) CheckPermissionNameUnique(db dataUtil.DB, permissionName string, parentId int64) int64 {
	var roleId int64 = 0
	err := db.Get(&roleId, "select permission_id from sys_Permission where permission_name=? and parent_id = ?", permissionName, parentId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func (md *SysPermissionDao) HasChildByPermissionId(db dataUtil.DB, permissionId int64) int {
	var count = 0
	err := db.Get(&count, "select count(1) from sys_Permission where parent_id = ?", permissionId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
func (md *SysPermissionDao) SelectPermissionListByRoleId(db dataUtil.DB, roleId int64, PermissionCheckStrictly bool) (roleIds []string) {
	var err error
	roleIds = make([]string, 0, 2)
	sqlstr := `select m.permission_id
		from sys_permission m
            left join sys_role_permission rm on m.permission_id = rm.permission_id
        where rm.role_id = ?`
	if PermissionCheckStrictly {
		sqlstr += " and m.permission_id not in (select m.parent_id from sys_permission m inner join sys_role_permission rm on m.permission_id = rm.permission_id and rm.role_id = ?)"
		err = db.Select(&roleIds, sqlstr, roleId, roleId)
	} else {
		err = db.Select(&roleIds, sqlstr, roleId)
	}

	if err != nil {
		panic(err)
	}
	return
}
