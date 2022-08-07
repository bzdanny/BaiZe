package systemServiceImpl

import (
	systemDao "github.com/bzdanny/BaiZe/app/system/systemDao"
	"github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/bzdanny/BaiZe/baize/utils/exceLize"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
	"strconv"
	"strings"
)

type RoleService struct {
	data              *datasource.Data
	roleDao           systemDao.IRoleDao
	rolePermissionDao systemDao.IRolePermissionDao
	roleDeptDao       systemDao.IRoleDeptDao
	userRoleDao       systemDao.IUserRoleDao
}

func NewRoleService(data *datasource.Data, rd *systemDaoImpl.SysRoleDao, rmd *systemDaoImpl.SysRolePermissionDao, rdd *systemDaoImpl.SysRoleDeptDao, urd *systemDaoImpl.SysUserRoleDao) *RoleService {
	return &RoleService{
		data:              data,
		roleDao:           rd,
		rolePermissionDao: rmd,
		roleDeptDao:       rdd,
		userRoleDao:       urd,
	}
}

func (roleService *RoleService) SelectRoleList(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, count *int64) {
	return roleService.roleDao.SelectRoleList(roleService.data.GetSlaveDb(), role)

}
func (roleService *RoleService) RoleExport(role *systemModels.SysRoleDQL) (data []byte) {
	list, _ := roleService.roleDao.SelectRoleList(roleService.data.GetSlaveDb(), role)
	rows := systemModels.SysRoleListToRows(list)
	return exceLize.SetRows(rows)
}

func (roleService *RoleService) SelectRoleById(roseId int64) (role *systemModels.SysRoleVo) {
	return roleService.roleDao.SelectRoleById(roleService.data.GetSlaveDb(), roseId)

}

func (roleService *RoleService) InsertRole(sysRole *systemModels.SysRoleAdd) {
	sysRole.RoleId = snowflake.GenID()
	tx, err := roleService.data.GetMasterDb().Beginx()
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
	roleService.roleDao.InsertRole(tx, sysRole)
	PermissionIds := sysRole.PermissionIds
	l := len(PermissionIds)
	if l != 0 {
		list := make([]*systemModels.SysRolePermission, 0, l)
		for _, PermissionId := range PermissionIds {
			intPermissionId, _ := strconv.ParseInt(PermissionId, 10, 64)
			list = append(list, &systemModels.SysRolePermission{RoleId: sysRole.RoleId, PermissionId: intPermissionId})
		}
		roleService.insertRolePermission(tx, list)
	}

	return
}

func (roleService *RoleService) UpdateRole(sysRole *systemModels.SysRoleEdit) {
	tx, err := roleService.data.GetMasterDb().Beginx()
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
	roleService.roleDao.UpdateRole(tx, sysRole)
	roleService.rolePermissionDao.DeleteRolePermissionByRoleId(tx, sysRole.RoleId)
	PermissionIds := sysRole.PermissionIds
	l := len(PermissionIds)
	if l != 0 {
		list := make([]*systemModels.SysRolePermission, 0, l)
		for _, PermissionId := range PermissionIds {
			intPermissionId, _ := strconv.ParseInt(PermissionId, 10, 64)
			list = append(list, &systemModels.SysRolePermission{RoleId: sysRole.RoleId, PermissionId: intPermissionId})
		}
		roleService.insertRolePermission(tx, list)
	}
	return
}

func (roleService *RoleService) UpdateRoleStatus(sysRole *systemModels.SysRoleEdit) {
	roleService.roleDao.UpdateRole(roleService.data.GetMasterDb(), sysRole)
	return
}
func (roleService *RoleService) AuthDataScope(sysRole *systemModels.SysRoleEdit) {
	tx, err := roleService.data.GetMasterDb().Beginx()
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
	roleService.roleDao.UpdateRole(tx, sysRole)
	roleService.roleDeptDao.DeleteRoleDeptByRoleId(tx, sysRole.RoleId)
	roleService.insertRoleDept(tx, sysRole)
	return
}

