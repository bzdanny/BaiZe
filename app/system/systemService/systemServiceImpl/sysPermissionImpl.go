package systemServiceImpl

import (
	systemDao "github.com/bzdanny/BaiZe/app/system/systemDao"
	systemDaoImpl "github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	systemModels "github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/utils"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
)

type PermissionService struct {
	data              *datasource.Data
	PermissionDao     systemDao.IPermissionDao
	rolePermissionDao systemDao.IRolePermissionDao
	roleDao           systemDao.IRoleDao
}

func NewPermissionService(data *datasource.Data, md *systemDaoImpl.SysPermissionDao, rmd *systemDaoImpl.SysRolePermissionDao, rd *systemDaoImpl.SysRoleDao) *PermissionService {
	return &PermissionService{
		data:              data,
		PermissionDao:     md,
		rolePermissionDao: rmd,
		roleDao:           rd,
	}
}

func (PermissionService *PermissionService) SelectPermissionList(Permission *systemModels.SysPermissionDQL, userId int64) (list []*systemModels.SysPermissionVo) {
	if utils.IsAdmin(userId) {
		list = PermissionService.PermissionDao.SelectPermissionList(PermissionService.data.GetSlaveDb(), Permission)
	} else {
		Permission.UserId = userId
		list = PermissionService.PermissionDao.SelectPermissionListByUserId(PermissionService.data.GetSlaveDb(), Permission)
	}
	return
}

func (PermissionService *PermissionService) SelectPermissionById(PermissionId int64) (Permission *systemModels.SysPermissionVo) {
	return PermissionService.PermissionDao.SelectPermissionById(PermissionService.data.GetSlaveDb(), PermissionId)

}
func (PermissionService *PermissionService) InsertPermission(Permission *systemModels.SysPermissionDML) {
	Permission.PermissionId = snowflake.GenID()
	PermissionService.PermissionDao.InsertPermission(PermissionService.data.GetMasterDb(), Permission)
}
func (PermissionService *PermissionService) UpdatePermission(Permission *systemModels.SysPermissionDML) {
	PermissionService.PermissionDao.UpdatePermission(PermissionService.data.GetMasterDb(), Permission)
}
func (PermissionService *PermissionService) DeletePermissionById(PermissionId int64) {
	PermissionService.PermissionDao.DeletePermissionById(PermissionService.data.GetMasterDb(), PermissionId)
}

func (PermissionService *PermissionService) SelectPermissionTreeByUserId(userId int64) (sysPermission []*systemModels.SysPermissionVo) {

	return
}

func (PermissionService *PermissionService) BuildPermissions(sysPermissions []*systemModels.SysPermissionVo) []*systemModels.RouterVo {
	routerVo := make([]*systemModels.RouterVo, 0, 2)

	return routerVo
}

func (PermissionService *PermissionService) CheckPermissionNameUnique(Permission *systemModels.SysPermissionDML) bool {
	RoleId := PermissionService.PermissionDao.CheckPermissionNameUnique(PermissionService.data.GetSlaveDb(), Permission.PermissionName, Permission.ParentId)
	if RoleId == Permission.PermissionId || RoleId == 0 {
		return false
	}
	return true
}

func (PermissionService *PermissionService) HasChildByPermissionId(PermissionId int64) bool {
	return PermissionService.PermissionDao.HasChildByPermissionId(PermissionService.data.GetSlaveDb(), PermissionId) > 0
}

func (PermissionService *PermissionService) CheckPermissionExistRole(PermissionId int64) bool {
	return PermissionService.rolePermissionDao.CheckPermissionExistRole(PermissionService.data.GetSlaveDb(), PermissionId) > 0
}
func (PermissionService *PermissionService) SelectPermissionListByRoleId(roleId int64) []string {
	role := PermissionService.roleDao.SelectRoleById(PermissionService.data.GetSlaveDb(), roleId)
	return PermissionService.PermissionDao.SelectPermissionListByRoleId(PermissionService.data.GetSlaveDb(), roleId, role.PermissionCheckStrictly)
}

func getChildPerms(Permission []*systemModels.SysPermissionVo, parentId int64) []*systemModels.SysPermissionVo {
	sysPermissions := make([]*systemModels.SysPermissionVo, 0, 2)
	for _, sysPermission := range Permission {
		if sysPermission.ParentId == parentId {
			recursionFn(Permission, sysPermission)
			sysPermissions = append(sysPermissions, sysPermission)
		}
	}
	return sysPermissions
}

/**
 * 递归列表
 *
 * @param Permission
 * @param s
 */
func recursionFn(Permission []*systemModels.SysPermissionVo, s *systemModels.SysPermissionVo) {
	childList := getChildList(Permission, s)
	s.Children = childList
	for _, sysPermission := range childList {
		if hasChild(Permission, sysPermission) {
			recursionFn(Permission, sysPermission)
		}
	}
}

/**
 * 判断是否有子节点
 */
func hasChild(list []*systemModels.SysPermissionVo, m *systemModels.SysPermissionVo) bool {
	return len(getChildList(list, m)) > 0
}

/**
 * 得到子节点列表
 */
func getChildList(Permission []*systemModels.SysPermissionVo, s *systemModels.SysPermissionVo) []*systemModels.SysPermissionVo {
	tlist := make([]*systemModels.SysPermissionVo, 0, 2)
	for _, sysPermission := range Permission {
		if sysPermission.ParentId == s.PermissionId {
			tlist = append(tlist, sysPermission)
		}

	}
	return tlist
}
