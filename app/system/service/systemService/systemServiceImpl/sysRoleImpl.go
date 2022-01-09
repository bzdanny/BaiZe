package systemServiceImpl

import (
	"baize/app/common/datasource"
	"baize/app/system/dao/systemDao"
	"baize/app/system/dao/systemDao/systemDaoImpl"
	"baize/app/system/models/loginModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
	"strconv"
	"strings"
)

var roleServiceImpl *roleService

func init() {
	roleServiceImpl = &roleService{
		roleDao:     systemDaoImpl.GetSysRoleDao(),
		roleMenuDao: systemDaoImpl.GetSysRoleMenuDao(),
		roleDeptDao: systemDaoImpl.GetSysRoleDeptDao(),
		userRoleDao: systemDaoImpl.GetSysUserRoleDao(),
	}
}

type roleService struct {
	roleDao     systemDao.IRoleDao
	roleMenuDao systemDao.IRoleMenuDao
	roleDeptDao systemDao.IRoleDeptDao
	userRoleDao systemDao.IUserRoleDao
}

func GetRoleService() *roleService {

	return roleServiceImpl
}

func (roleService *roleService) SelectRoleList(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, count *int64) {
	return roleService.roleDao.SelectRoleList(role)

}
func (roleService *roleService) RoleExport(role *systemModels.SysRoleDQL) (data []byte) {
	list, _ := roleService.roleDao.SelectRoleList(role)
	rows := systemModels.SysRoleListToRows(list)
	return exceLize.SetRows(rows)
}

func (roleService *roleService) SelectRoleById(roseId int64) (role *systemModels.SysRoleVo) {
	return roleService.roleDao.SelectRoleById(roseId)

}

func (roleService *roleService) InsertRole(sysRole *systemModels.SysRoleDML) {
	sysRole.RoleId = snowflake.GenID()
	tx, err := datasource.GetMasterDb().Beginx()
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
	roleService.roleDao.InsertRole(sysRole, tx)
	roleService.insertRoleMenu(sysRole, tx)
	return
}

func (roleService *roleService) UpdateRole(sysRole *systemModels.SysRoleDML) {
	tx, err := datasource.GetMasterDb().Beginx()
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
	roleService.roleDao.UpdateRole(sysRole, tx)
	roleService.roleMenuDao.DeleteRoleMenuByRoleId(sysRole.RoleId, tx)
	roleService.insertRoleMenu(sysRole, tx)
	return
}

func (roleService *roleService) UpdateRoleStatus(sysRole *systemModels.SysRoleDML) {
	roleService.roleDao.UpdateRole(sysRole)
	return
}
func (roleService *roleService) AuthDataScope(sysRole *systemModels.SysRoleDML) {
	tx, err := datasource.GetMasterDb().Beginx()
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
	roleService.roleDao.UpdateRole(sysRole, tx)
	roleService.roleDeptDao.DeleteRoleDeptByRoleId(sysRole.RoleId, tx)
	roleService.insertRoleDept(sysRole, tx)
	return
}

func (roleService *roleService) DeleteRoleByIds(ids []int64) {
	tx, err := datasource.GetMasterDb().Beginx()
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
	roleService.roleMenuDao.DeleteRoleMenu(ids, tx)
	roleService.roleDeptDao.DeleteRoleDept(ids, tx)
	roleService.roleDao.DeleteRoleByIds(ids, tx)
}
func (roleService *roleService) CountUserRoleByRoleId(ids []int64) bool {
	return roleService.userRoleDao.CountUserRoleByRoleId(ids) > 0
}

func (roleService *roleService) SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole) {
	return roleService.roleDao.SelectBasicRolesByUserId(userId)

}
func (roleService *roleService) SelectRoleAll(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo) {
	list, _ = roleService.roleDao.SelectRoleList(role)
	return
}

func (roleService *roleService) RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []*loginModels.Role) {
	loginRoles = make([]*loginModels.Role, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		loginRoles = append(loginRoles, &loginModels.Role{RoleId: role.RoleId, DataScope: role.DataScope})
	}
	return
}
func (roleService *roleService) SelectRoleListByUserId(userId int64) (list []int64) {
	return roleService.roleDao.SelectRoleListByUserId(userId)

}

func (roleService *roleService) insertRoleMenu(sysRole *systemModels.SysRoleDML, tx ...datasource.Transaction) {
	menuIds := sysRole.MenuIds
	if len(menuIds) != 0 {
		list := make([]*systemModels.SysRoleMenu, 0, len(menuIds))
		for _, menuId := range menuIds {
			intMenuId, _ := strconv.ParseInt(menuId, 10, 64)
			list = append(list, &systemModels.SysRoleMenu{RoleId: sysRole.RoleId, MenuId: intMenuId})
		}
		roleService.roleMenuDao.BatchRoleMenu(list, tx...)
	}

	return
}

func (roleService *roleService) CheckRoleNameUnique(sysRole *systemModels.SysRoleDML) bool {
	RoleId := roleService.roleDao.CheckRoleNameUnique(sysRole.RoleName)
	if RoleId == sysRole.RoleId || sysRole.RoleId == 0 {
		return false
	}
	return true
}

func (roleService *roleService) CheckRoleKeyUnique(sysRole *systemModels.SysRoleDML) bool {
	RoleId := roleService.roleDao.CheckRoleKeyUnique(sysRole.RoleKey)
	if RoleId == sysRole.RoleId || sysRole.RoleId == 0 {
		return false
	}
	return true
}
func (roleService *roleService) SelectUserRoleGroupByUserId(userId int64) string {
	roles := roleService.roleDao.SelectBasicRolesByUserId(userId)
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	return strings.Join(roleNames, ",")

}
func (roleService *roleService) insertRoleDept(sysRole *systemModels.SysRoleDML, tx ...datasource.Transaction) {
	deptIds := sysRole.DeptIds
	if len(deptIds) != 0 {
		list := make([]*systemModels.SysRoleDept, 0, len(deptIds))
		for _, deptId := range deptIds {
			intDeptId, _ := strconv.ParseInt(deptId, 10, 64)
			list = append(list, &systemModels.SysRoleDept{RoleId: sysRole.RoleId, DeptId: intDeptId})
		}
		roleService.roleDeptDao.BatchRoleDept(list, tx...)
	}

}
func (roleService *roleService) SelectAllocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	return roleService.roleDao.SelectAllocatedList(user)
}

func (roleService *roleService) SelectUnallocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	return roleService.roleDao.SelectUnallocatedList(user)

}

func (roleService *roleService) InsertAuthUsers(roleId int64, userIds []int64) {
	if len(userIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(userIds))
		for _, userId := range userIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		roleService.userRoleDao.BatchUserRole(list)
	}
}
func (roleService *roleService) DeleteAuthUsers(roleId int64, userIds []int64) {
	roleService.userRoleDao.DeleteUserRoleInfos(roleId, userIds)
}
func (roleService *roleService) DeleteAuthUserRole(userRole *systemModels.SysUserRole) {
	roleService.userRoleDao.DeleteUserRoleInfo(userRole)
}
