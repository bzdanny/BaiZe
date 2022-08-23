package systemDaoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSysDeptDao, NewSysDictDataDao, NewSysDictTypeDao, NewSysPermissionDao,
	NewSysPostDao, NewSysRoleDeptDao, NewSysRoleDao, NewSysRolePermissionDao, NewSysUserPostDao, NewSysUserRoleDao, GetSysUserDao)
