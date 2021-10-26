package systemService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
	"strconv"
	"strings"
)

func SelectRoleList(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, count *int64) {
	return systemDao.SelectRoleList(role)

}
func RoleExport(role *systemModels.SysRoleDQL) (data []byte) {
	sysUserList, _ := systemDao.SelectRoleList(role)
	rows := systemModels.SysRoleDMLListToRows(sysUserList)
	return exceLize.SetRows(rows)
}

func SelectRoleById(roseId int64) (role *systemModels.SysRoleVo) {
	return systemDao.SelectRoleById(roseId)

}

func InsertRole(sysRole *systemModels.SysRoleDML) {
	sysRole.RoleId = snowflake.GenID()
	systemDao.InsertRole(sysRole)
	InsertRoleMenu(sysRole)
	return
}

func UpdateRole(sysRole *systemModels.SysRoleDML) {
	systemDao.UpdateRole(sysRole)

	systemDao.DeleteRoleMenuByRoleId(sysRole.RoleId)

	InsertRoleMenu(sysRole)
	return
}

func UpdateRoleStatus(sysRole *systemModels.SysRoleDML) {
	systemDao.UpdateRole(sysRole)
	return
}
func AuthDataScope(sysRole *systemModels.SysRoleDML) {
	systemDao.UpdateRole(sysRole)

	systemDao.DeleteRoleDeptByRoleId(sysRole.RoleId)
	insertRoleDept(sysRole)
	return
}

func insertRoleDept(sysRole *systemModels.SysRoleDML) {
	deptIds := sysRole.DeptIds
	if len(deptIds) != 0 {
		list := make([]*systemModels.SysRoleDept, 0, len(deptIds))
		for _, deptId := range deptIds {
			intDeptId, _ := strconv.ParseInt(deptId, 10, 64)
			list = append(list, &systemModels.SysRoleDept{RoleId: sysRole.RoleId, DeptId: intDeptId})
		}
		systemDao.BatchRoleDept(list)
	}

	return
}
func DeleteRoleByIds(ids []int64) (err error) {

	systemDao.DeleteRoleMenu(ids)
	systemDao.DeleteRoleDept(ids)
	systemDao.DeleteRoleByIds(ids)
	return
}
func CountUserRoleByRoleId(ids []int64) bool {
	return systemDao.CountUserRoleByRoleId(ids) > 0
}

func SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole) {
	return systemDao.SelectBasicRolesByUserId(userId)

}
func SelectRoleAll(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo) {
	list, _ = systemDao.SelectRoleList(role)
	return
}

func RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []*loginModels.Role) {
	loginRoles = make([]*loginModels.Role, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		loginRoles = append(loginRoles, &loginModels.Role{RoleId: role.RoleId, DataScope: role.DataScope})
	}
	return
}
func SelectRoleListByUserId(userId int64) (list []int64) {
	return systemDao.SelectRoleListByUserId(userId)

}

func InsertRoleMenu(sysRole *systemModels.SysRoleDML) {
	menuIds := sysRole.MenuIds
	if len(menuIds) != 0 {
		list := make([]*systemModels.SysRoleMenu, 0, len(menuIds))
		for _, menuId := range menuIds {
			intMenuId, _ := strconv.ParseInt(menuId, 10, 64)
			list = append(list, &systemModels.SysRoleMenu{RoleId: sysRole.RoleId, MenuId: intMenuId})
		}
		systemDao.BatchRoleMenu(list)
	}

	return
}

func CheckRoleNameUnique(sysRole *systemModels.SysRoleDML) bool {
	RoleId := systemDao.CheckRoleNameUnique(sysRole.RoleName)
	if RoleId == sysRole.RoleId || sysRole.RoleId == 0 {
		return false
	}
	return true
}

func CheckRoleKeyUnique(sysRole *systemModels.SysRoleDML) bool {
	RoleId := systemDao.CheckRoleKeyUnique(sysRole.RoleKey)
	if RoleId == sysRole.RoleId || sysRole.RoleId == 0 {
		return false
	}
	return true
}
func SelectUserRoleGroupByUserId(userId int64) string {
	roles := systemDao.SelectBasicRolesByUserId(userId)
	roleNames := make([]string, 0, len(roles))
	//strings:=[len(roles)]string
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	return strings.Join(roleNames, ",")

}
