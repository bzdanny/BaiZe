package systemServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewDeptService, NewDictDataService, NewDictTypeService, NewPermissionService, NewPostService, NewRoleService, NewUserService, NewLoginService)
