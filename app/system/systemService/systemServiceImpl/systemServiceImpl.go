package systemServiceImpl

var ProviderSet = wire.NewSet(NewDeptService, NewDictDataService, NewDictTypeService, NewMenuService, NewPostService, NewRoleService, NewUserService)
