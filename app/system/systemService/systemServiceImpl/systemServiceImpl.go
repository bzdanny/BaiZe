package systemServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewDeptService, NewDictDataService, NewDictTypeService, NewMenuService, NewPostService, NewRoleService, NewUserService, NewLoginService)