func (roleService *RoleService) DeleteRoleByIds(ids []int64) {
	tx, err := roleService.data.GetMasterDb().Beginx()
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
	roleService.rolePermissionDao.DeleteRolePermission(tx, ids)
	roleService.roleDeptDao.DeleteRoleDept(tx, ids)
	roleService.roleDao.DeleteRoleByIds(tx, ids)
}
func (roleService *RoleService) CountUserRoleByRoleId(ids []int64) bool {
	return roleService.userRoleDao.CountUserRoleByRoleId(roleService.data.GetSlaveDb(), ids) > 0
}

func (roleService *RoleService) SelectBasicRolesByUserId(userId int64) (roles []*systemModels.SysRole) {
	return roleService.roleDao.SelectBasicRolesByUserId(roleService.data.GetSlaveDb(), userId)

}
func (roleService *RoleService) SelectRoleAll(role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo) {
	list, _ = roleService.roleDao.SelectRoleList(roleService.data.GetSlaveDb(), role)
	return
}

func (roleService *RoleService) RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []*baizeEntity.Role) {
	loginRoles = make([]*baizeEntity.Role, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		loginRoles = append(loginRoles, &baizeEntity.Role{RoleId: role.RoleId, DataScope: role.DataScope})
	}
	return
}
func (roleService *RoleService) SelectRoleListByUserId(userId int64) (list []int64) {
	return roleService.roleDao.SelectRoleListByUserId(roleService.data.GetSlaveDb(), userId)

}

func (roleService *RoleService) insertRolePermission(db dataUtil.DB, list []*systemModels.SysRolePermission) {
	if len(list) != 0 {
		roleService.rolePermissionDao.BatchRolePermission(db, list)
	}
	return
}

func (roleService *RoleService) CheckRoleNameUnique(id int64, roleName string) bool {
	RoleId := roleService.roleDao.CheckRoleNameUnique(roleService.data.GetSlaveDb(), roleName)
	if RoleId == id || RoleId == 0 {
		return false
	}
	return true
}

func (roleService *RoleService) CheckRoleKeyUnique(id int64, roleKey string) bool {
	RoleId := roleService.roleDao.CheckRoleKeyUnique(roleService.data.GetSlaveDb(), roleKey)
	if RoleId == id || RoleId == 0 {
		return false
	}
	return true
}
func (roleService *RoleService) SelectUserRoleGroupByUserId(userId int64) string {
	roles := roleService.roleDao.SelectBasicRolesByUserId(roleService.data.GetSlaveDb(), userId)
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	return strings.Join(roleNames, ",")

}
func (roleService *RoleService) insertRoleDept(db dataUtil.DB, sysRole *systemModels.SysRoleEdit) {
	deptIds := sysRole.DeptIds
	if len(deptIds) != 0 {
		list := make([]*systemModels.SysRoleDept, 0, len(deptIds))
		for _, deptId := range deptIds {
			intDeptId, _ := strconv.ParseInt(deptId, 10, 64)
			list = append(list, &systemModels.SysRoleDept{RoleId: sysRole.RoleId, DeptId: intDeptId})
		}
		roleService.roleDeptDao.BatchRoleDept(db, list)
	}

}
func (roleService *RoleService) SelectAllocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	return roleService.roleDao.SelectAllocatedList(roleService.data.GetSlaveDb(), user)
}

func (roleService *RoleService) SelectUnallocatedList(user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total *int64) {
	return roleService.roleDao.SelectUnallocatedList(roleService.data.GetSlaveDb(), user)

}

func (roleService *RoleService) InsertAuthUsers(roleId int64, userIds []int64) {
	if len(userIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(userIds))
		for _, userId := range userIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		roleService.userRoleDao.BatchUserRole(roleService.data.GetMasterDb(), list)
	}
}
func (roleService *RoleService) DeleteAuthUsers(roleId int64, userIds []int64) {
	roleService.userRoleDao.DeleteUserRoleInfos(roleService.data.GetMasterDb(), roleId, userIds)
}
func (roleService *RoleService) DeleteAuthUserRole(userRole *systemModels.SysUserRole) {
	roleService.userRoleDao.DeleteUserRoleInfo(roleService.data.GetMasterDb(), userRole)
}
